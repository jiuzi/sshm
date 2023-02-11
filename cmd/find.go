package cmd

import (
	"fmt"
	"github.com/jiuzi/sshm/dao"
	"github.com/jiuzi/sshm/model"
	"github.com/modood/table"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
)

var isById bool

var findDataCommand = &cobra.Command{
	Use:     "find",
	Short:   "Query machine by condition. Also can use `fd`",
	Aliases: []string{"fd"},
	Example: `
根据id查询
	sshm find --id 2
	sshm find -i 2
	sshm fd -i 2
根据指定的信息模糊查询[Name,Host]
	sshm fd 测试
`,
	Run: func(cmd *cobra.Command, args []string) {
		db := dao.InitDB()
		var machineDao = dao.NewMachineDao(db)

		if len(args) > 0 {
			arg := strings.TrimSpace(args[0])
			var machine *model.Machine
			if isById {
				id, err := strconv.ParseInt(arg, 10, 32)
				if err == nil {
					machine, err = machineDao.SelectById(int(id))
					table.Output([]*model.Machine{machine})
				} else {
					fmt.Printf("Sorry,find machine by %s is error😞!!!\n", arg)
					return
				}
			} else {
				machines, err := machineDao.SelectLikeName(arg)
				if err != nil {
					fmt.Printf("Sorry,find machine like %s is error😞!!!\n", arg)
					return
				} else {
					table.Output(machines)
				}
			}
		} else {
			machines, err := machineDao.SelectAll()
			if err != nil {
				fmt.Printf("Sorry,select all machine error😞!!!\n")
				return
			} else {
				table.Output(machines)
			}

		}
	},
}

func init() {
	findDataCommand.Flags().BoolVarP(&isById, "id", "i", false, "flag use id to query machine")
	rootCommand.AddCommand(findDataCommand)
}
