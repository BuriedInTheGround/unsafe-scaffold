package main

import (
	"net/http"
	"net/url"
)

func perpetuatheme(mux *http.ServeMux) {
	mux.HandleFunc("www.perpetuatheme.com/", func(w http.ResponseWriter, r *http.Request) {
		u := &url.URL{
			Scheme:   "https",
			Host:     "perpetuatheme.com",
			Path:     r.URL.Path,
			RawQuery: r.URL.RawQuery,
		}
		http.Redirect(w, r, u.String(), http.StatusMovedPermanently)
	})
	mux.Handle("perpetuatheme.com/", http.RedirectHandler(
		"https://github.com/perpetuatheme", http.StatusFound))
}
