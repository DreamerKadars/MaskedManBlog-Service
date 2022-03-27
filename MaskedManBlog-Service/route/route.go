package route

import (
	// "time"

	"MaskedManBlog-Service/service"
	"time"

	iris "github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/monitor"
)

var App *iris.Application

func init() {
	App = iris.New()
	// Initialize and start the monitor middleware.
	m := monitor.New(monitor.Options{
		RefreshInterval:     2 * time.Second,
		ViewRefreshInterval: 2 * time.Second,
		ViewTitle:           "MyServer Monitor",
	})
	// Manually stop monitoring on CMD/CTRL+C.
	iris.RegisterOnInterrupt(m.Stop)

	App.Post("/monitor", m.Stats)
	App.Get("/monitor", m.View)
	App.Get("/", service.Handler)

}