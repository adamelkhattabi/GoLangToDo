package main

import (
	"context"
	"fmt"

	"github.com/adamelkhattabi/GoLangToDo/application"
)

func main() {
    // Replace these values with your actual Redis connection details
    redisAddr := "redis-12519.c304.europe-west1-2.gce.redns.redis-cloud.com:12519"
    redisPassword := "miEKU3NvjfztOuVxr48loSNMb1VxjO0p"

    app := application.New(redisAddr, redisPassword)

    // Test Redis connection
    err := app.TestRedisConnection()
    if err != nil {
        fmt.Println("failed to test Redis connection:", err)
        return
    }

    // Start the application
    err = app.Start(context.TODO())
    if err != nil {
        fmt.Println("failed to start app:", err)
    }
}
