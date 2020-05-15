package api

import (
	"html/template"
	"net/http"
	"path"
	"log"
)

// Index godoc
// @Summary Index
// @Description renders podinfo UI
// @Tags HTTP API
// @Produce html
// @Router / [get]
// @Success 200 {string} string "OK"
func (s *Server) indexHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Scheme: %s", r.URL.Scheme)
	
	if r.URL.Scheme != "https" {
		target := "https://" + r.Host + r.URL.Path

	    if len(r.URL.RawQuery) > 0 {
	        target += "?" + r.URL.RawQuery
	    }

	    log.Printf("redirect to: %s", target)

		http.Redirect(w, r, target , http.StatusMovedPermanently)
	}

	log.Printf("no redirect: %s", r.URL.Scheme)


	tmpl, err := template.New("vue.html").ParseFiles(path.Join(s.config.UIPath, "vue.html"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(path.Join(s.config.UIPath, "vue.html") + err.Error()))
		return
	}

	data := struct {
		Title string
		Logo  string
	}{
		Title: s.config.Hostname,
		Logo:  s.config.UILogo,
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, path.Join(s.config.UIPath, "vue.html")+err.Error(), http.StatusInternalServerError)
	}
}
