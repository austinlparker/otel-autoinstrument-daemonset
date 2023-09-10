package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		messages := []string{
			"Keep your face always toward the sunshine - and shadows will fall behind you.",
			"Positive anything is better than negative nothing.",
			"The only limit to our realization of tomorrow is our doubts of today.",
			"Keep your face to the sunshine and you cannot see a shadow.",
			"Once you replace negative thoughts with positive ones, you'll start having positive results.",
		}

		rand.Seed(time.Now().UnixNano())
		message := messages[rand.Intn(len(messages))]
		fmt.Fprintf(w, message)
	})

	http.ListenAndServe(":8080", nil)
}
