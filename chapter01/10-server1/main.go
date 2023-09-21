// In this sec tion, we’ll show a minimal ser ver that retur ns the pat h comp onent
// of the URL used to access the ser ver. That is, if the request is for http://local-
// host:8000/hello , the respons e wi l l be URL.Path = "/hello" .

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler( w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}