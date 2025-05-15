// Migration for collection: /api/admin/teams teams - Teams - Successfully retrieved list of tags
package migrations
import (
    "github.com/pocketbase/pocketbase/core"
    "github.com/pocketbase/pocketbase/migrations"
    "github.com/pocketbase/pocketbase/tools/types"
)

func init() {
migrations.Register(func(app core.App) error {
		storesCol, _ := app.FindCollectionByNameOrId("stores")
vendorsCol, _ := app.FindCollectionByNameOrId("vendors")

		teamsCol, err := app.FindCollectionByNameOrId("teams")
		if err != nil { teamsCol = core.NewBaseCollection("teams") }
		teamsCol.ListRule = types.Pointer("")
		teamsCol.ViewRule = types.Pointer("")
		teamsCol.Fields.Add(
			&core.TextField{Name: "role" , Required: true},
			&core.EmailField{Name: "email"},
			&core.TextField{Name: "phone"},
			&core.TextField{Name: "avatar"},
			&core.BoolField{Name: "approved"},
			&core.BoolField{Name: "isJoined"},
			&core.BoolField{Name: "isCreator"},
			&core.TextField{Name: "zips"},
			&core.RelationField{Name: "storeId", CollectionId: storesCol.Id , Required: true},
			&core.RelationField{Name: "vendorId", CollectionId: vendorsCol.Id},
			&core.AutodateField{Name: "createdAt", OnCreate: true},
			&core.AutodateField{Name: "updatedAt", OnUpdate: true},
		)
		if err := app.Save(teamsCol); err != nil { return err }
		return nil
	}, nil)
}
