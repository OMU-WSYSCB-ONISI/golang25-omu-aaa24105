package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	fmt.Println("Week 06 課題（BMI）")

	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	http.HandleFunc("/bmi", bmihandler)

	fmt.Println("http://localhost:8080 で起動中（Week06 BMI）")
	http.ListenAndServe(":8080", nil)
}

func bmihandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintln(w, "エラー")
		return
	}

	weight, _ := strconv.ParseFloat(r.FormValue("weight"), 64)
	height, _ := strconv.ParseFloat(r.FormValue("height"), 64)

	if height == 0 {
		fmt.Fprintln(w, "身長が0では計算できません")
		return
	}

	hm := height / 100
	bmi := weight / (hm * hm)

	fmt.Fprintf(w, "BMI：%.2f", bmi)
}




