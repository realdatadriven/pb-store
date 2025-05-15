// Migration for collection: /api/vendors vendors - Vendors - The list of vendors
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

		vendorsCol, err := app.FindCollectionByNameOrId("vendors")
		if err != nil { vendorsCol = core.NewBaseCollection("vendors") }
		vendorsCol.ListRule = types.Pointer("")
		vendorsCol.ViewRule = types.Pointer("")
		vendorsCol.Fields.Add(
			&core.TextField{Name: "status"},
			&core.TextField{Name: "address"},
			&core.EmailField{Name: "email"},
			&core.TextField{Name: "phone"},
			&core.TextField{Name: "dialCode"},
			&core.TextField{Name: "name"},
			&core.EmailField{Name: "email2"},
			&core.TextField{Name: "banners"},
			&core.FileField{Name: "logo"},
			&core.TextField{Name: "countryName"},
			&core.TextField{Name: "country"},
			&core.TextField{Name: "about"},
			&core.TextField{Name: "businessName" , Required: true},
			&core.URLField{Name: "website"},
			&core.TextField{Name: "description"},
			&core.TextField{Name: "info"},
			&core.NumberField{Name: "shippingCharges", Min: types.Pointer[float64](-8388608), Max: types.Pointer[float64](8388607)},
			&core.NumberField{Name: "codCharges", Min: types.Pointer[float64](-8388608), Max: types.Pointer[float64](8388607)},
			&core.TextField{Name: "slug"},
			&core.TextField{Name: "featuredImage"},
			&core.BoolField{Name: "isEmailVerified"},
			&core.BoolField{Name: "isPhoneVerified"},
			&core.TextField{Name: "address_1"},
			&core.TextField{Name: "address_2"},
			&core.TextField{Name: "city"},
			&core.BoolField{Name: "isApproved"},
			&core.BoolField{Name: "isDeleted"},
			&core.TextField{Name: "state"},
			&core.TextField{Name: "tax_number"},
			&core.TextField{Name: "zip"},
			&core.NumberField{Name: "commissionRate", Min: types.Pointer[float64](-8388608), Max: types.Pointer[float64](8388607)},
			&core.RelationField{Name: "storeId", CollectionId: storesCol.Id , Required: true},
			&core.RelationField{Name: "userId", CollectionId: usersCol.Id , Required: true},
			&core.AutodateField{Name: "createdAt", OnCreate: true},
			&core.AutodateField{Name: "updatedAt", OnUpdate: true},
		)
		if err := app.Save(vendorsCol); err != nil { return err }
		return nil
	}, nil)
}
