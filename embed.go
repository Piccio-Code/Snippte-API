package FirstAPI

import (
	"embed"
	_ "embed"
	"io/fs"
)

//go:embed ui/html
var templatesFS embed.FS

var TemplatesFolder, _ = fs.Sub(templatesFS, "ui/html")

//go:embed ui/static
var staticFS embed.FS

var StaticFolder, _ = fs.Sub(staticFS, "ui/static")
