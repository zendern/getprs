package renderer

import "github.com/zendern/getprs/models"

type Renderer = func(statuses []models.PRStatus)

var Renderers = map[string]Renderer{
	"table": RenderTable,
	"json":  RenderJson,
	"text":  RenderText,
}
