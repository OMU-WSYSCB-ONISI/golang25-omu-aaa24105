package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	fortunes := []string{"大吉", "中吉", "吉", "凶"}

	http.HandleFunc("/webfortune", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		f := fortunes[rand.Intn(len(fortunes))]
		fmt.Fprintf(w, "今の運勢は%sです\n", f)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, `
<!doctype html><meta charset="utf-8">
<title>Webおみくじ</title>
<h1>Webおみくじ</h1>
<p><a href="/webfortune">/webfortune にアクセス</a> で運勢を表示します。</p>
`)
	})

	addr := ":8080"
	log.Printf("listening on %s ...", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

