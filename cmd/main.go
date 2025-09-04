package main

import (
	"log"
	"url-shortening-service/internal/router"
)

func main() {
	r, err := router.NewRouter()
	if err != nil {
		log.Fatalf("init router: %v", err)
	}
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("run server: %v", err)
	}
}
