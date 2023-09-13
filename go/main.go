package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

var startedAt = time.Now()

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/secret", secret)
	http.HandleFunc("/healthz", healthz)
	http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	podName := os.Getenv("HOSTNAME")

	res := fmt.Sprintf("<h1>Hello %s from pod %s!!!</h1>", name, podName)

	w.Write([]byte(res))
}

func secret(w http.ResponseWriter, r *http.Request) {
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")

	res := fmt.Sprintf("<h1>User: %s | Password: %s</h1>", user, password)

	w.Write([]byte(res))
}

func healthz(w http.ResponseWriter, r *http.Request) {
	duration := time.Since(startedAt).Seconds()

	// probe
	if duration < 10 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Healthcheck fail: Duration %v", duration)))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("All probes OK: Duration: %v", duration)))
}
