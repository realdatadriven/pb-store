// Migration for collection: /api/categories categories - Categories - Successfully retrieved list of categories with pagination details
package migrations

import (
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/tools/types"
)

func init() {
	migrations.Register(func(app core.App) error {
		usersCol, _ := app.FindCollectionByNameOrId("users")
		storesCol, _ := app.FindCollectionByNameOrId("stores")

		categoriesCol, err := app.FindCollectionByNameOrId("categories")
		if err != nil {
			categoriesCol = core.NewBaseCollection("categories")
		}
		categoriesCol.ListRule = types.Pointer("")
		categoriesCol.ViewRule = types.Pointer("")
		categoriesCol.Fields.Add(
			&core.BoolField{Name: "isActive"},
			&core.BoolField{Name: "isInternal"},
			&core.BoolField{Name: "isMegamenu"},
			&core.FileField{Name: "thumbnail"},
			&core.TextField{Name: "path"},
			&core.NumberField{Name: "level", Min: types.Pointer[float64](-2147483648), Max: types.Pointer[float64](2147483647)},
			&core.TextField{Name: "description"},
			&core.BoolField{Name: "isFeatured"},
			&core.TextField{Name: "keywords"},
			&core.NumberField{Name: "rank", Min: types.Pointer[float64](-2147483648), Max: types.Pointer[float64](2147483647)},
			&core.URLField{Name: "link"},
			&core.TextField{Name: "metaDescription"},
			&core.TextField{Name: "metaKeywords"},
			&core.TextField{Name: "metaTitle"},
			&core.TextField{Name: "name", Required: true},
			&core.RelationField{Name: "parentCategoryId", CollectionId: categoriesCol.Id},
			&core.TextField{Name: "store"},
			&core.TextField{Name: "slug"},
			&core.NumberField{Name: "activeProducts", Min: types.Pointer[float64](-2147483648), Max: types.Pointer[float64](2147483647)},
			&core.NumberField{Name: "inactiveProducts", Min: types.Pointer[float64](-2147483648), Max: types.Pointer[float64](2147483647)},
			&core.RelationField{Name: "userId", CollectionId: usersCol.Id},
			&core.RelationField{Name: "storeId", CollectionId: storesCol.Id, Required: true},
			&core.AutodateField{Name: "createdAt", OnCreate: true},
			&core.AutodateField{Name: "updatedAt", OnUpdate: true},
		)
		if err := app.Save(categoriesCol); err != nil {
			return err
		}
		return nil
	}, nil)
}
