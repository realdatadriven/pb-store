// Migration for collection: /api/countries countries - Countries - Successfully retrieved list of countries with pagination details
package migrations
import (
    "github.com/pocketbase/pocketbase/core"
    "github.com/pocketbase/pocketbase/migrations"
    "github.com/pocketbase/pocketbase/tools/types"
)

func init() {
migrations.Register(func(app core.App) error {
		
		countriesCol, err := app.FindCollectionByNameOrId("countries")
		if err != nil { countriesCol = core.NewBaseCollection("countries") }
		countriesCol.ListRule = types.Pointer("")
		countriesCol.ViewRule = types.Pointer("")
		countriesCol.Fields.Add(
			&core.TextField{Name: "iso2" , Required: true},
			&core.TextField{Name: "code"},
			&core.TextField{Name: "iso3" , Required: true},
			&core.NumberField{Name: "numCode", Min: types.Pointer[float64](-2147483648), Max: types.Pointer[float64](2147483647) , Required: true},
			&core.TextField{Name: "dialCode" , Required: true},
			&core.TextField{Name: "name" , Required: true},
			&core.TextField{Name: "displayName"},
			&core.TextField{Name: "regionId"},
		)
		if err := app.Save(countriesCol); err != nil { return err }
		return nil
	}, nil)
}
