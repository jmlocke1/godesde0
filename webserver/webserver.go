package webserver

import "net/http"

func MiWebServer() {
	http.HandleFunc("/", home)
	http.HandleFunc("/login", login)
	http.ListenAndServe(":3000", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./webserver/index.html")
}

func login(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./webserver/login.html")
}
