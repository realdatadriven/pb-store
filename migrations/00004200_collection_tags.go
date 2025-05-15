// Migration for collection: /api/admin/tags tags - Tags - Successfully retrieved list of tags
package migrations
import (
    "github.com/pocketbase/pocketbase/core"
    "github.com/pocketbase/pocketbase/migrations"
    "github.com/pocketbase/pocketbase/tools/types"
)

func init() {
migrations.Register(func(app core.App) error {
		productsCol, _ := app.FindCollectionByNameOrId("products")
storesCol, _ := app.FindCollectionByNameOrId("stores")
usersCol, _ := app.FindCollectionByNameOrId("users")

		tagsCol, err := app.FindCollectionByNameOrId("tags")
		if err != nil { tagsCol = core.NewBaseCollection("tags") }
		tagsCol.ListRule = types.Pointer("")
		tagsCol.ViewRule = types.Pointer("")
		tagsCol.Fields.Add(
			&core.TextField{Name: "name" , Required: true},
			&core.TextField{Name: "slug"},
			&core.TextField{Name: "description"},
			&core.TextField{Name: "type"},
			&core.TextField{Name: "colorCode"},
			&core.NumberField{Name: "rank", Min: types.Pointer[float64](-2147483648), Max: types.Pointer[float64](2147483647)},
			&core.BoolField{Name: "active"},
			&core.RelationField{Name: "productId", CollectionId: productsCol.Id , Required: true},
			&core.RelationField{Name: "storeId", CollectionId: storesCol.Id , Required: true},
			&core.RelationField{Name: "userId", CollectionId: usersCol.Id},
			&core.AutodateField{Name: "createdAt", OnCreate: true},
			&core.AutodateField{Name: "updatedAt", OnUpdate: true},
			&core.DateField{Name: "deletedAt"},
		)
		if err := app.Save(tagsCol); err != nil { return err }
		return nil
	}, nil)
}
