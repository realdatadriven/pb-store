package routes

import (
	"fmt"
	"net/http"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/router"
)

func StoresById(app *pocketbase.PocketBase, se *core.ServeEvent) *router.Route[*core.RequestEvent] {
	return se.Router.GET("/api/stores/{id}", func(e *core.RequestEvent) error {
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
