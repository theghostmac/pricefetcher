package main

import (
	"io"
	"log"
	"net/http"
)

func displayHomePage(writer http.ResponseWriter, request *http.Request) {
	log.Println("Hello from the CLI")
	data, _ := io.ReadAll(request.Body)

	log.Printf("%s\n", data)
}

func displayExitPage(writer http.ResponseWriter, request *http.Request) {
	log.Println("Goodbye from the CLI")
}

func main() {
	http.HandleFunc("/", displayHomePage)
	http.HandleFunc("/exit", displayExitPage)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("failed to start the server due to : %s", err)
	}
}
