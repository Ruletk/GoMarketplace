package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/register", handleRegister)
	http.HandleFunc("/password-reset", handlePasswordReset)
	http.HandleFunc("/password-change", handlePasswordChange)

	http.ListenAndServe(":8080", nil)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello from auth service in login route")
	fmt.Println("Hello from auth service in login route")
}

func handleRegister(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello from auth service in register route")
	fmt.Println("Hello from auth service in register route")
}

func handlePasswordReset(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello from auth service in password reset route")
	fmt.Println("Hello from auth service in password reset route")
}

func handlePasswordChange(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello from auth service in password change route")
	fmt.Println("Hello from auth service in password change route")
}
