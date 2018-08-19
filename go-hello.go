package main

import "net/http"

func main() {
	// Handle / path with hello function
	http.HandleFunc("/", helloWorldHandler)
	// listen on port 8080
	http.ListenAndServe(":8080", nil)
}

/*
  Go language uses MixedCaps or mixedCaps
  The visibility of a name outside a package is determined by whether its first character is upper case
  https://golang.org/doc/effective_go.html#mixed-caps

  For example:
  writeToMongoDB // unexported, only visible within the package
  WriteToMongoDB // exported
*/

func helloWorldHandler(rw http.ResponseWriter, req *http.Request) {

	// allow cross domain AJAX requests
	if origin := req.Header.Get("Origin"); origin != "" {
		rw.Header().Set("Access-Control-Allow-Origin", origin)
		rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		rw.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}

	// Stop here if its Preflighted OPTIONS request
	if req.Method == "OPTIONS" {
		rw.WriteHeader(http.StatusTeapot) // I'm a teapot
		return
	}

	// send my Hello World back
	rw.Write([]byte("Hello, world\n"))
}
