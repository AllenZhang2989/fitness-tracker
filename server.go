package main

// in go every program starts in the main package, which onctains the main function where execution begins

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	//for SQLite shit

	_ "github.com/mattn/go-sqlite3" // Import the SQLite driver
)

//fmt - formated input/output functions
//io/ioutil - contain functions to read and write files and data
//log -  logging functiolkns
//net/http - core package for building http servers

// Serve the HTML form
func formHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "form.html")
}


// Handle incoming POST requests
func messageHandler(w http.ResponseWriter, r *http.Request) {
	// if incoming data is a post request
	if r.Method == http.MethodPost {
		body, err := ioutil.ReadAll(r.Body) // Read the message body
		if err != nil {
			http.Error(w, "COCK Unable to read body", http.StatusBadRequest)
			return
		}
		message := string(body)
		fmt.Println("Received message:", message) // Log the received message

		// Respond to the sender
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Message received: " + message))
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
// processes http requests to the /send endpoint
// w http.ResponseWriter - outgoing stream to the client
// r *http.Request - incoming http request

func main() {
	
    // Serve the form at the root URL
    http.HandleFunc("/", formHandler)

    // Handle the /send endpoint for form submission
    http.HandleFunc("/send", messageHandler)

    fmt.Println("Server is listening on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))

	
}