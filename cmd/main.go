package main

import "url-shortening-service/internal/router"

func main() {
	r := router.NewRouter()
	r.Run(":8080")
}
