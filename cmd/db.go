package cmd

import (
	"adm/app/models"
	"github.com/gobuffalo/envy"
	"github.com/pressly/goose"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var migrationStep int
var migrationsDir = envy.Get("MIGRATIONS_DIR", "app/migrations")

func init() {
	RootCmd.AddCommand(migrationCommand)
	migrationCommand.AddCommand(statusCommand)

	migrationCommand.AddCommand(migrateUpCommand)
	migrateUpCommand.Flags().IntVarP(&migrationStep, "step", "s", 0, "Number of steps to migrate.")

	migrationCommand.AddCommand(migrateDownCommand)
	migrateDownCommand.Flags().IntVarP(&migrationStep, "step", "s", 0, "Number of steps to migrate.")

	migrationCommand.AddCommand(migrateRedoCommand)
	migrationCommand.AddCommand(migrateCreateCommand)
}

var migrationCommand = &cobra.Command{
	Use:   "migrate",
	Short: "Manage your apps database migrations",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		os.MkdirAll(envy.Get("MIGRATIONS_DIR", "app/migrations"), 0766)
	},
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var statusCommand = &cobra.Command{
	Use:   "status",
	Short: "Get the current status of database migrations",
	Run: func(cmd *cobra.Command, args []string) {
		if err := goose.Run("status", models.DB().DB, migrationsDir); err != nil {
			log.Fatal(err)
		}
	},
}

var migrateUpCommand = &cobra.Command{
	Use:   "up",
	Short: "Migrate the database up. By default, it runs all pending migrations.",
	Run: func(cmd *cobra.Command, args []string) {
		if migrationStep > 0 {
			migrations := getMigrations()

			for i := 0; i < migrationStep; i++ {
				currentVersion := getCurrentVersion()

				next, err := migrations.Next(currentVersion)
				if err != nil {
					if err == goose.ErrNoNextVersion {
						log.Fatalf("no migration %v\n", currentVersion)
					} else {
						log.Fatal(err)
					}
				}

				if err = next.Up(models.DB().DB); err != nil {
					log.Fatal(err)
				}
			}
		} else {
			if err := goose.Run("up", models.DB().DB, migrationsDir); err != nil {
				log.Fatal(err)
			}
		}
	},
}

var migrateDownCommand = &cobra.Command{
	Use:   "down",
	Short: "Migrate the database down. By default, it rolls back the db by 1 migration.",
	Run: func(cmd *cobra.Command, args []string) {
		if migrationStep > 0 {
			migrations := getMigrations()

			for i := 0; i < migrationStep; i++ {
				currentVersion := getCurrentVersion()

				current, err := migrations.Current(currentVersion)
				if err != nil {
					log.Fatalf("no migration %v", currentVersion)
				}

				if err = current.Down(models.DB().DB); err != nil {
					log.Fatal(err)
				}
			}
		} else {
			if err := goose.Run("down", models.DB().DB, migrationsDir); err != nil {
				log.Fatal(err)
			}
		}
	},
}

var migrateRedoCommand = &cobra.Command{
	Use:   "redo",
	Short: "Redo the last migration.",
	Run: func(cmd *cobra.Command, args []string) {
		if err := goose.Run("redo", models.DB().DB, migrationsDir); err != nil {
			log.Fatal(err)
		}
	},
}

var migrateCreateCommand = &cobra.Command{
	Use:   "create",
	Short: "Create a new migration.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatalln("Please specify a name for the migration")
		}

		if err := goose.Create(models.DB().DB, envy.Get("MIGRATIONS_DIR", ""), args[0], "sql"); err != nil {
			log.Fatal(err)
		}
	},
}

func getCurrentVersion() int64 {
	currentVersion, err := goose.GetDBVersion(models.DB().DB)
	if err != nil {
		log.Fatal(err)
	}
	return currentVersion
}

func getMigrations() goose.Migrations {
	migrations, err := goose.CollectMigrations(migrationsDir, int64(0), int64((1<<63)-1))
	if err != nil {
		log.Fatal(err)
	}

	return migrations
}
