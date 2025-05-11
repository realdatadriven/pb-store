package migrations

import (
	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/tools/types"
)

func init() {
	m.Register(func(app core.App) error {

		usersCol, err := app.FindCollectionByNameOrId("users")
		if err != nil {
			return err
		}
		storeCol, err := app.FindCollectionByNameOrId("stores")
		if err != nil {
			return err
		}
		collection, err := app.FindCollectionByNameOrId("categories")
		if err != nil {
			collection = core.NewBaseCollection("categories")
		}
		// restrict the list and view rules for record owners
		collection.ListRule = types.Pointer("")
		collection.ViewRule = types.Pointer("")
		// add extra fields in addition to the default ones
		collection.Fields.Add(
			&core.BoolField{Name: "isActive"},
			&core.BoolField{Name: "isInternal"},
			&core.BoolField{Name: "isMegamenu"},
			&core.TextField{Name: "thumbnail"},
			&core.TextField{Name: "path"},
			&core.NumberField{
				Name: "level",
				Min:  types.Pointer[float64](-2147483648),
				Max:  types.Pointer[float64](2147483647),
			},
			&core.TextField{Name: "description"},
			&core.BoolField{Name: "isFeatured"},
			&core.TextField{Name: "keywords"},
			&core.NumberField{
				Name: "rank",
				Min:  types.Pointer[float64](-2147483648),
				Max:  types.Pointer[float64](2147483647),
			},
			&core.URLField{Name: "link", Presentable: true},
			&core.TextField{Name: "metaDescription"},
			&core.TextField{Name: "metaKeywords"},
			&core.TextField{Name: "metaTitle"},
			&core.TextField{
				Name:     "name",
				Required: true,
			},
			&core.TextField{Name: "parentCategoryId"},
			&core.TextField{Name: "store"},
			&core.TextField{Name: "slug"},
			&core.NumberField{
				Name: "activeProducts",
				Min:  types.Pointer[float64](-2147483648),
				Max:  types.Pointer[float64](2147483647),
			},
			&core.NumberField{
				Name: "inactiveProducts",
				Min:  types.Pointer[float64](-2147483648),
				Max:  types.Pointer[float64](2147483647),
			},
			&core.RelationField{
				Name:         "userId",
				CollectionId: usersCol.Id, // links to the built-in users collection
				Required:     true,
			},
			&core.RelationField{
				Name:         "storeId",
				CollectionId: storeCol.Id,
				Required:     true,
			},
			&core.AutodateField{Name: "createdAt", OnCreate: true},
			&core.AutodateField{Name: "updatedAt", OnUpdate: true},
		)

		return app.Save(collection)
	}, /*func(app core.App) error { // optional revert operation
		    collection, err := app.FindCollectionByNameOrId("clients")
		    if err != nil {
		        return err
		    }

		    return app.Delete(collection)
		}*/nil)
}
