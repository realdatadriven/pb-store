// Migration for collection: /api/address address - Address - Successfully retrieved the list of addresses with pagination details
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

		addressCol, err := app.FindCollectionByNameOrId("address")
		if err != nil { addressCol = core.NewBaseCollection("address") }
		addressCol.ListRule = types.Pointer("")
		addressCol.ViewRule = types.Pointer("")
		addressCol.Fields.Add(
			&core.BoolField{Name: "active"},
			&core.TextField{Name: "address_1"},
			&core.TextField{Name: "address_2"},
			&core.TextField{Name: "city"},
			&core.TextField{Name: "countryCode"},
			&core.TextField{Name: "deliveryInstructions"},
			&core.EmailField{Name: "email"},
			&core.TextField{Name: "firstName"},
			&core.BoolField{Name: "isPrimary"},
			&core.BoolField{Name: "isResidential"},
			&core.TextField{Name: "lastName"},
			&core.TextField{Name: "lat"},
			&core.TextField{Name: "lng"},
			&core.TextField{Name: "locality"},
			&core.TextField{Name: "phone"},
			&core.TextField{Name: "state"},
			&core.RelationField{Name: "storeId", CollectionId: storesCol.Id , Required: true},
			&core.RelationField{Name: "userId", CollectionId: usersCol.Id},
			&core.TextField{Name: "zip"},
			&core.AutodateField{Name: "createdAt", OnCreate: true},
			&core.AutodateField{Name: "updatedAt", OnUpdate: true},
		)
		if err := app.Save(addressCol); err != nil { return err }
		return nil
	}, nil)
}
