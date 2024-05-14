package application

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redis/go-redis/v9"
)

type App struct {
  router http.Handler
	rdb *redis.Client
}

func New(redisAddr, redisPassword string) *App {
	
    // Intiliaze Redis Client
    rdb := redis.NewClient(&redis.Options{
        Addr:     redisAddr,
        Password: redisPassword,
        DB:       0, // default database
    })
		app := &App{
		router: loadRoutes(),
		rdb: rdb,
	}

    // Ping Redis to check if connection is successful
    
    return app
}



func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":3000",
		Handler: a.router,
	}

	err := server.ListenAndServe()
	if err != nil {
    return fmt.Errorf("failed to start server: %w", err)
	}
  
  return nil
}

func (a *App ) TestRedisConnection () error {
	// set a test key value pair
err := a.rdb.Set(context.Background(), "test_key", "test_value", 0).Err()
if err != nil {
	return fmt.Errorf("failed to set the key value pair: %w", err)
}

// Retrieve the value 
val, err := a.rdb.Get(context.Background(), "test_key").Result()
if err != nil {
	return fmt.Errorf("failed to get value: %w", err)
}

fmt.Println("Test key value", val)
fmt.Println("Connect to db succesfully!")

return nil
}