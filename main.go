package main

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/Masterminds/sprig"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

//go:embed templates/*.html
var templatesFS embed.FS

func renderTemplate(name string, files []string, data map[string]any) (string, error) {
	tmpl := template.New(name).Funcs(sprig.FuncMap())
	tmpl, err := tmpl.ParseFiles(files...)
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %v", err)
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("failed to execute template: %v", err)
	}
	//fmt.Println(buf.String())
	return buf.String(), nil
}

var templates = template.Must(template.ParseFS(templatesFS, "templates/*.html"))

func main() {
	app := pocketbase.New()

	// prints "Hello!" every 2 minutes
	app.Cron().MustAdd("hello", "*/2 * * * *", func() {
		log.Println("Hello!")
	})

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {

		// serves static files from the provided public dir (if exists)
		se.Router.GET("/public/{path...}", apis.Static(os.DirFS("./pb_public"), false))
		se.Router.GET("/static/{path...}", apis.Static(os.DirFS("./pb_public"), false))

		se.Router.GET("/", func(e *core.RequestEvent) error {
			// fmt.Println("REROUTED!")
			tmplData := map[string]any{
				"Title":   "Welcome to the Store",
				"Message": "Traditional goodies, handcrafted for you!",
			}
			html, err := renderTemplate("store", []string{
				"views/store/parts/header.html",
				"views/store/parts/footer.html",
				"views/store/layout.html",
				"views/store/products.html",
			}, tmplData)
			if err != nil {
				// or redirect to a dedicated 404 HTML page
				return e.NotFoundError("", err)
			}
			return e.HTML(http.StatusOK, html)
		})

		se.Router.GET("/product/:slug", func(e *core.RequestEvent) error {
			slug := e.Request.PathValue("slug")
			//db := app.Dao()
			records, err := app.FindRecordsByFilter(
				"products",
				"slug = {:slug}",
				"-slug",
				1,
				0,
				dbx.Params{"slug": slug},
			)
			if err != nil || len(records) == 0 {
				return e.String(404, "Product not found")
			}
			return templates.ExecuteTemplate(e.Response, "product.html", map[string]any{
				"Product": records[0],
			})
		})

		se.Router.GET("/cart", func(e *core.RequestEvent) error {
			return templates.ExecuteTemplate(e.Response, "cart.html", nil)
		})

		se.Router.GET("/checkout", func(e *core.RequestEvent) error {
			return templates.ExecuteTemplate(e.Response, "checkout.html", nil)
		})

		se.Router.POST("/order", func(e *core.RequestEvent) error {
			return e.String(http.StatusOK, "Order placed!")
		})

		se.Router.GET("/order/status/:id", func(e *core.RequestEvent) error {
			return templates.ExecuteTemplate(e.Response, "status.html", map[string]string{"Status": "Shipped"})
		})

		se.Router.GET("/login", func(e *core.RequestEvent) error {
			return templates.ExecuteTemplate(e.Response, "login.html", nil)
		})

		se.Router.GET("/register", func(e *core.RequestEvent) error {
			return templates.ExecuteTemplate(e.Response, "register.html", nil)
		})

		return se.Next()
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
