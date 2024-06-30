package main

import (
	"fmt"
	"go-server/pkg/controllers"
	"net/http"
)

func main() {
	http.HandleFunc("/validation", controllers.LunhCheck)
	fmt.Print("Server is running at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server", err)
	}
}
