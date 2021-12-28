package main

type HTTPrequest struct {
	Method string
}

import (
	"fmt"
)

func main() {
	r := HTTPrequest { Method: "HEAD"}
	
	switch r.Method {
		case "GET": println("GET Request")
		case "DELETE": println("DELETE Request")
		case "POST": println("POST Request")
		case "PUT": println("PUT request")
		default: println("Unhandled request")
	}
}
