package main

import (
	"fmt"
	"html"
	"io"
	"log"
	"net/http"
	"strconv"
)

func addCountHandler() {
	countHandler := func(w http.ResponseWriter, req *http.Request) {
		count := 0
		if c, err := req.Cookie("count"); err == nil {
			if count, err = strconv.Atoi(c.Value); err != nil {
				log.Default().Print(err)
				count = 0
			}
		}
		count += 1
		http.SetCookie(w, &http.Cookie{
			Name:  "count",
			Value: strconv.Itoa(count),
		})
		str := fmt.Sprintf("You have visited: %d times.", count)
		log.Default().Print(str)
		io.WriteString(w, str)
	}
	http.HandleFunc("/count", countHandler)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	addCountHandler()
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
