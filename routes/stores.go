package routes

import (
	"fmt"
	"net/http"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/router"
)

func Stores(app *pocketbase.PocketBase, se *core.ServeEvent) *router.Route[*core.RequestEvent] {
	return se.Router.GET("/api/stores/{id}", func(e *core.RequestEvent) error {
		id := e.Request.PathValue("id")
		fmt.Println("ID:", id)
		var err error
		if err != nil {
			return e.NotFoundError("", err)
		}
		return e.JSON(http.StatusOK, []map[string]any{})
	})
}
