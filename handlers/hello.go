package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Hello struct {
}

func (h *Hello) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
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
