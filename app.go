package main

import (
	"errors"
	"fmt"
	"github/metaufiq/mobile-legend-hero-service/src/presentation"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	presentation.ServeRoutes()

	fmt.Printf("\n[ENV VARIABLES]\n")
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Printf("%s: %s\n", pair[0], pair[1])
	}

	fmt.Printf("\n\nserver running...\n\n")
	err := http.ListenAndServe(":8080", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
