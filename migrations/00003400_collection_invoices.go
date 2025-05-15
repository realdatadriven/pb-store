// Migration for collection: /api/admin/invoices invoices - Invoices - Successfully retrieved list of invoices with pagination details
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

		invoicesCol, err := app.FindCollectionByNameOrId("invoices")
		if err != nil { invoicesCol = core.NewBaseCollection("invoices") }
		invoicesCol.ListRule = types.Pointer("")
		invoicesCol.ViewRule = types.Pointer("")
		invoicesCol.Fields.Add(
			&core.TextField{Name: "invoiceNo" , Required: true},
			&core.TextField{Name: "taxType"},
			&core.TextField{Name: "reportStatus"},
			&core.TextField{Name: "orderNo" , Required: true},
			&core.NumberField{Name: "totalTaxAmount", Min: types.Pointer[float64](-8388608), Max: types.Pointer[float64](8388607)},
			&core.NumberField{Name: "ix", Min: types.Pointer[float64](-8388608), Max: types.Pointer[float64](8388607)},
			&core.TextField{Name: "comment"},
			&core.RelationField{Name: "vendorId", CollectionId: vendorsCol.Id},
			&core.RelationField{Name: "storeId", CollectionId: storesCol.Id , Required: true},
			&core.RelationField{Name: "userId", CollectionId: usersCol.Id},
			&core.AutodateField{Name: "createdAt", OnCreate: true},
			&core.AutodateField{Name: "updatedAt", OnUpdate: true},
		)
		if err := app.Save(invoicesCol); err != nil { return err }
		return nil
	}, nil)
}
