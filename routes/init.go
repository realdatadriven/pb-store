package routes

import (
	"net/http"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/router"
	"github.com/realdatadriven/pocket_store/internals/env"
)

func Init(app *pocketbase.PocketBase, se *core.ServeEvent) *router.Route[*core.RequestEvent] {
	return se.Router.GET("/api/init", func(e *core.RequestEvent) error {
		record, err := e.App.FindFirstRecordByData("stores", "name", env.GetString("PB_INIT_STORE_NAME", "test"))
		if err != nil {
			return e.NotFoundError("Missing or invalid slug", err)
		}
		println(record.Id)

		return e.JSON(http.StatusOK, map[string]any{"storeOne": record})
	})
}
