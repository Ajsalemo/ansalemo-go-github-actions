package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		m := map[string]string{"msg": "Hello from Go with GitHub Actions!"}
		e := json.NewEncoder(w)
		e.Encode(m)
	})
	log.Println("Server listening on port 8080.")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}
