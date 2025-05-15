// Migration for collection: /api/admin/webhooks webhooks - Webhooks - Successfully retrieved list of webhooks with pagination details
package migrations
import (
    "github.com/pocketbase/pocketbase/core"
    "github.com/pocketbase/pocketbase/migrations"
    "github.com/pocketbase/pocketbase/tools/types"
)

func init() {
migrations.Register(func(app core.App) error {
		storesCol, _ := app.FindCollectionByNameOrId("stores")
usersCol, _ := app.FindCollectionByNameOrId("users")

		webhooksCol, err := app.FindCollectionByNameOrId("webhooks")
		if err != nil { webhooksCol = core.NewBaseCollection("webhooks") }
		webhooksCol.ListRule = types.Pointer("")
		webhooksCol.ViewRule = types.Pointer("")
		webhooksCol.Fields.Add(
			&core.TextField{Name: "url"},
			&core.TextField{Name: "password"},
			&core.TextField{Name: "name"},
			&core.TextField{Name: "description"},
			&core.TextField{Name: "event"},
			&core.NumberField{Name: "rank", Min: types.Pointer[float64](-2147483648), Max: types.Pointer[float64](2147483647)},
			&core.BoolField{Name: "active"},
			&core.RelationField{Name: "storeId", CollectionId: storesCol.Id , Required: true},
			&core.RelationField{Name: "userId", CollectionId: usersCol.Id},
			&core.AutodateField{Name: "createdAt", OnCreate: true},
			&core.AutodateField{Name: "updatedAt", OnUpdate: true},
		)
		if err := app.Save(webhooksCol); err != nil { return err }
		return nil
	}, nil)
}
