package main

import (
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/InVisionApp/go-health"
	"github.com/InVisionApp/go-health/checkers"
	"github.com/InVisionApp/go-health/handlers"
)

func main() {
	// Create a new health instance
	h := health.New()
	goodTestURL, _ := url.Parse("http://localhost:8001/time")

	// Create a couple of checks
	goodHTTPCheck, _ := checkers.NewHTTP(&checkers.HTTPConfig{
		URL: goodTestURL,
	})

	

	// Add the checks to the health instance
	h.AddChecks([]*health.Config{
		{
			Name:     "good-check",
			Checker:  goodHTTPCheck,
			Interval: time.Duration(10) * time.Second,
			Fatal:    true,
		},
	})

	//  Start the healthcheck process
	if err := h.Start(); err != nil {
		log.Fatalf("Unable to start healthcheck: %v", err)
	}

	log.Println("Server listening on :8008")

	// Define a healthcheck endpoint and use the built-in JSON handler
	http.HandleFunc("/healthcheck", handlers.NewJSONHandlerFunc(h, nil))
	http.ListenAndServe(":8008", nil)
}
