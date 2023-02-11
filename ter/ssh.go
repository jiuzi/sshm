package ter

import (
	"fmt"
	"github.com/jiuzi/sshm/model"
	"github.com/mitchellh/go-homedir"
	"golang.org/x/crypto/ssh"
	"log"
	"os"
	"time"
)

func NewClient(machine *model.Machine) (*ssh.Client, error) {

	config := &ssh.ClientConfig{
		Timeout:         time.Second * 5,
		User:            machine.User,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), //这个可以， 但是不够安全
	}

	if machine.Type == "password" {
		config.Auth = []ssh.AuthMethod{ssh.Password(machine.Password)}
	} else {
		config.Auth = []ssh.AuthMethod{publicKeyAuthFunc(machine.Key)}
	}
	addr := fmt.Sprintf("%s:%d", machine.Host, machine.Port)
	client, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		return nil, err
	}
	return client, nil

}

func publicKeyAuthFunc(kPath string) ssh.AuthMethod {
	keyPath, err := homedir.Expand(kPath)
	if err != nil {
		log.Fatal("find key's home dir failed", err)
	}
	key, err := os.ReadFile(keyPath)
	if err != nil {
		log.Fatal("ssh key file read failed", err)
	}
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatal("ssh key signer failed", err)
	}
	return ssh.PublicKeys(signer)
}
