package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/health", health)
	http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	podName, _ := os.Hostname()

	res := fmt.Sprintf("<h1>Hello K8S from pod %s!!!</h1>", podName[len(podName)-5:])

	w.Write([]byte(res))
}

func health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
