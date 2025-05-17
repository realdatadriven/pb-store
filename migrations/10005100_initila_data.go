// Migration for collection: /api/enquiry enquiry - Enquiry - Enquiry created successfully
package migrations

import (
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/migrations"
	"github.com/realdatadriven/pocket_store/internals/env"
)

func init() {
	migrations.Register(func(app core.App) error {

		user, err := app.FindAuthRecordByEmail("users", env.GetString("PB_INIT_USER_EMAIL", ""))
		if err != nil {
			collection, err := app.FindCollectionByNameOrId("users")
			if err != nil {
				return err
			} else {
				user = core.NewRecord(collection)
				user.Set("email", env.GetString("PB_INIT_USER_EMAIL", "test@domain.com"))
				user.Set("fisrtName", env.GetString("PB_INIT_USER_FIRST_NAME", "test"))
				user.Set("lastName", env.GetString("PB_INIT_USER_LAST_NAME", "test"))
				user.Set("role", "admin")
				user.SetPassword(env.GetString("PB_INIT_USER_PASSWOED", "test1234"))
				err = app.Save(user)
				if err != nil {
					return err
				}
			}
		}

		store, err := app.FindFirstRecordByData("stores", "name", env.GetString("PB_INIT_STORE_NAME", "test"))
		if err != nil {
			collection, err := app.FindCollectionByNameOrId("stores")
			if err != nil {
				return err
			} else {
				store = core.NewRecord(collection)
				store.Set("name", env.GetString("PB_INIT_STORE_NAME", "test"))
				store.Set("address_1", env.GetString("PB_INIT_STORE_ADRESS", "test"))
				store.Set("isActive", true)
				store.Set("isApproved", true)
				store.Set("userId", user.Id)
				err = app.Save(store)
				if err != nil {
					return err
				}
			}
		}

		vendor, err := app.FindFirstRecordByData("vendors", "businessName", env.GetString("PB_INIT_VENDOR_NAME", "test"))
		if err != nil {
			collection, err := app.FindCollectionByNameOrId("vendors")
			if err != nil {
				return err
			} else {
				vendor = core.NewRecord(collection)
				vendor.Set("businessName", env.GetString("PB_INIT_VENDOR_NAME", "test"))
				vendor.Set("email", env.GetString("PB_INIT_STORE_EMAIL", "test"))
				vendor.Set("address", env.GetString("PB_INIT_VENDOR_ADRESS", "test"))
				vendor.Set("storeId", store.Id)
				vendor.Set("userId", user.Id)
				err = app.Save(vendor)
				if err != nil {
					return err
				}
			}
		}

		category, err := app.FindFirstRecordByData("categories", "name", env.GetString("PB_INIT_CATEGORY", "test"))
		if err != nil {
			collection, err := app.FindCollectionByNameOrId("categories")
			if err != nil {
				return err
			} else {
				category = core.NewRecord(collection)
				category.Set("name", env.GetString("PB_INIT_CATEGORY", "test"))
				category.Set("storeId", store.Id)
				category.Set("userId", user.Id)
				err = app.Save(category)
				if err != nil {
					return err
				}
			}
		}

		product, err := app.FindFirstRecordByData("products", "name", env.GetString("PB_INIT_PRODUCT_TITLE", "test"))
		if err != nil {
			collection, err := app.FindCollectionByNameOrId("products")
			if err != nil {
				return err
			} else {
				product = core.NewRecord(collection)
				product.Set("title", env.GetString("PB_INIT_PRODUCT_TITLE", "test"))
				product.Set("status", env.GetString("PB_INIT_PRODUCT_STATUS", "published"))
				product.Set("currency", env.GetString("PB_INIT_PRODUCT_CURRENCY", "USD"))
				product.Set("currency", env.GetInt("PB_INIT_PRODUCT_PRICE", 10))
				product.Set("allowBackorder", true)
				product.Set("manageInventory", true)
				product.Set("categoryId", category.Id)
				product.Set("vendorId", vendor.Id)
				product.Set("storeId", store.Id)
				product.Set("userId", user.Id)
				err = app.Save(product)
				if err != nil {
					return err
				}
			}
		}

		return nil
	}, nil)
}
