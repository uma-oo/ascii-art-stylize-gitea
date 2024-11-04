package handler

import (
	"fmt"
	"net/http"
	"os"
	"strings"

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
	fmt.Println(r.URL.Path)
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



func HandleAssets(w http.ResponseWriter, r *http.Request) {
	if !requestMethodChecker(w, r, http.MethodGet) {
		return
	} 
	if !strings.HasPrefix(r.URL.Path, "/assets") {
		handleStatusCode(w, http.StatusNotFound)
		return
	} else {
		file_info, err := os.Stat(r.URL.Path[1:])
		if err != nil {
			handleStatusCode(w, http.StatusNotFound)
			return
		} else 
		if file_info.IsDir() {
			handleStatusCode(w, http.StatusForbidden)
			return 
		} else {
			http.ServeFile(w, r, r.URL.Path[1:])
		}
	}
}
