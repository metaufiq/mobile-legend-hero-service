package main

import (
	"errors"
	"fmt"
	"github/metaufiq/mobile-legend-hero-service/src/hero"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", hero.GetRoot)
	http.HandleFunc("/hero", hero.GetHero)
	fmt.Printf("server running...\n")
	err := http.ListenAndServe(":3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
