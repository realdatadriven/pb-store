package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/jsvm"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
	_ "github.com/realdatadriven/pocket_store/migrations"
)

// the default pb_public dir location is relative to the executable
func defaultPublicDir() string {
	if strings.HasPrefix(os.Args[0], os.TempDir()) {
		// most likely ran with go run
		return "./pb_public"
	}

	return filepath.Join(os.Args[0], "../pb_public")
}

func main() {
	// Load .env file
	_err := godotenv.Load()
	if _err != nil {
		fmt.Print("Error loading .env file")
	}

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

	// Template rendering function
	renderTemplate := func(name string, files []string, data map[string]any) (string, error) {
		tmpl, err := template.ParseFiles(files...)
		if err != nil {
			return "", err
		}

		var buf strings.Builder
		err = tmpl.ExecuteTemplate(&buf, name, data)
		if err != nil {
			return "", err
		}

		return buf.String(), nil
	}

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		// Serve static files
		se.Router.GET("/static/{path...}", func(e *core.RequestEvent) error {
			path := e.Request.PathValue("path")
			filePath := filepath.Join("pb_public", path)
			http.ServeFile(e.Response, e.Request, filePath)
			return nil
		})

		// Home route
		se.Router.GET("/", func(e *core.RequestEvent) error {
			// Fetch products
			products, err := app.FindRecordsByFilter(
				"products",
				"status = 'published'",
				"",
				12,
				0,
			)
			if err != nil {
				fmt.Printf("Error fetching products: %v\n", err)
				products = []*core.Record{}
			}

			// Fetch categories
			categories, err := app.FindRecordsByFilter(
				"categories",
				"isActive = true",
				"",
				10,
				0,
			)
			if err != nil {
				fmt.Printf("Error fetching categories: %v\n", err)
				categories = []*core.Record{}
			}

			tmplData := map[string]any{
				"Title":      "Pocket Store",
				"Message":    "Quality Products for Everyone",
				"Products":   products,
				"Categories": categories,
			}
			html, err := renderTemplate("store", []string{
				"views/store/parts/header.html",
				"views/store/parts/footer.html",
				"views/store/layout.html",
				"views/store/products.html",
			}, tmplData)
			if err != nil {
				fmt.Printf("Template rendering error: %v\n", err)
				return e.NotFoundError("", err)
			}
			return e.HTML(http.StatusOK, html)
		})

		// Product detail route
		se.Router.GET("/products/{slug}", func(e *core.RequestEvent) error {
			slug := e.Request.PathValue("slug")

			// Try to find product by slug first, then by ID
			var product *core.Record
			var err error

			// First try to find by slug
			products, err := app.FindRecordsByFilter(
				"products",
				fmt.Sprintf("slug = '%s' && status = 'published'", slug),
				"",
				1,
				0,
			)
			if err == nil && len(products) > 0 {
				product = products[0]
			} else {
				// If not found by slug, try by ID
				product, err = app.FindRecordById("products", slug)
				if err != nil {
					return e.NotFoundError("Product not found", err)
				}
				// Check if product is published
				if product.GetString("status") != "published" {
					return e.NotFoundError("Product not found", nil)
				}
			}

			tmplData := map[string]any{
				"Title":   "Pocket Store",
				"Product": product,
			}
			html, err := renderTemplate("store", []string{
				"views/store/parts/header.html",
				"views/store/parts/footer.html",
				"views/store/layout.html",
				"views/store/product_detail.html",
			}, tmplData)
			if err != nil {
				fmt.Printf("Template rendering error: %v\n", err)
				return e.NotFoundError("", err)
			}
			return e.HTML(http.StatusOK, html)
		})

		// Search route
		se.Router.GET("/search", func(e *core.RequestEvent) error {
			query := e.Request.URL.Query().Get("q")

			var products []*core.Record
			var err error

			if query != "" {
				// Search in title and description
				filter := fmt.Sprintf("(title ~ '%s' || description ~ '%s') && status = 'published'", query, query)
				products, err = app.FindRecordsByFilter(
					"products",
					filter,
					"",
					30,
					0,
				)
			} else {
				// If no query, show all products
				products, err = app.FindRecordsByFilter(
					"products",
					"status = 'published'",
					"",
					30,
					0,
				)
			}

			if err != nil {
				fmt.Printf("Error searching products: %v\n", err)
				products = []*core.Record{}
			}

			// Fetch categories for navigation
			categories, err := app.FindRecordsByFilter(
				"categories",
				"isActive = true",
				"",
				10,
				0,
			)
			if err != nil {
				fmt.Printf("Error fetching categories: %v\n", err)
				categories = []*core.Record{}
			}

			title := "Search Results"
			message := fmt.Sprintf("Found %d products", len(products))
			if query != "" {
				title = fmt.Sprintf("Search Results for \"%s\"", query)
			}

			tmplData := map[string]any{
				"Title":       "Pocket Store",
				"Message":     message,
				"Products":    products,
				"Categories":  categories,
				"SearchQuery": query,
				"PageTitle":   title,
			}
			html, err := renderTemplate("store", []string{
				"views/store/parts/header.html",
				"views/store/parts/footer.html",
				"views/store/layout.html",
				"views/store/products.html",
			}, tmplData)
			if err != nil {
				fmt.Printf("Template rendering error: %v\n", err)
				return e.NotFoundError("", err)
			}
			return e.HTML(http.StatusOK, html)
		})

		// Category filtering route
		se.Router.GET("/category/{slug}", func(e *core.RequestEvent) error {
			slug := e.Request.PathValue("slug")

			// Find category by slug
			categories, err := app.FindRecordsByFilter(
				"categories",
				fmt.Sprintf("slug = '%s' && isActive = true", slug),
				"",
				1,
				0,
			)
			if err != nil || len(categories) == 0 {
				return e.NotFoundError("Category not found", err)
			}

			category := categories[0]

			// Find products in this category
			products, err := app.FindRecordsByFilter(
				"products",
				fmt.Sprintf("categoryId = '%s' && status = 'published'", category.Id),
				"",
				30,
				0,
			)
			if err != nil {
				fmt.Printf("Error fetching products for category: %v\n", err)
				products = []*core.Record{}
			}

			// Fetch all categories for navigation
			allCategories, err := app.FindRecordsByFilter(
				"categories",
				"isActive = true",
				"",
				10,
				0,
			)
			if err != nil {
				fmt.Printf("Error fetching categories: %v\n", err)
				allCategories = []*core.Record{}
			}

			title := fmt.Sprintf("%s - Products", category.GetString("name"))
			message := fmt.Sprintf("Found %d products in %s", len(products), category.GetString("name"))

			tmplData := map[string]any{
				"Title":           "Pocket Store",
				"Message":         message,
				"Products":        products,
				"Categories":      allCategories,
				"CurrentCategory": category,
				"PageTitle":       title,
			}
			html, err := renderTemplate("store", []string{
				"views/store/parts/header.html",
				"views/store/parts/footer.html",
				"views/store/layout.html",
				"views/store/products.html",
			}, tmplData)
			if err != nil {
				fmt.Printf("Template rendering error: %v\n", err)
				return e.NotFoundError("", err)
			}
			return e.HTML(http.StatusOK, html)
		})

		// Cart route
		se.Router.GET("/cart", func(e *core.RequestEvent) error {
			tmplData := map[string]any{
				"Title": "Pocket Store",
			}
			html, err := renderTemplate("store", []string{
				"views/store/parts/header.html",
				"views/store/parts/footer.html",
				"views/store/layout.html",
				"views/store/cart.html",
			}, tmplData)
			if err != nil {
				fmt.Printf("Template rendering error: %v\n", err)
				return e.NotFoundError("", err)
			}
			return e.HTML(http.StatusOK, html)
		})

		// Checkout route
		se.Router.GET("/checkout", func(e *core.RequestEvent) error {
			tmplData := map[string]any{
				"Title": "Pocket Store",
			}
			html, err := renderTemplate("store", []string{
				"views/store/parts/header.html",
				"views/store/parts/footer.html",
				"views/store/layout.html",
				"views/store/checkout.html",
			}, tmplData)
			if err != nil {
				fmt.Printf("Template rendering error: %v\n", err)
				return e.NotFoundError("", err)
			}
			return e.HTML(http.StatusOK, html)
		})

		// Login route
		se.Router.GET("/login", func(e *core.RequestEvent) error {
			tmplData := map[string]any{
				"Title": "Pocket Store",
			}
			html, err := renderTemplate("store", []string{
				"views/store/parts/header.html",
				"views/store/parts/footer.html",
				"views/store/layout.html",
				"views/store/login.html",
			}, tmplData)
			if err != nil {
				fmt.Printf("Template rendering error: %v\n", err)
				return e.NotFoundError("", err)
			}
			return e.HTML(http.StatusOK, html)
		})

		// Register route
		se.Router.GET("/register", func(e *core.RequestEvent) error {
			tmplData := map[string]any{
				"Title": "Pocket Store",
			}
			html, err := renderTemplate("store", []string{
				"views/store/parts/header.html",
				"views/store/parts/footer.html",
				"views/store/layout.html",
				"views/store/register.html",
			}, tmplData)
			if err != nil {
				fmt.Printf("Template rendering error: %v\n", err)
				return e.NotFoundError("", err)
			}
			return e.HTML(http.StatusOK, html)
		})

		return se.Next()
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
