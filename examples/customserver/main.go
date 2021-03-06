package main

import (
	"github.com/yarf-framework/yarf"
	"net/http"
	"time"
)

// Define a simple resource
type Hello struct {
	yarf.Resource // Extend the yarf.Resource by composition
}

// Implement the GET handler
func (h *Hello) Get(c *yarf.Context) error {
	c.Render("Hello world!")

	return nil
}

// Entry point of the executable application
// This time we setup a custom Go http server and use YARF as a router.
func main() {
	// Create a new empty YARF server
	y := yarf.New()

	// Add route/resource
	y.Add("/", new(Hello))

	// Configure custom http server and set the yarf object as Handler.
	s := &http.Server{
		Addr:           ":8080",
		Handler:        y,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
