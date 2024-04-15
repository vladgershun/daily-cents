package main

import (
	"html/template"
	"io/fs"
	"net/http"
	"path/filepath"

	"github.com/vladgershun/daily-cents/internal/models"
	"github.com/vladgershun/daily-cents/ui"
)

type templateData struct {
	Banks []models.Bank
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := fs.Glob(ui.Files, "html/pages/*.tmpl")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		patterns := []string{
			"html/base.tmpl",
			// "html/partials/*.tmpl",
			page,
		}

		ts, err := template.New(name).ParseFS(ui.Files, patterns...)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}

func (app *application) newTemplateData(r *http.Request) templateData {
	return templateData{}
}
