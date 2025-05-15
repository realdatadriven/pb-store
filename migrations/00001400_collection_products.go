// Migration for collection: /api/products products - Products - Successfully retrieved list of products with pagination details
package migrations

import (
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/tools/types"
)

func init() {
	migrations.Register(func(app core.App) error {
		vendorsCol, _ := app.FindCollectionByNameOrId("vendors")
		storesCol, _ := app.FindCollectionByNameOrId("stores")
		usersCol, _ := app.FindCollectionByNameOrId("users")
		categoriesCol, _ := app.FindCollectionByNameOrId("categories")

		productsCol, err := app.FindCollectionByNameOrId("products")
		if err != nil {
			productsCol = core.NewBaseCollection("products")
		}
		productsCol.ListRule = types.Pointer("")
		productsCol.ViewRule = types.Pointer("")
		productsCol.Fields.Add(
			&core.BoolField{Name: "active"},
			&core.SelectField{Name: "status", Values: []string{"draft", "proposed", "published", "rejected"}},
			&core.TextField{Name: "type"},
			&core.RelationField{Name: "categoryId", CollectionId: categoriesCol.Id, Required: true},
			&core.TextField{Name: "currency"},
			&core.TextField{Name: "instructions"},
			&core.TextField{Name: "description"},
			&core.TextField{Name: "hsnCode"},
			&core.TextField{Name: "images"},
			&core.FileField{Name: "thumbnail"},
			&core.TextField{Name: "keywords"},
			&core.URLField{Name: "link"},
			&core.TextField{Name: "metaTitle"},
			&core.TextField{Name: "metaDescription"},
			&core.TextField{Name: "title", Required: true},
			&core.TextField{Name: "subtitle"},
			&core.NumberField{Name: "popularity", Min: types.Pointer[float64](-8388608), Max: types.Pointer[float64](8388607)},
			&core.NumberField{Name: "rank", Min: types.Pointer[float64](-8388608), Max: types.Pointer[float64](8388607)},
			&core.TextField{Name: "slug"},
			&core.TextField{Name: "expiryDate"},
			&core.NumberField{Name: "weight", Min: types.Pointer[float64](-8388608), Max: types.Pointer[float64](8388607)},
			&core.TextField{Name: "mfgDate"},
			&core.NumberField{Name: "mrp", Min: types.Pointer[float64](-8388608), Max: types.Pointer[float64](8388607)},
			&core.NumberField{Name: "price", Min: types.Pointer[float64](-8388608), Max: types.Pointer[float64](8388607)},
			&core.NumberField{Name: "costPerItem", Min: types.Pointer[float64](-8388608), Max: types.Pointer[float64](8388607)},
			&core.TextField{Name: "sku"},
			&core.NumberField{Name: "stock", Min: types.Pointer[float64](-8388608), Max: types.Pointer[float64](8388607)},
			&core.BoolField{Name: "allowBackorder", Required: true},
			&core.BoolField{Name: "manageInventory", Required: true},
			&core.NumberField{Name: "shippingWeight", Min: types.Pointer[float64](-8388608), Max: types.Pointer[float64](8388607)},
			&core.NumberField{Name: "shippingHeight", Min: types.Pointer[float64](-8388608), Max: types.Pointer[float64](8388607)},
			&core.NumberField{Name: "shippingLen", Min: types.Pointer[float64](-8388608), Max: types.Pointer[float64](8388607)},
			&core.NumberField{Name: "shippingWidth", Min: types.Pointer[float64](-8388608), Max: types.Pointer[float64](8388607)},
			&core.NumberField{Name: "height", Min: types.Pointer[float64](-8388608), Max: types.Pointer[float64](8388607)},
			&core.NumberField{Name: "width", Min: types.Pointer[float64](-8388608), Max: types.Pointer[float64](8388607)},
			&core.TextField{Name: "barcode"},
			&core.TextField{Name: "qrcode"},
			&core.NumberField{Name: "shippingCost", Min: types.Pointer[float64](-8388608), Max: types.Pointer[float64](8388607)},
			&core.BoolField{Name: "returnAllowed"},
			&core.BoolField{Name: "replaceAllowed"},
			&core.BoolField{Name: "allowReviews"},
			&core.NumberField{Name: "len", Min: types.Pointer[float64](-8388608), Max: types.Pointer[float64](8388607)},
			&core.TextField{Name: "productTags"},
			&core.TextField{Name: "originCountry"},
			&core.TextField{Name: "weightUnit"},
			&core.TextField{Name: "dimensionUnit"},
			&core.TextField{Name: "collectionId"},
			&core.TextField{Name: "styleCode"},
			&core.BoolField{Name: "isCustomizable"},
			&core.RelationField{Name: "vendorId", CollectionId: vendorsCol.Id},
			&core.RelationField{Name: "storeId", CollectionId: storesCol.Id, Required: true},
			&core.RelationField{Name: "userId", CollectionId: usersCol.Id},
			&core.TextField{Name: "groupedSku"},
			&core.TextField{Name: "remoteSku"},
			&core.AutodateField{Name: "createdAt", OnCreate: true},
			&core.AutodateField{Name: "updatedAt", OnUpdate: true},
			&core.DateField{Name: "deletedAt"},
		)
		if err := app.Save(productsCol); err != nil {
			return err
		}
		return nil
	}, nil)
}
