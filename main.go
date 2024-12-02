package main

import (
	"fmt"
	"net/http"
)

// handler to serve the HTML page
func handler(w http.ResponseWriter, r *http.Request) {
	html := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Hello, Dog!</title>
	</head>
	<body>
		<h1>Hello, World!</h1>
		<img src="/dog.jpg" alt="A cute dog" width="300">
	</body>
	</html>
	`
	fmt.Fprint(w, html)
}

// handler to serve the dog image
func imageHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "dog.jpg")
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/dog.jpg", imageHandler)
	http.ListenAndServe(":8080", nil)
}
