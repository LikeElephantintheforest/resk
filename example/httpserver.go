package main

import "net/http"

func main() {
	http.HandleFunc("/hello",
		func(writer http.ResponseWriter, request *http.Request) {
			writer.Write([]byte("helloWord"))
		})
	http.ListenAndServe(":8082", nil)
}
