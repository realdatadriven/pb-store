// Migration for collection: /api/currencies currencies - Currencies - Successfully retrieved list of currencies with pagination details
package migrations
import (
    "github.com/pocketbase/pocketbase/core"
    "github.com/pocketbase/pocketbase/migrations"
    "github.com/pocketbase/pocketbase/tools/types"
)

func init() {
migrations.Register(func(app core.App) error {
		
		currenciesCol, err := app.FindCollectionByNameOrId("currencies")
		if err != nil { currenciesCol = core.NewBaseCollection("currencies") }
		currenciesCol.ListRule = types.Pointer("")
		currenciesCol.ViewRule = types.Pointer("")
		currenciesCol.Fields.Add(
			&core.TextField{Name: "code" , Required: true},
			&core.TextField{Name: "decimalDigits" , Required: true},
			&core.TextField{Name: "rounding" , Required: true},
			&core.TextField{Name: "namePlural" , Required: true},
			&core.TextField{Name: "symbol" , Required: true},
			&core.TextField{Name: "symbolNative" , Required: true},
			&core.TextField{Name: "name" , Required: true},
			&core.BoolField{Name: "includesTax"},
		)
		if err := app.Save(currenciesCol); err != nil { return err }
		return nil
	}, nil)
}
