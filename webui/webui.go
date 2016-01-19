package webui

import "fmt"
import "net/http"

func init() {
	http.HandleFunc("/", rootHandler)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header()["Content-Type"] = []string{"text/plain"}
	fmt.Fprint(w, "Hello, world!")
}
