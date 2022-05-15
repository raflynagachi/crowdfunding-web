package cmd

import (
	"github.com/raflynagachi/crowdfunding-web/database"
	"github.com/spf13/cobra"
)

var dbMigrate = &cobra.Command{
	Use:   "db:migrate",
	Short: "do database migration",
	Run: func(cmd *cobra.Command, args []string) {

		database.MigrateDB(false)
	},
}

var dbMigrateFresh = &cobra.Command{
	Use:   "db:migrate-fresh",
	Short: "do database fresh migration",
	Run: func(cmd *cobra.Command, args []string) {
		database.MigrateDB(true)
	},
}

var dbSeed = &cobra.Command{
	Use:   "db:seed",
	Short: "do database seeding",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO
		// run migration command
		panic("not implemented - seed")
	},
}

func init() {
	RootCmd.AddCommand(dbMigrate, dbMigrateFresh, dbSeed)
}
