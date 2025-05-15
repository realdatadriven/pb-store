// Migration for collection: /api/admin/plugin plugin - Plugin - Successfully retrieved list of plugins with pagination details
package migrations
import (
    "github.com/pocketbase/pocketbase/core"
    "github.com/pocketbase/pocketbase/migrations"
    "github.com/pocketbase/pocketbase/tools/types"
)

func init() {
migrations.Register(func(app core.App) error {
		usersCol, _ := app.FindCollectionByNameOrId("users")

		pluginCol, err := app.FindCollectionByNameOrId("plugin")
		if err != nil { pluginCol = core.NewBaseCollection("plugin") }
		pluginCol.ListRule = types.Pointer("")
		pluginCol.ViewRule = types.Pointer("")
		pluginCol.Fields.Add(
			&core.TextField{Name: "code" , Required: true},
			&core.BoolField{Name: "isActive"},
			&core.RelationField{Name: "userId", CollectionId: usersCol.Id},
		)
		if err := app.Save(pluginCol); err != nil { return err }
		return nil
	}, nil)
}
