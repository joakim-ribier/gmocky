package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joakim-ribier/gmocky/helpers"
	"github.com/joakim-ribier/gmocky/utils"
)

var port = "8080"
var action = "writeMockURL"

func main() {
	if len(os.Args) == 2 {
		action = os.Args[1]
	}

	if v := os.Getenv("GMOCKY_PORT"); v != "" {
		port = v
	}

	switch action {
	case "start":
		start()
	default:
		helpers.NewMockURLWriter(port).Print()
	}
}

func start() {
	fmt.Printf("# Server waiting on :%s...\n\r", port)
	http.HandleFunc(utils.HandlerPattern, func(w http.ResponseWriter, r *http.Request) {
		helpers.NewResponseWriter(w, r).Write()
	})
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
}
