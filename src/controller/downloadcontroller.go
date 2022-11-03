package controller

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fileUrl := "https://www.africau.edu/images/default/sample.pdf"
	filename, err := GetFilename(fileUrl)
	fmt.Println(filename)
	if err != nil {
		http.Error(w, "error url", 502)
		return
	}
	resp, err := http.Get(fileUrl)
	if err != nil {
		http.Error(w, "error url", 502)
		return
	}
	defer resp.Body.Close()
	w.Header().Set("Content-Disposition", "attachment; filename= "+filename)
	w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
	w.Header().Set("Content-Lenght", resp.Header.Get("Content-Lenght"))
	_, err = io.Copy(w, resp.Body)
	if err != nil {
		http.Error(w, "Remote server error", 503)
		return
	}
}

func GetFilename(inputUrl string) (string, error) {
	u, err := url.Parse(inputUrl)
	if err != nil {
		return "", err
	}
	u.RawQuery = ""
	return path.Base(u.String()), nil
}
