package main

import (
	"fmt"
	"net/http"
)

//register custom handler
type customHandler struct{}

//implement ServerHTTP method to satisfy the Header interface
func (h customHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	//set the response type in the header to be plain text
	w.Header().Set("response type", "plain/http")

	//print hello world if request method is a "GET"
	if r.Method == "GET" {
		fmt.Fprintf(w, "hello from golang")

	} else {

		//set the status of the response to a 405 if the method is not a "GET"
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "method not allowed")
	}

}

func main() {

	//instantiate a handler
	handler := customHandler{}

	//register the handler to the root route
	http.Handle("/", handler)

	fmt.Println("server listening on port :3000")

	//start the server on port 3000
	err := http.ListenAndServe(":4000", nil)

	//handle errors if they occur
	if err != nil {
		fmt.Println("error starting server", err)

	}
}
