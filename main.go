package main

import "net/http"
import "log"
import "fmt"
//import "html"
import "io"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		reqHostname := r.URL.Query().Get("q")
		reqMethod := r.Method
		client := &http.Client{}
		req, err := http.NewRequest(reqMethod, reqHostname, nil)
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if resp.StatusCode > 299 {
			log.Fatalf("Response failed with status code: %d and\nbody: %s\n", resp.StatusCode, body)
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, "%s", body)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
