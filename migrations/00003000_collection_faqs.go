// Migration for collection: /api/faqs faqs - Faqs - The list of faqs
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

		faqsCol, err := app.FindCollectionByNameOrId("faqs")
		if err != nil {
			faqsCol = core.NewBaseCollection("faqs")
		}
		faqsCol.ListRule = types.Pointer("")
		faqsCol.ViewRule = types.Pointer("")
		faqsCol.Fields.Add(
			&core.TextField{Name: "question"},
			&core.TextField{Name: "answer"},
			&core.TextField{Name: "zip"},
			&core.TextField{Name: "status"},
			&core.TextField{Name: "topic"},
			&core.NumberField{Name: "rank"},
			&core.RelationField{Name: "storeId", CollectionId: storesCol.Id, Required: true},
			&core.RelationField{Name: "userId", CollectionId: usersCol.Id},
			&core.AutodateField{Name: "createdAt", OnCreate: true},
			&core.AutodateField{Name: "updatedAt", OnUpdate: true},
		)
		if err := app.Save(faqsCol); err != nil {
			return err
		}
		return nil
	}, nil)
}
