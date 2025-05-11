package migrations

import (
	"litestore/internals/env"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	if env.GetString("PB_SU_EMAIL", "") != "" && env.GetString("PB_SU_PASS", "") != "" {
		m.Register(func(app core.App) error {
			superusers, err := app.FindCollectionByNameOrId(core.CollectionNameSuperusers)
			if err != nil {
				return err
			}
			record := core.NewRecord(superusers)
			record.Set("email", env.GetString("PB_SU_EMAIL", ""))
			record.Set("password", env.GetString("PB_SU_PASS", ""))
			return app.Save(record)
		}, nil)
	}
}
