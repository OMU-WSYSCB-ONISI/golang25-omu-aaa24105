package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")

		now := time.Now().Format("15:04")

		ua := r.Header.Get("User-Agent")
		if ua == "" {
			ua = "不明"
		}

		fmt.Fprintf(w, "今の時刻は %s で，利用しているブラウザは %s ，ですね。\n", now, ua)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, `
<!doctype html><meta charset="utf-8">
<title>info エンドポイント</title>
<h1>/info エンドポイント</h1>
<p><a href="/info">/info にアクセス</a> すると現在時刻とUser-Agentを返します。</p>
`)
	})

	addr := ":8080"
	log.Printf("listening on %s ...", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
