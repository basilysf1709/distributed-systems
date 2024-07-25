package middleware

import (
    "log"
    "net/http"
    "time"
    "go-api/internal/services/nameservice"
    "encoding/json"
    "strconv"
)


// In middlewares, a function is returned to promote function chaining og http.HandleFunc so that each middleware chain returns a function chain, for easier wrapping and functionality
// LoggingMiddleware logs the incoming HTTP request & its duration
func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        log.Printf("Started %s %s", r.Method, r.URL.Path)
        next.ServeHTTP(w, r)
        log.Printf("Completed in %v", time.Since(start))
    })
}


type ErrorResponse struct {
    Error       string `json:"error"`
    Description string `json:"description"`
    Status      int    `json:"status"`
}

func RateLimitMiddleware(tb *nameservice.TokenBucket) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            if !tb.Allow() {
                // Create an ErrorResponse struct
                timeLeft := tb.SecondsUntilNextRequest()
                errResp := ErrorResponse{
                    Error:       "Too Many Requests",
                    Description: "You have exceeded the rate limit. Please wait for " + strconv.Itoa(timeLeft) +  " seconds before making another request.",
                    Status:      http.StatusTooManyRequests,
                }

                // Encode the ErrorResponse to JSON
                jsonResponse, err := json.Marshal(errResp)
                if err != nil {
                    http.Error(w, "Internal Server Error", http.StatusInternalServerError)
                    return
                }

                // Send the JSON response with HTTP 429 status
                w.Header().Set("Content-Type", "application/json")
                w.WriteHeader(http.StatusTooManyRequests)
                w.Write(jsonResponse)
                return
            }
            next.ServeHTTP(w, r)
        })
    }
}
