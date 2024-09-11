package main

import "net/http"
import "log"
import "fmt"
//import "html"
import "io"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		reqHostname := r.URL.Query().Get("q")
		res, err := http.Get(reqHostname)
		if err != nil {
			log.Fatal(err)
		}
		body, err := io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode > 299 {
			log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s", body)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
