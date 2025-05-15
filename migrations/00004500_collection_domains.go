// Migration for collection: /api/admin/domains domains - Domains - Successfully retrieved list of blogs
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

		domainsCol, err := app.FindCollectionByNameOrId("domains")
		if err != nil { domainsCol = core.NewBaseCollection("domains") }
		domainsCol.ListRule = types.Pointer("")
		domainsCol.ViewRule = types.Pointer("")
		domainsCol.Fields.Add(
			&core.TextField{Name: "status"},
			&core.TextField{Name: "name" , Required: true},
			&core.TextField{Name: "description"},
			&core.BoolField{Name: "isPrimary"},
			&core.BoolField{Name: "isPropagated"},
			&core.BoolField{Name: "isActive"},
			&core.BoolField{Name: "isDeleted"},
			&core.BoolField{Name: "isVerified"},
			&core.RelationField{Name: "storeId", CollectionId: storesCol.Id , Required: true},
			&core.RelationField{Name: "userId", CollectionId: usersCol.Id},
			&core.AutodateField{Name: "createdAt", OnCreate: true},
			&core.AutodateField{Name: "updatedAt", OnUpdate: true},
		)
		if err := app.Save(domainsCol); err != nil { return err }
		return nil
	}, nil)
}
