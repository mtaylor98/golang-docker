//This is a basic webserver written in Go

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func helloHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		res,
		`<doctype html>
<html>
					<head>
								<title>Hello Gopher</title>
					</head>
					<body>
								Hello Gopher </br>
								It is really awesome that Docker is written in Go!
					</body>
</html>`,
				)
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
				fmt.Fprint(w, "Go web app powered by Docker")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Server", "A Go Web Server")
	w.WriteHeader(200)
	w.Write([]byte("OK"))
}

func treyHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		res,
		`<doctype html>
<html>
				<head>
							<title>Trey Face</title>
				</head>
				<body>
							<img src="https://i.imgur.com/l3SzCI5.jpg?1">
				</body>
</html>`,
				)
}

func main () {
				http.HandleFunc("/", treyHandler)
				http.HandleFunc("/lb-status", healthHandler)
				http.HandleFunc("/hello", helloHandler)
				http.HandleFunc("/default", defaultHandler)
				err := http.ListenAndServe(":80", nil)
				if err != nil {
								log.Fatal("ListenAndServe: ", err)
								return
				}
}
