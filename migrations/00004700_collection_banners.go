// Migration for collection: /api/banners banners - Banners - Successfully retrieved list of banners
package migrations
import (
    "github.com/pocketbase/pocketbase/core"
    "github.com/pocketbase/pocketbase/migrations"
    "github.com/pocketbase/pocketbase/tools/types"
)

func init() {
migrations.Register(func(app core.App) error {
		pagesCol, _ := app.FindCollectionByNameOrId("pages")
storesCol, _ := app.FindCollectionByNameOrId("stores")
usersCol, _ := app.FindCollectionByNameOrId("users")

		bannersCol, err := app.FindCollectionByNameOrId("banners")
		if err != nil { bannersCol = core.NewBaseCollection("banners") }
		bannersCol.ListRule = types.Pointer("")
		bannersCol.ViewRule = types.Pointer("")
		bannersCol.Fields.Add(
			&core.BoolField{Name: "active"},
			&core.BoolField{Name: "demo"},
			&core.TextField{Name: "groupId"},
			&core.TextField{Name: "groupTitle"},
			&core.TextField{Name: "heading"},
			&core.TextField{Name: "img" , Required: true},
			&core.TextField{Name: "imgCdn"},
			&core.URLField{Name: "link"},
			&core.RelationField{Name: "pageId", CollectionId: pagesCol.Id},
			&core.TextField{Name: "pageType"},
			&core.BoolField{Name: "isLinkExternal"},
			&core.NumberField{Name: "rank", Min: types.Pointer[float64](-2147483648), Max: types.Pointer[float64](2147483647)},
			&core.RelationField{Name: "storeId", CollectionId: storesCol.Id , Required: true},
			&core.RelationField{Name: "userId", CollectionId: usersCol.Id},
			&core.TextField{Name: "type"},
			&core.BoolField{Name: "isMobile"},
			&core.TextField{Name: "description"},
			&core.TextField{Name: "title"},
			&core.NumberField{Name: "bannerId", Min: types.Pointer[float64](-2147483648), Max: types.Pointer[float64](2147483647)},
			&core.NumberField{Name: "fieldGrid", Min: types.Pointer[float64](-2147483648), Max: types.Pointer[float64](2147483647)},
			&core.BoolField{Name: "scroll"},
			&core.AutodateField{Name: "createdAt", OnCreate: true},
			&core.AutodateField{Name: "updatedAt", OnUpdate: true},
		)
		if err := app.Save(bannersCol); err != nil { return err }
		return nil
	}, nil)
}
