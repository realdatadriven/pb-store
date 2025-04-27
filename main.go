package main

import (
	"bytes"
	"embed"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/Masterminds/sprig"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/jsvm"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
	"github.com/pocketbase/pocketbase/tools/hook"
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

	// ---------------------------------------------------------------
	// Optional plugin flags:
	// ---------------------------------------------------------------

	var hooksDir string
	app.RootCmd.PersistentFlags().StringVar(
		&hooksDir,
		"hooksDir",
		"",
		"the directory with the JS app hooks",
	)

	var hooksWatch bool
	app.RootCmd.PersistentFlags().BoolVar(
		&hooksWatch,
		"hooksWatch",
		true,
		"auto restart the app on pb_hooks file change; it has no effect on Windows",
	)

	var hooksPool int
	app.RootCmd.PersistentFlags().IntVar(
		&hooksPool,
		"hooksPool",
		15,
		"the total prewarm goja.Runtime instances for the JS app hooks execution",
	)

	var migrationsDir string
	app.RootCmd.PersistentFlags().StringVar(
		&migrationsDir,
		"migrationsDir",
		"",
		"the directory with the user defined migrations",
	)

	var automigrate bool
	app.RootCmd.PersistentFlags().BoolVar(
		&automigrate,
		"automigrate",
		true,
		"enable/disable auto migrations",
	)

	var publicDir string
	app.RootCmd.PersistentFlags().StringVar(
		&publicDir,
		"publicDir",
		defaultPublicDir(),
		"the directory to serve static files",
	)

	var indexFallback bool
	app.RootCmd.PersistentFlags().BoolVar(
		&indexFallback,
		"indexFallback",
		true,
		"fallback the request to index.html on missing static path, e.g. when pretty urls are used with SPA",
	)

	app.RootCmd.ParseFlags(os.Args[1:])

	// ---------------------------------------------------------------
	// Plugins and hooks:
	// ---------------------------------------------------------------

	// load jsvm (pb_hooks and pb_migrations)
	jsvm.MustRegister(app, jsvm.Config{
		MigrationsDir: migrationsDir,
		HooksDir:      hooksDir,
		HooksWatch:    hooksWatch,
		HooksPoolSize: hooksPool,
	})

	// migrate command (with js templates)
	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		TemplateLang: migratecmd.TemplateLangJS,
		Automigrate:  automigrate,
		Dir:          migrationsDir,
	})

	// GitHub selfupdate
	//ghupdate.MustRegister(app, app.RootCmd, ghupdate.Config{})

	// prints "Hello!" every 2 minutes
	app.Cron().MustAdd("hello", "*/2 * * * *", func() {
		log.Println("Hello!")
	})

	// static route to serves files from the provided public dir
	// (if publicDir exists and the route path is not already defined)
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

		se.Router.GET("/auth/{_type}", func(e *core.RequestEvent) error {
			_type := e.Request.PathValue("_type")
			fmt.Println(_type)
			tmplData := map[string]any{"Title": "Login"}
			if _type == "" || _type == "login" {
				_type = "login"
				tmplData = map[string]any{"Title": "Login"}
			} else if _type == "recovery" {
				tmplData = map[string]any{"Title": "Recover Password"}
			} else if _type == "signup" {
				tmplData = map[string]any{"Title": "Signup"}
			}
			fmt.Println(fmt.Sprintf("templates/auth/%s.html", _type))
			html, err := renderTemplate("auth", []string{
				fmt.Sprintf("templates/auth/%s.html", _type),
				"templates/auth/layout.html",
			}, tmplData)
			if err != nil {
				return e.NotFoundError("", err)
			}
			return e.HTML(http.StatusOK, html)
		})

		se.Router.POST("/auth/{_type}", func(e *core.RequestEvent) error {
			_type := e.Request.PathValue("_type")
			fmt.Println(_type)
			tmplData := map[string]any{"Title": "Login"}
			if _type == "" || _type == "login" {
				_type = "login"
				email := e.Request.FormValue("email")
				password := e.Request.FormValue("password")
				// fmt.Println(email, password)
				record, err := app.FindAuthRecordByEmail("users", email)
				if err != nil {
					fmt.Printf("%s", err)
					tmplData = map[string]any{
						"Title": "Login",
						"msg":   "User or password incorrect!",
					}
				}
				if record.ValidatePassword(password) {
					token, err := record.NewAuthToken()
					if err != nil {
						fmt.Printf("%s", err)
						tmplData = map[string]any{
							"Title": "Login",
							"msg":   "User or password incorrect!",
						}
					} else {
						fmt.Println(record, token)
						_val, _ := json.Marshal(map[string]any{"token": token, "record": record})
						http.SetCookie(e.Response, &http.Cookie{
							Name:  "session",
							Value: string(_val),
							Path:  "/",
						})
						//http.Redirect(e.Response, e.Request, "/", http.StatusSeeOther)
					}
				} else {
					fmt.Printf("Password not valid!")
					tmplData = map[string]any{
						"Title": "Login",
						"msg":   "User or password incorrect!",
					}
				}

				/*http.SetCookie(w, &http.Cookie{
					Name:  "session",
					Value: "authenticated",
					Path:  "/",
				})
				http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
				*/
			} else if _type == "recovery" {
				tmplData = map[string]any{"Title": "Recover Password"}
			} else if _type == "signup" {
				tmplData = map[string]any{"Title": "Signup"}
			}
			fmt.Println(fmt.Sprintf("templates/auth/%s.html", _type))
			html, err := renderTemplate("auth", []string{
				fmt.Sprintf("templates/auth/%s.html", _type),
				"templates/auth/layout.html",
			}, tmplData)
			if err != nil {
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
	app.OnServe().Bind(&hook.Handler[*core.ServeEvent]{
		Func: func(se *core.ServeEvent) error {
			if !se.Router.HasRoute(http.MethodGet, "/{path...}") {
				//se.Router.GET("/{path...}", apis.Static(os.DirFS(publicDir), indexFallback))
			}
			return se.Next()
		},
		Priority: 999, // execute as latest as possible to allow users to provide their own route
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}

// the default pb_public dir location is relative to the executable
func defaultPublicDir() string {
	if strings.HasPrefix(os.Args[0], os.TempDir()) {
		// most likely ran with go run
		return "./pb_public"
	}

	return filepath.Join(os.Args[0], "../pb_public")
}
