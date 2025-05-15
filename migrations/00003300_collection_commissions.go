// Migration for collection: /api/admin/commissions commissions - Invoices - Successfully retrieved commission details
package migrations
import (
    "github.com/pocketbase/pocketbase/core"
    "github.com/pocketbase/pocketbase/migrations"
    "github.com/pocketbase/pocketbase/tools/types"
)

func init() {
migrations.Register(func(app core.App) error {
		
		commissionsCol, err := app.FindCollectionByNameOrId("commissions")
		if err != nil { commissionsCol = core.NewBaseCollection("commissions") }
		commissionsCol.ListRule = types.Pointer("")
		commissionsCol.ViewRule = types.Pointer("")
		commissionsCol.Fields.Add(

		)
		if err := app.Save(commissionsCol); err != nil { return err }
		return nil
	}, nil)
}
