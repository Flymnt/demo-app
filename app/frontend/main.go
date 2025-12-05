package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("OK"))
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		page := `
<!DOCTYPE html>
<html>
<head>
  <title>Flymnt demo frontend</title>
</head>
<body>
  <h1>Flymnt demo-app frontend</h1>
  <p>This is a simple Go HTTP server used for Port.io demo purposes.</p>
  <p>Downstream services (hard-coded for demo):</p>
  <ul>
    <li>service-a ("Hello from service-a" on port 8081)</li>
    <li>service-b ("Hello from service-b" on port 8082)</li>
  </ul>
</body>
</html>
`
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = fmt.Fprint(w, page)
	})

	addr := ":8080"
	log.Printf("frontend listening on %s\n", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("frontend failed: %v", err)
	}
}
