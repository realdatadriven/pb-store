// Migration for collection: /api/admin/stores stores - Store - Store retrieved successfully
package migrations
import (
    "github.com/pocketbase/pocketbase/core"
    "github.com/pocketbase/pocketbase/migrations"
    "github.com/pocketbase/pocketbase/tools/types"
)

func init() {
migrations.Register(func(app core.App) error {
		usersCol, _ := app.FindCollectionByNameOrId("users")

		storesCol, err := app.FindCollectionByNameOrId("stores")
		if err != nil { storesCol = core.NewBaseCollection("stores") }
		storesCol.ListRule = types.Pointer("")
		storesCol.ViewRule = types.Pointer("")
		storesCol.Fields.Add(
			&core.TextField{Name: "name" , Required: true},
			&core.TextField{Name: "homepageTitle"},
			&core.TextField{Name: "slug"},
			&core.BoolField{Name: "isActive"},
			&core.BoolField{Name: "isApproved"},
			&core.BoolField{Name: "isFeatured"},
			&core.BoolField{Name: "isClosed"},
			&core.TextField{Name: "ownerFirstName"},
			&core.TextField{Name: "ownerLastName"},
			&core.TextField{Name: "ownerPhone"},
			&core.TextField{Name: "ownerEmail"},
			&core.TextField{Name: "address_1" , Required: true},
			&core.TextField{Name: "address_2"},
			&core.TextField{Name: "zipCode"},
			&core.TextField{Name: "city"},
			&core.TextField{Name: "state"},
			&core.TextField{Name: "country"},
			&core.NumberField{Name: "lat", Min: types.Pointer[float64](-8388608), Max: types.Pointer[float64](8388607)},
			&core.NumberField{Name: "lng", Min: types.Pointer[float64](-8388608), Max: types.Pointer[float64](8388607)},
			&core.TextField{Name: "currencyCode"},
			&core.TextField{Name: "currencySymbol"},
			&core.TextField{Name: "currencySymbolNative"},
			&core.TextField{Name: "currencyDecimalDigits"},
			&core.TextField{Name: "lang"},
			&core.TextField{Name: "description"},
			&core.TextField{Name: "metaTitle"},
			&core.TextField{Name: "metaDescription"},
			&core.TextField{Name: "metaKeywords"},
			&core.NumberField{Name: "freeShippingOn", Min: types.Pointer[float64](-8388608), Max: types.Pointer[float64](8388607)},
			&core.NumberField{Name: "minimumOrderValue", Min: types.Pointer[float64](-8388608), Max: types.Pointer[float64](8388607)},
			&core.NumberField{Name: "shippingCharges", Min: types.Pointer[float64](-8388608), Max: types.Pointer[float64](8388607)},
			&core.NumberField{Name: "commissionRate", Min: types.Pointer[float64](-8388608), Max: types.Pointer[float64](8388607)},
			&core.TextField{Name: "timeZone"},
			&core.TextField{Name: "businessPhone"},
			&core.TextField{Name: "businessEmail"},
			&core.TextField{Name: "businessLegalName"},
			&core.TextField{Name: "businessName"},
			&core.FileField{Name: "logo"},
			&core.TextField{Name: "favicon"},
			&core.TextField{Name: "orderPrefix"},
			&core.TextField{Name: "productImageAspectRatio"},
			&core.TextField{Name: "homePageSliderBannerImageHeight"},
			&core.TextField{Name: "androidScheme"},
			&core.TextField{Name: "androidPackage"},
			&core.TextField{Name: "androidAppJson"},
			&core.NumberField{Name: "androidVersionCode", Min: types.Pointer[float64](-2147483648), Max: types.Pointer[float64](2147483647)},
			&core.TextField{Name: "androidWebviewUrl"},
			&core.TextField{Name: "themePrimaryColor"},
			&core.TextField{Name: "themeSecondaryColor"},
			&core.TextField{Name: "themeFontFamily"},
			&core.TextField{Name: "themeFontColor"},
			&core.TextField{Name: "dimensionUnit"},
			&core.TextField{Name: "weightUnit"},
			&core.BoolField{Name: "isParent"},
			&core.RelationField{Name: "userId", CollectionId: usersCol.Id},
			&core.TextField{Name: "sitemap"},
			&core.AutodateField{Name: "createdAt", OnCreate: true},
			&core.AutodateField{Name: "updatedAt", OnUpdate: true},
		)
		if err := app.Save(storesCol); err != nil { return err }
		return nil
	}, nil)
}
