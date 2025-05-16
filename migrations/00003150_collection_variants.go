// Migration for collection: /api/admin/variants variants - Variants - List of Variants
package migrations

import (
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/tools/types"
)

func init() {
	migrations.Register(func(app core.App) error {
		productsCol, _ := app.FindCollectionByNameOrId("products")
		vendorsCol, _ := app.FindCollectionByNameOrId("vendors")
		storesCol, _ := app.FindCollectionByNameOrId("stores")
		usersCol, _ := app.FindCollectionByNameOrId("users")

		variantsCol, err := app.FindCollectionByNameOrId("variants")
		if err != nil {
			variantsCol = core.NewBaseCollection("variants")
		}
		variantsCol.ListRule = types.Pointer("")
		variantsCol.ViewRule = types.Pointer("")
		variantsCol.Fields.Add(
			&core.TextField{Name: "title"},
			&core.TextField{Name: "variantValue", Required: true},
			&core.RelationField{Name: "productId", CollectionId: productsCol.Id, Required: true},
			&core.TextField{Name: "sku"},
			&core.TextField{Name: "barcode"},
			&core.TextField{Name: "batchNo"},
			&core.NumberField{Name: "reorderLevel", Min: types.Pointer[float64](-8388608), Max: types.Pointer[float64](8388607)},
			&core.NumberField{Name: "stock", Min: types.Pointer[float64](-8388608), Max: types.Pointer[float64](8388607)},
			&core.BoolField{Name: "allowBackorder", Required: true},
			&core.BoolField{Name: "manageInventory", Required: true},
			&core.TextField{Name: "hsCode"},
			&core.TextField{Name: "originCountry"},
			&core.TextField{Name: "midCode"},
			&core.TextField{Name: "material"},
			&core.NumberField{Name: "weight", Min: types.Pointer[float64](-8388608), Max: types.Pointer[float64](8388607)},
			&core.NumberField{Name: "length", Min: types.Pointer[float64](-8388608), Max: types.Pointer[float64](8388607)},
			&core.NumberField{Name: "height", Min: types.Pointer[float64](-8388608), Max: types.Pointer[float64](8388607)},
			&core.NumberField{Name: "width", Min: types.Pointer[float64](-8388608), Max: types.Pointer[float64](8388607)},
			&core.NumberField{Name: "price", Min: types.Pointer[float64](-8388608), Max: types.Pointer[float64](8388607)},
			&core.NumberField{Name: "costPerItem", Min: types.Pointer[float64](-8388608), Max: types.Pointer[float64](8388607)},
			&core.DateField{Name: "mfgDate"},
			&core.DateField{Name: "expiryDate"},
			&core.BoolField{Name: "returnAllowed"},
			&core.BoolField{Name: "replaceAllowed"},
			&core.NumberField{Name: "mrp", Min: types.Pointer[float64](-8388608), Max: types.Pointer[float64](8388607)},
			&core.TextField{Name: "img"},
			&core.FileField{Name: "thumbnail"},
			&core.TextField{Name: "images"},
			&core.TextField{Name: "description"},
			&core.NumberField{Name: "len", Min: types.Pointer[float64](-8388608), Max: types.Pointer[float64](8388607)},
			&core.NumberField{Name: "rank", Min: types.Pointer[float64](-8388608), Max: types.Pointer[float64](8388607)},
			&core.NumberField{Name: "shippingWeight", Min: types.Pointer[float64](-8388608), Max: types.Pointer[float64](8388607)},
			&core.NumberField{Name: "shippingHeight", Min: types.Pointer[float64](-8388608), Max: types.Pointer[float64](8388607)},
			&core.NumberField{Name: "shippingLen", Min: types.Pointer[float64](-8388608), Max: types.Pointer[float64](8388607)},
			&core.NumberField{Name: "shippingWidth", Min: types.Pointer[float64](-8388608), Max: types.Pointer[float64](8388607)},
			&core.NumberField{Name: "shippingCost", Min: types.Pointer[float64](-8388608), Max: types.Pointer[float64](8388607)},
			&core.NumberField{Name: "variantRank", Min: types.Pointer[float64](-8388608), Max: types.Pointer[float64](8388607)},
			&core.RelationField{Name: "vendorId", CollectionId: vendorsCol.Id},
			&core.RelationField{Name: "storeId", CollectionId: storesCol.Id, Required: true},
			&core.RelationField{Name: "userId", CollectionId: usersCol.Id},
			&core.AutodateField{Name: "createdAt", OnCreate: true},
			&core.AutodateField{Name: "updatedAt", OnUpdate: true},
			&core.DateField{Name: "deletedAt"},
		)
		if err := app.Save(variantsCol); err != nil {
			return err
		}
		return nil
	}, nil)
}
