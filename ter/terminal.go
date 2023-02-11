package ter

import (
	"fmt"
	"github.com/jiuzi/sshm/model"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
	"io"
	"os"
	"runtime"
	"time"
)

type term struct {
	Session   *ssh.Session
	exitMsg   string
	stdout    io.Reader
	stdin     io.Writer
	stderr    io.Reader
	Password  string
	LoginUser string
}

func (t term) shell() error {
	defer func() {
		if t.exitMsg == "" {
			_, _ = fmt.Fprintln(os.Stdout, "[sshm]: the connection was closed on the remote side on ", time.Now().Format(time.RFC822))
		} else {
			_, _ = fmt.Fprintln(os.Stdout, t.exitMsg)
		}
	}()
	fd := int(os.Stdin.Fd())
	if !terminal.IsTerminal(fd) {
		osName := runtime.GOOS
		return fmt.Errorf("%s fd %d is not a terminal,can't create pty of ssh", osName, fd)
	}
	//使用VT100终端来实现tab键提示，上下键查看历史命令，clear键清屏等操作
	//VT100 start
	//windows下不支持VT100
	//使用VT100终端来实现tab键提示，上下键查看历史命令，clear键清屏等操作
	//VT100 start
	//windows下不支持VT100
	state, err := terminal.MakeRaw(fd)
	if err != nil {
		return err
	}
	defer func(fd int, oldState *terminal.State) {
		_ = terminal.Restore(fd, oldState)
	}(fd, state)

	//打开伪终端
	//https://tools.ietf.org/html/rfc4254#page-11
	termWidth, termHeight, err := terminal.GetSize(fd)
	if err != nil {
		return err
	}
	termType := os.Getenv("TERM")
	if termType == "" {
		termType = "xterm-256color"
	}

	err = t.Session.RequestPty(termType, termHeight, termWidth, ssh.TerminalModes{})
	if err != nil {
		return err
	}

	t.stdin, err = t.Session.StdinPipe()
	if err != nil {
		return err
	}
	t.stdout, err = t.Session.StdoutPipe()
	if err != nil {
		return err
	}
	t.stderr, err = t.Session.StderrPipe()

	go func() {
		_, _ = io.Copy(os.Stderr, t.stderr)
	}()
	go func() {
		_, _ = io.Copy(os.Stdout, t.stdout)
	}()
	go func() {
		_, _ = io.Copy(t.stdin, os.Stdin)
	}()
	//启动一个远程shell
	//https://tools.ietf.org/html/rfc4254#page-13
	err = t.Session.Shell()
	if err != nil {
		return err
	}
	//等待远程命令结束或远程shell退出
	return t.Session.Wait()
}

func RunTerminal(machine *model.Machine) error {
	client, err := NewClient(machine)
	if err != nil {
		return err
	}
	defer func(client *ssh.Client) {
		_ = client.Close()
	}(client)
	session, err := client.NewSession()
	if err != nil {
		return err
	}
	defer func(session *ssh.Session) {
		_ = session.Close()
	}(session)

	s := term{
		Session:   session,
		Password:  machine.Password,
		LoginUser: machine.User,
	}
	return s.shell()
}
