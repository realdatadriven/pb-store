package migrations

import (
	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/tools/types"
)

func init() {
	m.Register(func(app core.App) error {
		usersCol, err := app.FindCollectionByNameOrId("users")
		if err != nil {
			return err
		}
		collection, err := app.FindCollectionByNameOrId("stores")
		if err != nil {
			collection = core.NewBaseCollection("stores")
		}
		// set rules
		collection.ViewRule = types.Pointer("@request.auth.id != '' && @request.body.user = @request.auth.id")
		collection.CreateRule = types.Pointer("@request.auth.id != '' && @request.body.user = @request.auth.id")
		collection.UpdateRule = types.Pointer(`
			@request.auth.id != '' &&
			(@request.body.user:isset = false || @request.body.user = @request.auth.id)
		`)
		collection.Fields.Add(
			&core.TextField{Name: "name", Required: true},
			&core.RelationField{Name: "userId", CollectionId: usersCol.Id, Required: true},
			&core.AutodateField{Name: "createdAt", OnCreate: true},
			&core.AutodateField{Name: "updatedAt", OnUpdate: true},
		)
		if err := app.Save(collection); err != nil {
			return err
		}
		return nil
	}, nil)
}
