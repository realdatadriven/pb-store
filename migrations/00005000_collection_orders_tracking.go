// Migration for collection: /api/admin/orders/tracking tracking - Orders - Order tracking successfully created
package migrations
import (
    "github.com/pocketbase/pocketbase/core"
    "github.com/pocketbase/pocketbase/migrations"
    "github.com/pocketbase/pocketbase/tools/types"
)

func init() {
migrations.Register(func(app core.App) error {
		cartsCol, _ := app.FindCollectionByNameOrId("carts")
vendorsCol, _ := app.FindCollectionByNameOrId("vendors")
storesCol, _ := app.FindCollectionByNameOrId("stores")
usersCol, _ := app.FindCollectionByNameOrId("users")

		trackingCol, err := app.FindCollectionByNameOrId("tracking")
		if err != nil { trackingCol = core.NewBaseCollection("tracking") }
		trackingCol.ListRule = types.Pointer("")
		trackingCol.ViewRule = types.Pointer("")
		trackingCol.Fields.Add(
			&core.NumberField{Name: "orderNo", Min: types.Pointer[float64](-2147483648), Max: types.Pointer[float64](2147483647) , Required: true},
			&core.TextField{Name: "batchNo"},
			&core.TextField{Name: "parentOrderNo"},
			&core.NumberField{Name: "otp", Min: types.Pointer[float64](-8388608), Max: types.Pointer[float64](8388607)},
			&core.BoolField{Name: "isEmailSentToVendor"},
			&core.TextField{Name: "status"},
			&core.RelationField{Name: "cartId", CollectionId: cartsCol.Id , Required: true},
			&core.TextField{Name: "userPhone"},
			&core.TextField{Name: "userFirstName"},
			&core.TextField{Name: "userLastName"},
			&core.TextField{Name: "userEmail"},
			&core.TextField{Name: "comment"},
			&core.BoolField{Name: "needAddress" , Required: true},
			&core.BoolField{Name: "selfTakeout" , Required: true},
			&core.NumberField{Name: "shippingCharges", Min: types.Pointer[float64](-8388608), Max: types.Pointer[float64](8388607)},
			&core.NumberField{Name: "total", Min: types.Pointer[float64](-8388608), Max: types.Pointer[float64](8388607)},
			&core.NumberField{Name: "subtotal", Min: types.Pointer[float64](-8388608), Max: types.Pointer[float64](8388607)},
			&core.NumberField{Name: "discount", Min: types.Pointer[float64](-8388608), Max: types.Pointer[float64](8388607)},
			&core.NumberField{Name: "tax", Min: types.Pointer[float64](-8388608), Max: types.Pointer[float64](8388607)},
			&core.TextField{Name: "currencySymbol"},
			&core.TextField{Name: "currencyCode"},
			&core.NumberField{Name: "codCharges", Min: types.Pointer[float64](-8388608), Max: types.Pointer[float64](8388607)},
			&core.NumberField{Name: "codPaid", Min: types.Pointer[float64](-8388608), Max: types.Pointer[float64](8388607)},
			&core.BoolField{Name: "paid"},
			&core.NumberField{Name: "paySuccess", Min: types.Pointer[float64](-2147483648), Max: types.Pointer[float64](2147483647)},
			&core.NumberField{Name: "amountRefunded", Min: types.Pointer[float64](-8388608), Max: types.Pointer[float64](8388607)},
			&core.NumberField{Name: "amountDue", Min: types.Pointer[float64](-8388608), Max: types.Pointer[float64](8388607)},
			&core.NumberField{Name: "amountPaid", Min: types.Pointer[float64](-8388608), Max: types.Pointer[float64](8388607)},
			&core.NumberField{Name: "totalDiscount", Min: types.Pointer[float64](-8388608), Max: types.Pointer[float64](8388607)},
			&core.NumberField{Name: "totalAmountRefunded", Min: types.Pointer[float64](-8388608), Max: types.Pointer[float64](8388607)},
			&core.TextField{Name: "paymentMethod"},
			&core.TextField{Name: "platform"},
			&core.TextField{Name: "couponUsed"},
			&core.TextField{Name: "coupon"},
			&core.TextField{Name: "paymentStatus"},
			&core.TextField{Name: "paymentCurrency"},
			&core.TextField{Name: "paymentMsg"},
			&core.TextField{Name: "paymentReferenceId"},
			&core.TextField{Name: "paymentGateway"},
			&core.TextField{Name: "paymentId"},
			&core.NumberField{Name: "paymentAmount", Min: types.Pointer[float64](-8388608), Max: types.Pointer[float64](8388607)},
			&core.TextField{Name: "paymentMode"},
			&core.TextField{Name: "paymentDate"},
			&core.TextField{Name: "shippingAddressId"},
			&core.TextField{Name: "billingAddressId"},
			&core.TextField{Name: "invoiceUrl"},
			&core.RelationField{Name: "vendorId", CollectionId: vendorsCol.Id},
			&core.RelationField{Name: "storeId", CollectionId: storesCol.Id , Required: true},
			&core.RelationField{Name: "userId", CollectionId: usersCol.Id},
			&core.AutodateField{Name: "createdAt", OnCreate: true},
			&core.AutodateField{Name: "updatedAt", OnUpdate: true},
		)
		if err := app.Save(trackingCol); err != nil { return err }
		return nil
	}, nil)
}
