package commands

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
	"github.com/wisnuragaprawida/project/bootstrap"
	"github.com/wisnuragaprawida/project/pkg/log"
)

func init() {
	registerCommand(migrateDown)
}

func migrateDown(dep *bootstrap.Dependency) *cobra.Command {
	return &cobra.Command{
		Use:   "migrate:down",
		Short: "migrate down",
		Run: func(cmd *cobra.Command, args []string) {
			log.Info("rollback process started")
			m, err := migrate.New("file://files/db_migrations/", dep.GetConfig().Database.Write)
			if err != nil {
				log.Fatal(err)
			}
			err = m.Down()
			if err != nil {
				log.Fatal(err)
			}
			log.Info("rollback process finished")
		},
	}
}
