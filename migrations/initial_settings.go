package migrations

import (
	"litestore/internals/env"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		settings := app.Settings()

		// for all available settings fields you could check
		// https://github.com/pocketbase/pocketbase/blob/develop/core/settings_model.go#L121-L130
		settings.Meta.AppName = env.GetString("PB_META_APP_NAME", "test")
		settings.Meta.SenderName = env.GetString("PB_META_SENDER_NAME", "Support")
		settings.Meta.SenderAddress = env.GetString("PB_META_SENDER_ADRESS", "Support")
		settings.Meta.AppURL = "https://example.com"
		settings.Logs.MaxDays = env.GetInt("PB_LOGS_MAX_DAYS", 2)
		settings.Logs.LogAuthId = env.GetBool("PB_LOGS_AUTH_ID", false)
		settings.Logs.LogIP = env.GetBool("PB_LOGS_IP", false)

		// SMTP
		settings.SMTP.Enabled = env.GetBool("PB_SMTP_ENABLED", false)
		settings.SMTP.Host = env.GetString("PB_SMTP_HOST", "")
		settings.SMTP.Port = env.GetInt("PB_SMTP_PORT", 587)
		settings.SMTP.Username = env.GetString("PB_SMTP_USER", "")
		settings.SMTP.Password = env.GetString("PB_SMTP_PASS", "")
		//settings.SMTP.LocalName = env.GetString("PB_SMTP_LOCAL_NAME", "")

		// BACKUPS
		settings.Backups.Cron = env.GetString("PB_BACKUPS_CRON", "0 0 * * *")
		settings.Backups.CronMaxKeep = env.GetInt("PB_BACKUPS_MAX_KEEP", 3)
		// BACKUPS S3
		settings.Backups.S3.Enabled = env.GetBool("PB_S3_BACKUPS_ENABLED", false)
		settings.Backups.S3.Endpoint = env.GetString("PB_S3_BACKUPS_ENDPOINT", "")
		settings.Backups.S3.Bucket = env.GetString("PB_S3_BACKUPS_BUCKET", "")
		settings.Backups.S3.Region = env.GetString("PB_S3_BACKUPS_REGION", "")
		settings.Backups.S3.AccessKey = env.GetString("PB_S3_BACKUPS_ACCESS_KEY", "")
		settings.Backups.S3.Secret = env.GetString("PB_S3_BACKUPS_SECRET", "")
		settings.Backups.S3.ForcePathStyle = env.GetBool("PB_S3_BACKUPS_FORCE_PATH_STYLE", false)

		// S3
		settings.S3.Enabled = env.GetBool("PB_S3_ENABLED", false)
		settings.S3.Endpoint = env.GetString("PB_S3_ENDPOINT", "")
		settings.S3.Bucket = env.GetString("PB_S3_BUCKET", "")
		settings.S3.Region = env.GetString("PB_S3_REGION", "")
		settings.S3.AccessKey = env.GetString("PB_S3_ACCESS_KEY", "")
		settings.S3.Secret = env.GetString("PB_S3_SECRET", "")
		settings.S3.ForcePathStyle = env.GetBool("PB_S3_FORCE_PATH_STYLE", false)
		return app.Save(settings)
	}, nil)
}
