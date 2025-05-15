// Migration for collection: /api/carts carts - Cart - Successfully retrieved the requested cart
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

		cartsCol, err := app.FindCollectionByNameOrId("carts")
		if err != nil { cartsCol = core.NewBaseCollection("carts") }
		cartsCol.ListRule = types.Pointer("")
		cartsCol.ViewRule = types.Pointer("")
		cartsCol.Fields.Add(
			&core.EmailField{Name: "email"},
			&core.TextField{Name: "phone"},
			&core.TextField{Name: "billingAddressId"},
			&core.TextField{Name: "shippingAddressId"},
			&core.TextField{Name: "regionId"},
			&core.RelationField{Name: "userId", CollectionId: usersCol.Id},
			&core.TextField{Name: "couponCode"},
			&core.NumberField{Name: "discountAmount", Min: types.Pointer[float64](-8388608), Max: types.Pointer[float64](8388607) , Required: true},
			&core.TextField{Name: "couponAppliedDate"},
			&core.BoolField{Name: "needAddress" , Required: true},
			&core.BoolField{Name: "isCodAvailable" , Required: true},
			&core.TextField{Name: "paymentId"},
			&core.SelectField{Name: "type", Values: []string{"default","swap","draft_order","payment_link","claim"}},
			&core.DateField{Name: "completedAt"},
			&core.DateField{Name: "paymentAuthorizedAt"},
			&core.TextField{Name: "idempotencyKey"},
			&core.TextField{Name: "salesChannelId"},
			&core.NumberField{Name: "qty", Min: types.Pointer[float64](-2147483648), Max: types.Pointer[float64](2147483647) , Required: true},
			&core.NumberField{Name: "shippingCharges", Min: types.Pointer[float64](-8388608), Max: types.Pointer[float64](8388607) , Required: true},
			&core.TextField{Name: "paymentMethod"},
			&core.TextField{Name: "shippingMethod"},
			&core.NumberField{Name: "subtotal", Min: types.Pointer[float64](-8388608), Max: types.Pointer[float64](8388607) , Required: true},
			&core.NumberField{Name: "codCharges", Min: types.Pointer[float64](-8388608), Max: types.Pointer[float64](8388607) , Required: true},
			&core.NumberField{Name: "tax", Min: types.Pointer[float64](-8388608), Max: types.Pointer[float64](8388607) , Required: true},
			&core.NumberField{Name: "total", Min: types.Pointer[float64](-8388608), Max: types.Pointer[float64](8388607) , Required: true},
			&core.NumberField{Name: "savingAmount", Min: types.Pointer[float64](-8388608), Max: types.Pointer[float64](8388607) , Required: true},
			&core.TextField{Name: "userAuthToken"},
			&core.TextField{Name: "currencyCode"},
			&core.TextField{Name: "currencySymbol"},
			&core.TextField{Name: "shippingRateId"},
			&core.NumberField{Name: "currencyDecimalDigits", Min: types.Pointer[float64](-8388608), Max: types.Pointer[float64](8388607)},
			&core.RelationField{Name: "storeId", CollectionId: storesCol.Id , Required: true},
			&core.AutodateField{Name: "createdAt", OnCreate: true},
			&core.AutodateField{Name: "updatedAt", OnUpdate: true},
			&core.DateField{Name: "deletedAt"},
		)
		if err := app.Save(cartsCol); err != nil { return err }
		return nil
	}, nil)
}
