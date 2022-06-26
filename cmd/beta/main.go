package main

import (
	"fmt"
	"log"
	"multi-deploy/pkg/fizzbuzz"
	"net/http"
	"os"
	"strconv"
)

func handler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	n := q.Get("n")
	num, _ := strconv.Atoi(n)
	fmt.Fprintf(w, "beta: %s\n", fizzbuzz.Fizzbuzz(num))
}

func main() {
	// Cloud Runで動かす想定
	// 設定されている PORT 環境変数を参照
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/fizzbuzz", handler)
	err := http.ListenAndServe(":"+port, nil)
	fmt.Printf("[START] server. port: %s\n", port)
	if err != nil {
		log.Fatal(err)
	}

}
