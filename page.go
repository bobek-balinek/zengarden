package zengarden

import (
	"path/filepath"
)

// Page is a single non-post page not in the _posts directory.
type Page struct {
	site *Site
	vars Context
}

func (p *Page) toURL() string {
	from := p.vars["path"].(string)
	return joinURL(p.site.cfg.BaseURL, filepath.ToSlash(from[len(p.site.cfg.Source):]))
}

func (p *Page) toPath() string {
	from := p.vars["path"].(string)
	return filepath.ToSlash(filepath.Join(p.site.cfg.Target, from[len(p.site.cfg.Source):]))
}

// Pages is a collection of non-post pages.
type Pages []*Page

func (p Pages) context() []Context {
	ctx := make([]Context, 0, len(p))

	for _, page := range p {
		ctx = append(ctx, page.vars)
	}

	return ctx
}

func (p Pages) convert(siteVars Context) error {
	for _, page := range p {
		src := page.vars["path"].(string)

		if err := convertFile(src, page.toPath(), page.toURL(), page.site, nil); err != nil {
			return err
		}
	}

	return nil
}

func (p Pages) findIndex() *Page {
	for _, page := range p {
		if page.vars["url"] == "/index.md" || page.vars["url"] == "/index.html" {
			return page
		}
	}

	return nil
}

func (p Pages) Len() int {
	return len(p)
}

func (p Pages) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Pages) Less(i, j int) bool {
	return p[i].toPath() < p[j].toPath()
}
