// Migration for collection: /api/returns returns - Returns - Return request successfully created
package migrations
import (
    "github.com/pocketbase/pocketbase/core"
    "github.com/pocketbase/pocketbase/migrations"
    "github.com/pocketbase/pocketbase/tools/types"
)

func init() {
migrations.Register(func(app core.App) error {
		
		returnsCol, err := app.FindCollectionByNameOrId("returns")
		if err != nil { returnsCol = core.NewBaseCollection("returns") }
		returnsCol.ListRule = types.Pointer("")
		returnsCol.ViewRule = types.Pointer("")
		returnsCol.Fields.Add(
			&core.TextField{Name: "message" , Required: true},
		)
		if err := app.Save(returnsCol); err != nil { return err }
		return nil
	}, nil)
}
