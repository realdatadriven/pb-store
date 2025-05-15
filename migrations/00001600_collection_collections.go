// Migration for collection: /api/collections collections - Products - Successfully retrieved collection of product
package migrations
import (
    "github.com/pocketbase/pocketbase/core"
    "github.com/pocketbase/pocketbase/migrations"
    "github.com/pocketbase/pocketbase/tools/types"
)

func init() {
migrations.Register(func(app core.App) error {
		
		collectionsCol, err := app.FindCollectionByNameOrId("collections")
		if err != nil { collectionsCol = core.NewBaseCollection("collections") }
		collectionsCol.ListRule = types.Pointer("")
		collectionsCol.ViewRule = types.Pointer("")
		collectionsCol.Fields.Add(

		)
		if err := app.Save(collectionsCol); err != nil { return err }
		return nil
	}, nil)
}
