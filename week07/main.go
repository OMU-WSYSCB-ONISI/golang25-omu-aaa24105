package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	fmt.Println("Week 07 課題（加減乗除）")

	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	http.HandleFunc("/cal02", calpmjhHandler)

	fmt.Println("http://localhost:8080 で起動中（Week07 加減乗除）")
	http.ListenAndServe(":8080", nil)
}

func calpmjhHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	x, _ := strconv.ParseFloat(r.FormValue("x"), 64)
	y, _ := strconv.ParseFloat(r.FormValue("y"), 64)
	op := r.FormValue("cal0")

	switch op {
	case "+":
		fmt.Fprintln(w, x+y)
	case "-":
		fmt.Fprintln(w, x-y)
	case "*":
		fmt.Fprintln(w, x*y)
	case "/":
		if y == 0 {
			fmt.Fprintln(w, "0で割ることはできません")
			return
		}
		fmt.Fprintln(w, x/y)
	}
}
