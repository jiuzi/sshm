package cmd

import (
	"fmt"
	"github.com/jiuzi/sshm/dao"
	"github.com/jiuzi/sshm/model"
	"github.com/manifoldco/promptui"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"strconv"
)

var addDataCommand = &cobra.Command{
	Use:   "add",
	Short: "add one ssh machine",
	Run: func(cmd *cobra.Command, args []string) {

		// sshm add admin@12.12.12.12 -p 1212

		// sshm ssh --type=passwd

		fmt.Println("Begin add one machine...")
		name := formInfo("name:", func(input string) error {
			if len(input) == 0 {
				return errors.New("You must input name!")
			}
			return nil
		})
		host := formInfo("Host you can input ip:", func(input string) error {
			if len(input) == 0 {
				return errors.New("You must input Host!")
			}
			return nil
		})
		ip := formInfo("IP:", func(input string) error {
			if len(input) == 0 {
				return errors.New("You must input IP!")
			}
			return nil
		})
		portStr := formInfo("Port:", func(input string) error {
			if len(input) == 0 {
				return errors.New("You must input port!")
			}
			_, err := strconv.ParseInt(input, 10, 32)
			if err != nil {
				return errors.New("port is a number!")
			}
			return nil
		})

		user := formInfo("User:", func(input string) error {
			if len(input) == 0 {
				return errors.New("You must input User!")
			}
			return nil
		})
		Password := formInfo("Password:", func(input string) error {
			if len(input) == 0 {
				return errors.New("You must input Password!")
			}
			return nil
		})
		portNum, _ := strconv.ParseInt(portStr, 10, 32)
		machine := &model.Machine{
			Name:     name,
			Host:     host,
			Ip:       ip,
			Port:     int(portNum),
			User:     user,
			Password: Password,
			Type:     "password",
		}
		db := dao.InitDB()
		var machineDao = dao.NewMachineDao(db)
		err := machineDao.Add(machine)
		if err != nil {
			fmt.Println("Sorry,add one machine error")
		} else {
			fmt.Println("Congratulations,add one machine success，id is ", machine.ID)
		}

	},
}

func formInfo(label string, validate func(input string) error) string {
	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}
	prompt := promptui.Prompt{
		Label:     label,
		Templates: templates,
		Validate:  validate,
	}
	result, _ := prompt.Run()
	return result
}

func init() {
	rootCommand.AddCommand(addDataCommand)
}
