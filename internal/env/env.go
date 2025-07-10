package env

import (
	"fmt"
	"os"
	"strconv"
)

func GetString(key, fallback string) string{
	val, ok := os.LookupEnv(key)
    if !ok{
        fmt.Printf("key %v has not been fount\n", key)
        return fallback
    }
    return val
}

func GetInt(key string, fallback int) int {
    val, ok := os.LookupEnv(key)
    if !ok {
        return fallback
    }
    valAsInt, err := strconv.Atoi(val)
    if err != nil{
        return fallback
    }
    return valAsInt
}