package migrations

import (
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/migrations"
)

func init() {
	migrations.Register(func(app core.App) error {
		// Find the users collection
		usersCol, err := app.FindCollectionByNameOrId("users")
		if err != nil {
			return err
		}
		// Check if the firstName field already exists
		firstNameField := usersCol.Fields.GetByName("firstName")
		if firstNameField == nil {
			// Add the firstName field to the users collection
			usersCol.Fields.Add(
				&core.TextField{
					Name: "firstName",
				},
			)
		}
		// Check if the lastName field already exists
		lastNameField := usersCol.Fields.GetByName("lastName")
		if lastNameField == nil {
			// Add the lastName field to the users collection
			usersCol.Fields.Add(
				&core.TextField{
					Name: "lastName",
				},
			)
		}
		timeDiffField := usersCol.Fields.GetByName("timeDiff")
		if timeDiffField == nil {
			usersCol.Fields.Add(
				&core.NumberField{
					Name: "timeDiff",
				},
			)
		}
		phoneField := usersCol.Fields.GetByName("phone")
		if phoneField == nil {
			usersCol.Fields.Add(
				&core.TextField{
					Name: "phone",
				},
			)
		}
		// Check if the role field already exists
		roleField := usersCol.Fields.GetByName("role")
		if roleField == nil {
			// Add the role field to the users collection
			usersCol.Fields.Add(
				&core.SelectField{
					Name:   "role",
					Values: []string{"superuser", "admin", "user"},
					//Default: "user",
				},
			)
		}
		// Save the updated users collection
		if err := app.Save(usersCol); err != nil {
			return err
		}
		return nil
	}, nil)
}
