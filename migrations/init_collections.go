
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


	m.Register(func(app core.App) error {
		usersCol, err := app.FindCollectionByNameOrId("users")
		if err != nil {
			return err
		}
		storesCol, err := app.FindCollectionByNameOrId("storesCol")
		if err != nil { return err }

		collection, err := app.FindCollectionByNameOrId("vendors")
		if err != nil {
			collection = core.NewBaseCollection("vendors")
		}

		collection.Fields.Add(
			&core.TextField{Name: "name"},
			&core.TextField{Name: "description"},
			&core.NumberField{Name: "price", Min: types.Pointer(0.0), Max: types.Pointer(999999.99)},
			&core.NumberField{Name: "quantity", Min: types.Pointer(0.0)},
			&core.BoolField{Name: "isPublished"},
			&core.RelationField{Name: "userId", CollectionId: usersCol.Id, Required: false},
			&core.RelationField{Name: "storeId", CollectionId: storesCol.Id, Required: false},
			&core.AutodateField{Name: "createdAt", OnCreate: true},
			&core.AutodateField{Name: "updatedAt", OnUpdate: true},
		)

		if err := app.Save(collection); err != nil {
			return err
		}
		return nil
	}, nil)


	m.Register(func(app core.App) error {
		usersCol, err := app.FindCollectionByNameOrId("users")
		if err != nil {
			return err
		}
		storesCol, err := app.FindCollectionByNameOrId("storesCol")
		if err != nil { return err }

		collection, err := app.FindCollectionByNameOrId("categories")
		if err != nil {
			collection = core.NewBaseCollection("categories")
		}

		collection.Fields.Add(
			&core.TextField{Name: "name", Required: true},
			&core.RelationField{Name: "userId", CollectionId: usersCol.Id, Required: false},
			&core.RelationField{Name: "storeId", CollectionId: storesCol.Id, Required: false},
			&core.AutodateField{Name: "createdAt", OnCreate: true},
			&core.AutodateField{Name: "updatedAt", OnUpdate: true},
		)

		if err := app.Save(collection); err != nil {
			return err
		}
		return nil
	}, nil)


	m.Register(func(app core.App) error {
		usersCol, err := app.FindCollectionByNameOrId("users")
		if err != nil {
			return err
		}
		storesCol, err := app.FindCollectionByNameOrId("storesCol")
		if err != nil { return err }
		vendorsCol, err := app.FindCollectionByNameOrId("vendorsCol")
		if err != nil { return err }
		categoriesCol, err := app.FindCollectionByNameOrId("categoriesCol")
		if err != nil { return err }

		collection, err := app.FindCollectionByNameOrId("products")
		if err != nil {
			collection = core.NewBaseCollection("products")
		}

		collection.Fields.Add(
			&core.TextField{Name: "title", Required: true},
			&core.NumberField{Name: "price", Min: types.Pointer(0.0), Max: types.Pointer(999999.99)},
			&core.TextField{Name: "sku"},
			&core.RelationField{Name: "userId", CollectionId: usersCol.Id, Required: false},
			&core.RelationField{Name: "storeId", CollectionId: storesCol.Id, Required: false},
			&core.RelationField{Name: "vendorId", CollectionId: vendorsCol.Id, Required: false},
			&core.RelationField{Name: "categoryId", CollectionId: categoriesCol.Id, Required: false},
			&core.AutodateField{Name: "createdAt", OnCreate: true},
			&core.AutodateField{Name: "updatedAt", OnUpdate: true},
		)

		if err := app.Save(collection); err != nil {
			return err
		}
		return nil
	}, nil)

}
