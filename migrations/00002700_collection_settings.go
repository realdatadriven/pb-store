// Migration for collection: /api/settings settings - Settings - Successfully retrieved settings
package migrations
import (
    "github.com/pocketbase/pocketbase/core"
    "github.com/pocketbase/pocketbase/migrations"
    "github.com/pocketbase/pocketbase/tools/types"
)

func init() {
migrations.Register(func(app core.App) error {
		
		settingsCol, err := app.FindCollectionByNameOrId("settings")
		if err != nil { settingsCol = core.NewBaseCollection("settings") }
		settingsCol.ListRule = types.Pointer("")
		settingsCol.ViewRule = types.Pointer("")
		settingsCol.Fields.Add(
			&core.TextField{Name: "name" , Required: true},
			&core.TextField{Name: "description"},
			&core.FileField{Name: "logo"},
			&core.TextField{Name: "address_1"},
			&core.TextField{Name: "address_2"},
			&core.TextField{Name: "city"},
			&core.TextField{Name: "state"},
			&core.TextField{Name: "country"},
			&core.TextField{Name: "phone"},
			&core.EmailField{Name: "email"},
			&core.TextField{Name: "zipCode"},
			&core.TextField{Name: "defaultStoreId"},
			&core.TextField{Name: "certificateServerIp"},
			&core.AutodateField{Name: "createdAt", OnCreate: true},
			&core.AutodateField{Name: "updatedAt", OnUpdate: true},
		)
		if err := app.Save(settingsCol); err != nil { return err }
		return nil
	}, nil)
}
