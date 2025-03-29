package main

import (
	"net/http"
	"path/filepath"
	"strings"
)

func filesystemMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		if shouldAddExtension(path) {
			path = appendHTMLExtension(path)
		}

		if strings.HasPrefix(path, "//index.html") {
			path = path[1:]
		}

		r.URL.Path = path

		next.ServeHTTP(w, r)
	})
}

func shouldAddExtension(path string) bool {
	return !strings.HasSuffix(path, ".html") && filepath.Ext(path) == ""
}

func appendHTMLExtension(path string) string {
	if strings.Contains(path, "?") {
		parts := strings.SplitN(path, "?", 2)
		return parts[0] + "/index.html?" + parts[1]
	}
	return path + "/index.html"
}
