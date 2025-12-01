package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Week 08 課題（平均と分布）")

	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	http.HandleFunc("/sum", sumhandler)

	fmt.Println("http://localhost:8080 で起動中（Week08）")
	http.ListenAndServe(":8080", nil)
}

func sumhandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		fmt.Fprintln(w, "フォーム読み込みエラー")
		return
	}

	raw := r.FormValue("dd")
	if raw == "" {
		fmt.Fprintln(w, "データが空です")
		return
	}

	replacer := strings.NewReplacer("　", "", "，", ",")
	raw = replacer.Replace(raw)
	raw = strings.ReplaceAll(raw, " ", "")

	arr := strings.Split(raw, ",")

	var sum, cnt int
	bunpu := make([]int, 11)

	for _, s := range arr {
		if s == "" {
			continue
		}


		v, err := strconv.Atoi(s)
		if err != nil {

			continue
		}


		if v < 0 || v > 100 {
			continue
		}


		index := v / 10
		bunpu[index]++

		sum += v
		cnt++
	}

	if cnt == 0 {
		fmt.Fprintln(w, "有効な数字が入力されていません")
		return
	}

	avg := float64(sum) / float64(cnt)

	fmt.Fprintf(w, "平均値：%.2f<br><br>", avg)
	fmt.Fprintln(w, "【分布】<br>")

	for i := 0; i < 10; i++ {
		fmt.Fprintf(w, "%2d〜%2d点：%d<br>", i*10, i*10+9, bunpu[i])
	}
	fmt.Fprintf(w, "100点：%d<br>", bunpu[10])
}
