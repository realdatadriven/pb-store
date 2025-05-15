// Migration for collection: /api/menu menu - Menu (Public) - Successfully retrieved list of banners
package migrations
import (
    "github.com/pocketbase/pocketbase/core"
    "github.com/pocketbase/pocketbase/migrations"
    "github.com/pocketbase/pocketbase/tools/types"
)

func init() {
migrations.Register(func(app core.App) error {
		usersCol, _ := app.FindCollectionByNameOrId("users")

		menuCol, err := app.FindCollectionByNameOrId("menu")
		if err != nil { menuCol = core.NewBaseCollection("menu") }
		menuCol.ListRule = types.Pointer("")
		menuCol.ViewRule = types.Pointer("")
		menuCol.Fields.Add(
			&core.BoolField{Name: "active"},
			&core.TextField{Name: "name"},
			&core.TextField{Name: "menuId"},
			&core.URLField{Name: "link"},
			&core.RelationField{Name: "userId", CollectionId: usersCol.Id},
		)
		if err := app.Save(menuCol); err != nil { return err }
		return nil
	}, nil)
}
