package cmd

import (
	"fmt"
	"github.com/jiuzi/sshm/dao"
	"github.com/modood/table"
	"github.com/spf13/cobra"
)

var listDataCommand = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "Show all machine. Also can use 'ls' ",
	Run: func(cmd *cobra.Command, args []string) {
		db := dao.InitDB()
		var machineDao = dao.NewMachineDao(db)

		machines, err := machineDao.SelectAll()
		if err != nil {
			fmt.Printf("Sorry,select all machine error😞!!!\n")
			return
		} else {
			table.Output(machines)
		}

	},
}

func init() {
	rootCommand.AddCommand(listDataCommand)
}
