package utils

import (
    "math/rand"
    "time"
)

func RandomString(length int) string {
    letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
    rand.Seed(time.Now().UnixNano())
    s := make([]byte, length)
    for i := range s {
        s[i] = letters[rand.Intn(len(letters))]
    }
    return string(s)
}
