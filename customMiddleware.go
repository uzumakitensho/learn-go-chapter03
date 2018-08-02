package main

import (
	"fmt"
	"net/http"
)

func middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Executing middleware before request phase!")
		handler.ServeHTTP(w, r)
		fmt.Println("Executing middleware after response phase!")
	})
}

func mainLogic(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Executing main handler..")
	w.Write([]byte("OK"))
}

func main() {
	mailLogicHandler := http.HandlerFunc(mainLogic)
	http.Handle("/", middleware(mailLogicHandler))
	http.ListenAndServe(":8008", nil)
}
