// Migration for collection: /api/reels reels - Reel - The list of role
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

		reelsCol, err := app.FindCollectionByNameOrId("reels")
		if err != nil { reelsCol = core.NewBaseCollection("reels") }
		reelsCol.ListRule = types.Pointer("")
		reelsCol.ViewRule = types.Pointer("")
		reelsCol.Fields.Add(
			&core.RelationField{Name: "productId", CollectionId: productsCol.Id},
			&core.TextField{Name: "name" , Required: true},
			&core.URLField{Name: "link" , Required: true},
			&core.TextField{Name: "type"},
			&core.TextField{Name: "video"},
			&core.TextField{Name: "poster"},
			&core.BoolField{Name: "active"},
			&core.RelationField{Name: "storeId", CollectionId: storesCol.Id , Required: true},
			&core.RelationField{Name: "userId", CollectionId: usersCol.Id},
			&core.AutodateField{Name: "createdAt", OnCreate: true},
			&core.AutodateField{Name: "updatedAt", OnUpdate: true},
		)
		if err := app.Save(reelsCol); err != nil { return err }
		return nil
	}, nil)
}
