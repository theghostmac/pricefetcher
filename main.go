package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func displayHomePage(writer http.ResponseWriter, request *http.Request) {
	log.Println("Hello from the CLI")
	data, err := io.ReadAll(request.Body)
	if err != nil {
		http.Error(writer, "Oops", http.StatusBadRequest)
		return	
	}

	_, err = fmt.Fprintf(writer, "Data is %s\n", data)
	if err != nil {
		fmt.Println("Error writing the writer: ", err)
	}
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
