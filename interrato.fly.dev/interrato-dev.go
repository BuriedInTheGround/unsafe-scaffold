package main

import (
	"embed"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"net/url"
)

//go:embed interrato.dev
var interratoDevContent embed.FS

var goGetTmpl = template.Must(template.New("go-get").Parse(`
{{ $repo := or .GitRepo (printf "https://github.com/BuriedInTheGround/%s" .Name) }}
<!DOCTYPE html>
<html lang="en">
<head>
    <meta name="go-import" content="interrato.dev/{{ .Name }} git {{ $repo }}">
    <meta http-equiv="refresh" content="0;url={{ or .Redirect $repo }}">
</head>
<body>
    You will be redirected to the <a href="{{ or .Redirect $repo }}">project page</a>...
</body>
</html>
`))

type goGetHandler struct {
	Name     string
	GitRepo  string
	Redirect string
}

func (h goGetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	goGetTmpl.Execute(w, h)
}

func interratoDev(mux *http.ServeMux) {
	mux.HandleFunc("www.interrato.dev/", func(w http.ResponseWriter, r *http.Request) {
		u := &url.URL{
			Scheme:   "https",
			Host:     "interrato.dev",
			Path:     r.URL.Path,
			RawQuery: r.URL.RawQuery,
		}
		http.Redirect(w, r, u.String(), http.StatusMovedPermanently)
	})

	content, err := fs.Sub(interratoDevContent, "interrato.dev")
	if err != nil {
		log.Fatal(err)
	}
	mux.Handle("interrato.dev/static/", http.FileServer(http.FS(content)))

	for _, page := range []struct {
		Title        string
		Description  string
		Path         string
		TemplateName string
	}{
		{
			Title:        "Home",
			Description:  "Computer Engineer, Cybersecurity at the University of Padua.",
			Path:         "/",
			TemplateName: "home.html",
		},
		{
			Title:        "CV",
			Description:  "Simone Ragusa's Curriculum Vitae",
			Path:         "/cv/",
			TemplateName: "cv.html",
		},
	} {
		page := page
		mux.HandleFunc("interrato.dev"+page.Path, func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != page.Path {
				http.NotFound(w, r)
				return
			}
			tmpl, err := template.ParseFS(content,
				"templates/base.html", "templates/"+page.TemplateName)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			err = tmpl.Execute(w, page)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		})
	}

	// go get handlers
	mux.Handle("interrato.dev/carbonize/", goGetHandler{
		Name: "carbonize",
	})
	mux.Handle("interrato.dev/fine/", goGetHandler{
		Name: "fine",
	})
	mux.Handle("interrato.dev/olaf/", goGetHandler{
		Name: "olaf",
	})
	mux.Handle("interrato.dev/pigowa/", goGetHandler{
		Name: "pigowa",
	})

	// redirects
	for path, url := range map[string]string{
		"/resume/": "https://interrato.dev/cv/",
	} {
		path, url := path, url
		mux.HandleFunc("interrato.dev"+path, func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, url, http.StatusFound)
		})
	}
}
