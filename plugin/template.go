package plugin

var (
	tmpl = `
package main

import (
	"github.com/itzmanish/go-micro/v2/plugin"

	"{{.Path}}"
)

var Plugin = plugin.Config{
	Name: "{{.Name}}",
	Type: "{{.Type}}",
	Path: "{{.Path}}",
	NewFunc: {{.Name}}.{{.NewFunc}},
}
`
)
