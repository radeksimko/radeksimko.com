package app

import (
	"net/http"
)

func init() {
	http.HandleFunc("/robots.txt", plainText)
	http.HandleFunc("/pinterest-24f5e.html", plainText)
	http.HandleFunc("/attachments/cv-radek-simko.pdf", linkedIn)
	http.HandleFunc("/.*", mainPage)
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	h := http.RedirectHandler("http://about.radeksimko.com", 302)
	h.ServeHTTP(w, r)
}

func linkedIn(w http.ResponseWriter, r *http.Request) {
	h := http.RedirectHandler("http://linkedin.com/in/radeksimko", 302)
	h.ServeHTTP(w, r)
}

func plainText(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(200)
}
