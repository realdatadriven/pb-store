// Migration for collection: /api/admin/warehouse warehouse - Warehouse - Successfully retrieved list of warehouse
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

		warehouseCol, err := app.FindCollectionByNameOrId("warehouse")
		if err != nil {
			warehouseCol = core.NewBaseCollection("warehouse")
		}
		warehouseCol.ListRule = types.Pointer("")
		warehouseCol.ViewRule = types.Pointer("")
		warehouseCol.Fields.Add(
			&core.BoolField{Name: "isActive"},
			&core.TextField{Name: "name", Required: true},
			&core.TextField{Name: "description"},
			&core.TextField{Name: "address_1", Required: true},
			&core.TextField{Name: "address_2"},
			&core.NumberField{Name: "lat"},
			&core.NumberField{Name: "lng"},
			&core.TextField{Name: "zip", Required: true},
			&core.TextField{Name: "city", Required: true},
			&core.TextField{Name: "state", Required: true},
			&core.TextField{Name: "countryCode", Required: true},
			&core.RelationField{Name: "vendorId", CollectionId: vendorsCol.Id},
			&core.RelationField{Name: "storeId", CollectionId: storesCol.Id, Required: true},
			&core.RelationField{Name: "userId", CollectionId: usersCol.Id},
			&core.AutodateField{Name: "createdAt", OnCreate: true},
			&core.AutodateField{Name: "updatedAt", OnUpdate: true},
		)
		if err := app.Save(warehouseCol); err != nil {
			return err
		}
		return nil
	}, nil)
}
