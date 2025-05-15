// Migration for collection: /api/admin/wishlists wishlists - Wishlists - The list of wishlists
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
variantsCol, _ := app.FindCollectionByNameOrId("variants")

		wishlistsCol, err := app.FindCollectionByNameOrId("wishlists")
		if err != nil { wishlistsCol = core.NewBaseCollection("wishlists") }
		wishlistsCol.ListRule = types.Pointer("")
		wishlistsCol.ViewRule = types.Pointer("")
		wishlistsCol.Fields.Add(
			&core.BoolField{Name: "active"},
			&core.RelationField{Name: "productId", CollectionId: productsCol.Id , Required: true},
			&core.RelationField{Name: "storeId", CollectionId: storesCol.Id , Required: true},
			&core.RelationField{Name: "userId", CollectionId: usersCol.Id},
			&core.RelationField{Name: "variantId", CollectionId: variantsCol.Id , Required: true},
			&core.AutodateField{Name: "createdAt", OnCreate: true},
			&core.AutodateField{Name: "updatedAt", OnUpdate: true},
		)
		if err := app.Save(wishlistsCol); err != nil { return err }
		return nil
	}, nil)
}
