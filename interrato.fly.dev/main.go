package main

import (
	"log"
	"net/http"
	"net/url"
	"time"
)

func main() {
	mux := http.NewServeMux()
	perpetuatheme(mux)
	interratoDev(mux)

	s := http.Server{
		Addr: ":8080",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Header.Get("X-Forwarded-Proto") == "http" {
				u := &url.URL{
					Scheme:   "https",
					Host:     r.URL.Host,
					Path:     r.URL.Path,
					RawQuery: r.URL.RawQuery,
				}
				http.Redirect(w, r, u.String(), http.StatusMovedPermanently)
			}
			w.Header().Set("Strict-Transport-Security", "max-age=63072000; includeSubDomains; preload")
			w.Header().Set("Referrer-Policy", "no-referrer, strict-origin-when-cross-origin")
			w.Header().Set("X-Content-Type-Options", "nosniff")
			w.Header().Set("X-Frame-Options", "DENY")
			mux.ServeHTTP(w, r)
		}),
		ReadTimeout:  1 * time.Minute,
		WriteTimeout: 1 * time.Minute,
	}

	log.Fatal(s.ListenAndServe())
}
