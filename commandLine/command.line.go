package commandline

import (
	"github.com/pius706975/backend/database"
	libs "github.com/pius706975/backend/libs/server"
	"github.com/spf13/cobra"
)

var InitCommand = cobra.Command{
	Short: "on backend",
	Long:  "online notes backend",
}

func init() {
	InitCommand.AddCommand(libs.ServeCMD)
	InitCommand.AddCommand(database.MigrateCMD)
}

func Run(args []string) error {
	InitCommand.SetArgs(args)

	return InitCommand.Execute()
}
