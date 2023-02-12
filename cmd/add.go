package cmd

import (
	"fmt"
	"github.com/jiuzi/sshm/dao"
	"github.com/jiuzi/sshm/model"
	"github.com/manifoldco/promptui"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
)

var addDataCommand = &cobra.Command{
	Use:   "add",
	Short: "add one ssh machine",
	Run: func(cmd *cobra.Command, args []string) {

		// sshm add name user@ip:port
		// Please input password:
		// sshm add "TestName" admin@12.12.12.12:22
		var newMachine = &model.Machine{}

		if len(args) == 2 {
			newMachine.Name = args[0]
			userAndHost := strings.Split(args[1], "@")
			newMachine.User = userAndHost[0]

			if strings.Contains(userAndHost[1], ":") {
				hostAndPort := strings.Split(userAndHost[1], ":")
				newMachine.Ip = hostAndPort[0]
				newMachine.Host = hostAndPort[0]
				port, _ := strconv.ParseInt(hostAndPort[1], 10, 32)
				newMachine.Port = int(port)
			} else {
				newMachine.Ip = userAndHost[1]
				newMachine.Host = userAndHost[1]
				newMachine.Port = 22
			}
			newMachine.Password = formInfo("Password:", func(input string) error {
				if len(input) == 0 {
					return errors.New("You must input Password!")
				}
				return nil
			})
			newMachine.Type = "password"

		} else {
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
			password := formInfo("Password:", func(input string) error {
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
				Password: password,
				Type:     "password",
			}
			newMachine = machine
		}

		db := dao.InitDB()
		var machineDao = dao.NewMachineDao(db)
		err := machineDao.Add(newMachine)
		if err != nil {
			fmt.Println("Sorry,add one machine error")
		} else {
			fmt.Println("Congratulations,add one machine success，id is ", newMachine.ID)
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
