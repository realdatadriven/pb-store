// Migration for collection: /api/enquiry enquiry - Enquiry - Enquiry created successfully
package migrations
import (
    "github.com/pocketbase/pocketbase/core"
    "github.com/pocketbase/pocketbase/migrations"
    "github.com/pocketbase/pocketbase/tools/types"
)

func init() {
migrations.Register(func(app core.App) error {
		productsCol, _ := app.FindCollectionByNameOrId("products")
usersCol, _ := app.FindCollectionByNameOrId("users")
storesCol, _ := app.FindCollectionByNameOrId("stores")

		enquiryCol, err := app.FindCollectionByNameOrId("enquiry")
		if err != nil { enquiryCol = core.NewBaseCollection("enquiry") }
		enquiryCol.ListRule = types.Pointer("")
		enquiryCol.ViewRule = types.Pointer("")
		enquiryCol.Fields.Add(
			&core.RelationField{Name: "productId", CollectionId: productsCol.Id , Required: true},
			&core.RelationField{Name: "userId", CollectionId: usersCol.Id},
			&core.RelationField{Name: "storeId", CollectionId: storesCol.Id , Required: true},
			&core.TextField{Name: "name" , Required: true},
			&core.EmailField{Name: "email" , Required: true},
			&core.TextField{Name: "phone"},
			&core.TextField{Name: "message" , Required: true},
			&core.TextField{Name: "status"},
			&core.BoolField{Name: "notificationSent"},
			&core.TextField{Name: "errorMessage"},
			&core.TextField{Name: "errorStack"},
			&core.TextField{Name: "processingTime"},
			&core.TextField{Name: "metadata"},
			&core.AutodateField{Name: "createdAt", OnCreate: true},
			&core.AutodateField{Name: "updatedAt", OnUpdate: true},
		)
		if err := app.Save(enquiryCol); err != nil { return err }
		return nil
	}, nil)
}
