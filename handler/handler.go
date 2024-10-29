package handler

import (
	"net/http"

	"asciiWeb/internal"
)

type Data struct {
	Text      string
	Banner    string
	FormError string
	AsciiArt  string
}

var Pagedata = Data{}

// handler for the path "/"
func HandleMainPage(w http.ResponseWriter, r *http.Request) {
	Pagedata = Data{}
	if r.URL.Path != `/` {
		handleStatusCode(w, http.StatusNotFound)
		return
	}
	if !requestMethodChecker(w, r, http.MethodGet) {
		return
	}
	renderTemplate(w, "index.html", Pagedata)
}

// handler for the path "/ascii-art
func HandleAsciiArt(w http.ResponseWriter, r *http.Request) {
	if !requestMethodChecker(w, r, http.MethodPost) {
		return
	}

	if status := extractFormData(r); status != 200 {
		handleStatusCode(w, status)
		return
	}

	if status := validateFormData(); status != 200 {
		handleStatusCode(w, status)
		return
	}

	asciiArt, status := internal.Ascii(Pagedata.Text, Pagedata.Banner)
	if status != 200 {
		handleStatusCode(w, status)
		return
	}

	Pagedata.AsciiArt = asciiArt
	renderTemplate(w, "index.html", Pagedata)
	Pagedata = Data{}
}
