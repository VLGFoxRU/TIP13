package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof" // регистрирует /debug/pprof/*
	"example.com/pz13-pprof-lab/internal/work"
)

func main() {
	// Эндпоинт, вызывающий “тяжёлую” работу.
	http.HandleFunc("/work", func(w http.ResponseWriter, r *http.Request) {
		defer work.TimeIt("Fib(38)")()
		n := 38 // достаточно тяжело для CPU
		res := work.FibFast(n)
		w.Header().Set("Content-Type", "text/plain")
		_, _ = w.Write([]byte((fmtInt(res))))
	})

	log.Println("Server on :8080; pprof on /debug/pprof/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func fmtInt(v int) string { return fmt.Sprintf("%d\n", v) }
