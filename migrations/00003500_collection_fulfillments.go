// Migration for collection: /api/admin/fulfillments fulfillments - Fulfillments - Successfully retrieved list of fulfillments with pagination details
package migrations

import (
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/tools/types"
)

func init() {
	migrations.Register(func(app core.App) error {
		ordersCol, _ := app.FindCollectionByNameOrId("orders")
		vendorsCol, _ := app.FindCollectionByNameOrId("vendors")
		storesCol, _ := app.FindCollectionByNameOrId("stores")
		usersCol, _ := app.FindCollectionByNameOrId("users")

		fulfillmentsCol, err := app.FindCollectionByNameOrId("fulfillments")
		if err != nil {
			fulfillmentsCol = core.NewBaseCollection("fulfillments")
		}
		fulfillmentsCol.ListRule = types.Pointer("")
		fulfillmentsCol.ViewRule = types.Pointer("")
		fulfillmentsCol.Fields.Add(
			&core.BoolField{Name: "shippingSync"},
			&core.TextField{Name: "store"},
			&core.BoolField{Name: "active"},
			&core.NumberField{Name: "orderNo", Min: types.Pointer[float64](-2147483648), Max: types.Pointer[float64](2147483647), Required: true},
			&core.TextField{Name: "trackingNumber"},
			&core.TextField{Name: "trackingUrl"},
			&core.RelationField{Name: "orderId", CollectionId: ordersCol.Id},
			&core.TextField{Name: "batchNo", Required: true},
			&core.TextField{Name: "fulfillmentOrderId"},
			&core.TextField{Name: "shipmentId"},
			&core.TextField{Name: "shippingProvider"},
			&core.TextField{Name: "shipmentLabel"},
			&core.TextField{Name: "invoiceUrl"},
			&core.TextField{Name: "courierName"},
			&core.TextField{Name: "courierId"},
			&core.TextField{Name: "shippingStatus"},
			&core.TextField{Name: "status"},
			&core.TextField{Name: "shippingInfo"},
			&core.TextField{Name: "manifest"},
			&core.NumberField{Name: "weight"},
			&core.NumberField{Name: "length"},
			&core.NumberField{Name: "breadth"},
			&core.NumberField{Name: "height"},
			&core.RelationField{Name: "vendorId", CollectionId: vendorsCol.Id},
			&core.RelationField{Name: "storeId", CollectionId: storesCol.Id, Required: true},
			&core.RelationField{Name: "userId", CollectionId: usersCol.Id},
			&core.AutodateField{Name: "createdAt", OnCreate: true},
			&core.AutodateField{Name: "updatedAt", OnUpdate: true},
		)
		if err := app.Save(fulfillmentsCol); err != nil {
			return err
		}
		return nil
	}, nil)
}
