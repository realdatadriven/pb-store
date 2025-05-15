// Migration for collection: /api/pages pages - Pages - Successfully retrieved list of pages
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

		pagesCol, err := app.FindCollectionByNameOrId("pages")
		if err != nil { pagesCol = core.NewBaseCollection("pages") }
		pagesCol.ListRule = types.Pointer("")
		pagesCol.ViewRule = types.Pointer("")
		pagesCol.Fields.Add(
			&core.TextField{Name: "name" , Required: true},
			&core.TextField{Name: "slug" , Required: true},
			&core.TextField{Name: "type"},
			&core.EditorField{Name: "content"},
			&core.TextField{Name: "metaDescription"},
			&core.TextField{Name: "metaKeywords"},
			&core.TextField{Name: "metaTitle"},
			&core.TextField{Name: "status"},
			&core.NumberField{Name: "rank", Min: types.Pointer[float64](-8388608), Max: types.Pointer[float64](8388607)},
			&core.RelationField{Name: "storeId", CollectionId: storesCol.Id , Required: true},
			&core.RelationField{Name: "userId", CollectionId: usersCol.Id},
			&core.AutodateField{Name: "createdAt", OnCreate: true},
			&core.AutodateField{Name: "updatedAt", OnUpdate: true},
		)
		if err := app.Save(pagesCol); err != nil { return err }
		return nil
	}, nil)
}
