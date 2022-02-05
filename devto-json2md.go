// devto-json2md converts dev.to JSON exports into Markdown files
// usage:
// devto-json2md articles.json <output directory>
package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"k8s.io/klog/v2"
)

type Article struct {
	CreatedAt      time.Time `json:"created_at"`
	BodyMarkdown   string    `json:"body_markdown"`
	Slug           string    `json:"slug"`
	Description    string    `json:"description"`
	Title          string    `json:"title"`
	Path           string    `json:"path"`
	CachedUserName string    `json:"cached_username"`
	MainImage      string    `json:"main_image"`
	Published      bool      `json:"published"`
}

// outputTemplate is based on Hugo's frontmatter spec: https://gohugo.io/content-management/front-matter/
var outputTemplate = `---
title: {{ .Title|escape }}
date: {{ .CreatedAt.Format "2006-01-02" }}
description: {""{ .Description|escape }}
slug: {{ .Slug }}
{{- if not .Published }}
draft: true
{{- end }}
{{- with .MainImage }}
main_image: {{.}}
{{- end }}
---
{{ .BodyMarkdown }}
`

// escapeYAML is a low-tech way of escaping YAML strings
func escapeYAML(s string) string {
	s = strings.ReplaceAll(s, `"`, `\"`)
	return `"` + s + `"`
}

func main() {
	klog.InitFlags(nil)

	fm := template.FuncMap{
		"escape": escapeYAML,
	}
	t, err := template.New("articles").Funcs(fm).Parse(outputTemplate)
	if err != nil {
		klog.Exitf("template failed to parse: %v", err)
	}

	if len(os.Args) != 3 {
		klog.Exitf("Usage: devto-json2md /path/to/articles.json /path/to/output/directory")
	}

	articles := []Article{}

	input := os.Args[1]
	outDir := os.Args[2]
	bs, err := os.ReadFile(input)
	if err != nil {
		klog.Exitf("unable to read input: %v", err)
	}

	err = json.Unmarshal(bs, &articles)
	if err != nil {
		klog.Exitf("unmarshal: %w", err)
	}

	for i, a := range articles {
		path := filepath.Join(outDir, a.Slug) + ".md"
		klog.Infof("#%d: Writing out %s ...", i+1, path)

		f, err := os.Create(path)
		if err != nil {
			klog.Exitf("create failed for %s: %v", path, err)
		}

		err = t.Execute(f, a)
		if err != nil {
			klog.Exitf("execute failed: %v", err)
		}

	}
}
