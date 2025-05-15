// Migration for collection: /api/inventory inventory - Inventory - Successfully retrieved list of inventory
package migrations
import (
    "github.com/pocketbase/pocketbase/core"
    "github.com/pocketbase/pocketbase/migrations"
    "github.com/pocketbase/pocketbase/tools/types"
)

func init() {
migrations.Register(func(app core.App) error {
		productsCol, _ := app.FindCollectionByNameOrId("products")
variantsCol, _ := app.FindCollectionByNameOrId("variants")
vendorsCol, _ := app.FindCollectionByNameOrId("vendors")
storesCol, _ := app.FindCollectionByNameOrId("stores")
usersCol, _ := app.FindCollectionByNameOrId("users")

		inventoryCol, err := app.FindCollectionByNameOrId("inventory")
		if err != nil { inventoryCol = core.NewBaseCollection("inventory") }
		inventoryCol.ListRule = types.Pointer("")
		inventoryCol.ViewRule = types.Pointer("")
		inventoryCol.Fields.Add(
			&core.RelationField{Name: "productId", CollectionId: productsCol.Id , Required: true},
			&core.RelationField{Name: "variantId", CollectionId: variantsCol.Id},
			&core.TextField{Name: "sku" , Required: true},
			&core.TextField{Name: "barcode"},
			&core.NumberField{Name: "stock", Min: types.Pointer[float64](-2147483648), Max: types.Pointer[float64](2147483647)},
			&core.NumberField{Name: "lowStockThreshold", Min: types.Pointer[float64](-2147483648), Max: types.Pointer[float64](2147483647)},
			&core.NumberField{Name: "reorderPoint", Min: types.Pointer[float64](-2147483648), Max: types.Pointer[float64](2147483647)},
			&core.TextField{Name: "costPrice"},
			&core.TextField{Name: "location"},
			&core.TextField{Name: "warehouseId"},
			&core.NumberField{Name: "leadTime", Min: types.Pointer[float64](-2147483648), Max: types.Pointer[float64](2147483647)},
			&core.BoolField{Name: "active"},
			&core.RelationField{Name: "vendorId", CollectionId: vendorsCol.Id},
			&core.RelationField{Name: "storeId", CollectionId: storesCol.Id , Required: true},
			&core.RelationField{Name: "userId", CollectionId: usersCol.Id},
			&core.AutodateField{Name: "createdAt", OnCreate: true},
			&core.AutodateField{Name: "updatedAt", OnUpdate: true},
			&core.DateField{Name: "deletedAt"},
		)
		if err := app.Save(inventoryCol); err != nil { return err }
		return nil
	}, nil)
}
