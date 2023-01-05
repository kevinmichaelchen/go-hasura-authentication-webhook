package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handle)
	http.ListenAndServe(":8085", nil)
}

func handle(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("authorization")
	fmt.Println("token", token)

	// Sleep so we can test HTTP caching.
	// If Hasura is caching these responses, this endpoint shouldn't even be called.
	time.Sleep(5 * time.Second)

	if !isAllowed(token) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	res := Response{
		UserID: "foobar",
		Role:   "boomi",
		Custom: "",
	}

	b, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}

	w.Header().Set("cache-control", "max-age=604800")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func isAllowed(token string) bool {
	return true
}

type Response struct {
	UserID string `json:"X-Hasura-User-Id"`
	Role   string `json:"X-Hasura-Role"`
	Custom string `json:"X-Hasura-Custom"`
}
