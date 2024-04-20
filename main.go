package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: irio <url>")
		os.Exit(1)
	}
	url := os.Args[1]

	if !strings.Contains(url, "https://") {
		url = "https://" + url
	}
	resp, err := http.Get(url)

	if err != nil {
		log.Fatalf(err.Error())
	}
	for k, v := range resp.Header {
		fmt.Printf("\033[33m%s:\033[0m ", k)
		fmt.Printf("%s\n", v)
	}
}
