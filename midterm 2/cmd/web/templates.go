package main

import (
	"html/template"
	"path/filepath"
	"website.maksat.com/internal/models"
)

// List of all products in order to display them in the home page
type BaseTemplate struct {
	User     *models.User
	Product  *models.Product
	Products []*models.Product
}

// caching the tamplates in order not to parse them each time from the files
func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}
	// Return a slice of all filepaths shown in the string
	pages, err := filepath.Glob("./ui/html/pages/*.html")
	if err != nil {
		return nil, err
	}
	// Iterate through that slice above to cache them out into the templates
	// That should create a convenient map for us
	for _, page := range pages {
		name := filepath.Base(page)
		files := []string{
			"./ui/html/base.html",
			page,
		}
		tmps, err := template.ParseFiles(files...)
		if err != nil {
			return nil, err
		}

		cache[name] = tmps

	}
	return cache, nil
}
