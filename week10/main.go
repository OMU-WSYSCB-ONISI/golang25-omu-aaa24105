package main

import (
	"fmt"
	"html"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
)

const (
	dataDir  = "public"
	fileName = "memo.txt"
)

var saveFile = filepath.Join(dataDir, fileName)

func main() {
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		fmt.Printf("ディレクトリの作成に失敗しました: %v\n", err)
		return
	}

	fmt.Printf("Go version: %s\n", runtime.Version())
	fmt.Printf("Server is running at http://localhost:8080/memo\n")

	http.HandleFunc("/hello", hellohandler)
	http.HandleFunc("/memo", memo)
	http.HandleFunc("/mwrite", mwrite)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Failed to launch server: %v", err)
	}
}

func hellohandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "こんにちは from Codespace !")
}

func memo(w http.ResponseWriter, r *http.Request){
	text, err := os.ReadFile(saveFile)
	if err != nil {
		text = []byte("")
	}

	htmlText := html.EscapeString(string(text))

	htmlContent := fmt.Sprintf(`
<!DOCTYPE html>
<html lang="ja">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Go Memo App</title>
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css" rel="stylesheet">
</head>
<body class="bg-light">
    <div class="container mt-5">
        <div class="card shadow-sm">
            <div class="card-header bg-primary text-white">
                <h2 class="h5 mb-0">シンプルメモ帳</h2>
            </div>
            <div class="card-body">
                <form method="POST" action="/mwrite">
                    <div class="mb-3">
                        <textarea name="text" class="form-control" rows="10" placeholder="ここにメモを入力...">%s</textarea>
                    </div>
                    <div class="d-grid gap-2">
                        <button type="submit" class="btn btn-primary">保存する</button>
                    </div>
                </form>
            </div>
            <div class="card-footer text-muted text-end">
                Go Server is running
            </div>
        </div>
    </div>
</body>
</html>`, htmlText)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(htmlContent))
}

func mwrite(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Form parsing error", http.StatusInternalServerError)
		return
	}

	text := r.FormValue("text")

	if err := os.WriteFile(saveFile, []byte(text), 0644); err != nil {
		http.Error(w, "Failed to save file", http.StatusInternalServerError)
		fmt.Printf("Save error: %v\n", err)
		return
	}

	fmt.Println("Saved content length:", len(text))

	http.Redirect(w, r, "/memo", http.StatusSeeOther)
}
