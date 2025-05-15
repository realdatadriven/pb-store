// Migration for collection: /api/admin/templates templates - Templates - The list of templates
package migrations
import (
    "github.com/pocketbase/pocketbase/core"
    "github.com/pocketbase/pocketbase/migrations"
    "github.com/pocketbase/pocketbase/tools/types"
)

func init() {
migrations.Register(func(app core.App) error {
		storesCol, _ := app.FindCollectionByNameOrId("stores")
usersCol, _ := app.FindCollectionByNameOrId("users")

		templatesCol, err := app.FindCollectionByNameOrId("templates")
		if err != nil { templatesCol = core.NewBaseCollection("templates") }
		templatesCol.ListRule = types.Pointer("")
		templatesCol.ViewRule = types.Pointer("")
		templatesCol.Fields.Add(
			&core.TextField{Name: "templateId" , Required: true},
			&core.TextField{Name: "dltTemplateId"},
			&core.TextField{Name: "subject"},
			&core.TextField{Name: "fromEmail"},
			&core.TextField{Name: "toEmail"},
			&core.TextField{Name: "description"},
			&core.EditorField{Name: "content"},
			&core.TextField{Name: "sampleData"},
			&core.TextField{Name: "type"},
			&core.BoolField{Name: "isActive"},
			&core.RelationField{Name: "storeId", CollectionId: storesCol.Id , Required: true},
			&core.RelationField{Name: "userId", CollectionId: usersCol.Id},
			&core.AutodateField{Name: "createdAt", OnCreate: true},
			&core.AutodateField{Name: "updatedAt", OnUpdate: true},
		)
		if err := app.Save(templatesCol); err != nil { return err }
		return nil
	}, nil)
}
