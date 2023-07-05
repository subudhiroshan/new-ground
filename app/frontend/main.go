package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

var (
	backendHost = "localhost"
	backendPort = "8080"
)

func callBackend(method string) (string, error) {
	requestURL := fmt.Sprintf("http://%s:%s/%s", backendHost, backendPort, method)
	resp, err := http.Get(requestURL)
	if err != nil {
		return "", err
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	body := string(bodyBytes)

	return body, nil
}

func get(w http.ResponseWriter, _ *http.Request) {
	res, err := callBackend("")
	if err != nil {
		fmt.Fprintf(w, "failed to get current value, %s\n", err)
	}

	fmt.Fprintln(w, res)
}

func process(w http.ResponseWriter, _ *http.Request) {
	res, err := callBackend("")
	if err != nil {
		fmt.Fprintf(w, "failed to get current value, %s\n", err)
	}

	currVal, err := strconv.Atoi(res)
	if err != nil {
		fmt.Fprintf(w, "failed to parse current value, %s\n", err)
	}

	if currVal > 50 {
		res, err = callBackend("/zero")
		if err != nil {
			fmt.Fprintf(w, "failed to zero value, %s\n", err)
		}
	} else if currVal%2 != 0 {
		res, err = callBackend("/double")
		if err != nil {
			fmt.Fprintf(w, "failed to double value, %s\n", err)
		}
	} else {
		res, err = callBackend("/increase")
		if err != nil {
			fmt.Fprintf(w, "failed to increase value, %s\n", err)
		}
	}
}

func init() {
	if host := os.Getenv("BACKEND_HOST"); host != "" {
		backendHost = host
	}

	if port := os.Getenv("BACKEND_PORT"); port != "" {
		backendPort = port
	}
}

func main() {
	http.HandleFunc("/", get)
	http.HandleFunc("/process", process)

	log.Fatal(http.ListenAndServe("0.0.0.0:8090", nil))
}
