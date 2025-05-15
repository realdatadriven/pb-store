// Migration for collection: /api/admin/roles roles - Role - The list of role
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

		rolesCol, err := app.FindCollectionByNameOrId("roles")
		if err != nil { rolesCol = core.NewBaseCollection("roles") }
		rolesCol.ListRule = types.Pointer("")
		rolesCol.ViewRule = types.Pointer("")
		rolesCol.Fields.Add(
			&core.TextField{Name: "name" , Required: true},
			&core.BoolField{Name: "active"},
			&core.TextField{Name: "permissions"},
			&core.RelationField{Name: "storeId", CollectionId: storesCol.Id , Required: true},
			&core.RelationField{Name: "userId", CollectionId: usersCol.Id},
			&core.AutodateField{Name: "createdAt", OnCreate: true},
			&core.AutodateField{Name: "updatedAt", OnUpdate: true},
		)
		if err := app.Save(rolesCol); err != nil { return err }
		return nil
	}, nil)
}
