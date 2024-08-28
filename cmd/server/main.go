package main

import (
	"github.com/vegetarian23/go-initializr/internal/routers"
)

func main() {
	r := routers.SetupRouter()

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
