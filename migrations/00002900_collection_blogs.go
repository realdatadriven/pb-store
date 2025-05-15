// Migration for collection: /api/blogs blogs - Blogs - Successfully retrieved list of blogs
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

		blogsCol, err := app.FindCollectionByNameOrId("blogs")
		if err != nil { blogsCol = core.NewBaseCollection("blogs") }
		blogsCol.ListRule = types.Pointer("")
		blogsCol.ViewRule = types.Pointer("")
		blogsCol.Fields.Add(
			&core.TextField{Name: "status"},
			&core.TextField{Name: "title"},
			&core.EditorField{Name: "content"},
			&core.TextField{Name: "slug"},
			&core.RelationField{Name: "storeId", CollectionId: storesCol.Id , Required: true},
			&core.RelationField{Name: "userId", CollectionId: usersCol.Id},
			&core.BoolField{Name: "isPublished"},
			&core.BoolField{Name: "isFeatured"},
			&core.NumberField{Name: "views", Min: types.Pointer[float64](-8388608), Max: types.Pointer[float64](8388607)},
			&core.NumberField{Name: "likes", Min: types.Pointer[float64](-8388608), Max: types.Pointer[float64](8388607)},
			&core.AutodateField{Name: "createdAt", OnCreate: true},
			&core.AutodateField{Name: "updatedAt", OnUpdate: true},
		)
		if err := app.Save(blogsCol); err != nil { return err }
		return nil
	}, nil)
}
