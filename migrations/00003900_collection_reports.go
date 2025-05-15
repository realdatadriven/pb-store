// Migration for collection: /api/admin/reports reports - Reports - The list of reports
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

		reportsCol, err := app.FindCollectionByNameOrId("reports")
		if err != nil { reportsCol = core.NewBaseCollection("reports") }
		reportsCol.ListRule = types.Pointer("")
		reportsCol.ViewRule = types.Pointer("")
		reportsCol.Fields.Add(
			&core.TextField{Name: "name" , Required: true},
			&core.TextField{Name: "query" , Required: true},
			&core.TextField{Name: "description"},
			&core.RelationField{Name: "storeId", CollectionId: storesCol.Id , Required: true},
			&core.RelationField{Name: "userId", CollectionId: usersCol.Id},
			&core.AutodateField{Name: "createdAt", OnCreate: true},
			&core.AutodateField{Name: "updatedAt", OnUpdate: true},
		)
		if err := app.Save(reportsCol); err != nil { return err }
		return nil
	}, nil)
}
