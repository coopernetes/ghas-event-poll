package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/google/go-github/v47/github"
)

func main() {
	client := github.NewClient(nil)

	repos, _, err := client.Repositories.List(context.Background(), "coopernetes", nil)
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello %s!", r.UserAgent())
	})
	http.HandleFunc("/repos", func(w http.ResponseWriter, r *http.Request) {
		var sb strings.Builder
		for _, repo := range repos {
			fmt.Fprintf(&sb, "%s\n", repo.GetFullName())
		}
		fmt.Fprintf(w, "%s", sb.String())
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
