package routes

import (
	"fmt"
	"net/http"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/router"
)

func Products(app *pocketbase.PocketBase, se *core.ServeEvent) *router.Route[*core.RequestEvent] {
	return se.Router.GET("/api/products", func(e *core.RequestEvent) error {
		info, err := e.RequestInfo()
		limit, ok := info.Body["limit"].(int)
		if !ok {
			limit = 10
		}
		offset, ok := info.Body["offset"].(int)
		if !ok {
			offset = 0
		}
		records, err := app.FindRecordsByFilter(
			"products",                        // collection
			"status = {:status}",              // filter
			"",                                // sort -published
			limit,                             // limit
			offset,                            // offset
			dbx.Params{"status": "published"}, // optional filter params
		)
		if err != nil {
			return e.NotFoundError("Missing or invalid slug", err)
		}
		response := map[string]any{
			"pageSize": limit,
			"page":     offset,
			"count":    len(records),
			"data":     records,
		}
		return e.JSON(http.StatusOK, response)
	})
}

func ProductsById(app *pocketbase.PocketBase, se *core.ServeEvent) *router.Route[*core.RequestEvent] {
	return se.Router.GET("/api/products/{id}", func(e *core.RequestEvent) error {
		id := e.Request.PathValue("id")
		fmt.Println("ID:", id)
		/*collection, err := app.FindCollectionByNameOrId("stores")
		if err != nil {
			return e.NotFoundError("", err)
		}*/
		record, err := app.FindRecordById("stores", id)
		if err != nil {
			return e.NotFoundError("", err)
		}
		return e.JSON(http.StatusOK, record)
	})
}

func ProductsBySlug(app *pocketbase.PocketBase, se *core.ServeEvent) *router.Route[*core.RequestEvent] {
	return se.Router.GET("/api/products/{slug}", func(e *core.RequestEvent) error {
		slug := e.Request.PathValue("slug")
		fmt.Println("Slug:", slug)
		record, err := e.App.FindFirstRecordByData("products", "slug", slug)
		if err != nil {
			return e.NotFoundError("Missing or invalid slug", err)
		}
		info, err := e.RequestInfo()
		if err != nil {
			return e.BadRequestError("Failed to retrieve request info", err)
		}
		canAccess, err := e.App.CanAccessRecord(record, info, record.Collection().ViewRule)
		if !canAccess {
			return e.ForbiddenError("", err)
		}
		return e.JSON(http.StatusOK, record)
	})
}
