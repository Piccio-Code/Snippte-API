package FirstAPI

import (
	"embed"
	_ "embed"
)

//go:embed ui/html
var TemplatesFS embed.FS

//go:embed ui/static
var StaticsFS embed.FS
