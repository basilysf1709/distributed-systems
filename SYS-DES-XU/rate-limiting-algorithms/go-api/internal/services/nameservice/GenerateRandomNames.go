package nameservice

import (
    "go-api/pkg/utils"
)

func GenerateRandomNames(count int, length int) []string {
    var names []string
    for i := 0; i < count; i++ {
        names = append(names, utils.RandomString(length))
    }
    return names
}
