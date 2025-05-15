// Migration for collection: /api/chats chats - Chats - The requested chat
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

		chatsCol, err := app.FindCollectionByNameOrId("chats")
		if err != nil { chatsCol = core.NewBaseCollection("chats") }
		chatsCol.ListRule = types.Pointer("")
		chatsCol.ViewRule = types.Pointer("")
		chatsCol.Fields.Add(
			&core.TextField{Name: "customerId" , Required: true},
			&core.RelationField{Name: "storeId", CollectionId: storesCol.Id , Required: true},
			&core.RelationField{Name: "userId", CollectionId: usersCol.Id},
			&core.AutodateField{Name: "createdAt", OnCreate: true},
			&core.AutodateField{Name: "updatedAt", OnUpdate: true},
		)
		if err := app.Save(chatsCol); err != nil { return err }
		return nil
	}, nil)
}
