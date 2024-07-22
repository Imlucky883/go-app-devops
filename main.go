// main.go
package main

import (
	"log"
	"net/http"
)

func main() {
	// Serve the image from the "images" directory
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))

	// Handle the root URL
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// HTML response to display the image
		html := `
        <!DOCTYPE html>
        <html>
        <head>
            <title>Simple Go Image Server</title>
        </head>
        <body>
            <h1>Welcome to the Simple Go Image Server</h1>
            <img src="/static/gofer.png" alt="Sample Image" />
        </body>
        </html>`
		w.Write([]byte(html))
	})

	// Start the server on port 8080
	log.Println("Starting server on :8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}
