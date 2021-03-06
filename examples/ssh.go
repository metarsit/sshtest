package examples

import (
	"net"

	"golang.org/x/crypto/ssh"
)

// RunCmdOverSSH establishes connection to remote machine and run command
func RunCmdOverSSH(addr, username, password, cmd string) ([]byte, error) {
	cfg := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.HostKeyCallback(
			func(hostname string, remote net.Addr, key ssh.PublicKey) error {
				return nil
			},
		),
	}

	conn, err := ssh.Dial("tcp", addr, cfg)
	if err != nil {
		return nil, err
	}

	session, err := conn.NewSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()

	return session.CombinedOutput(cmd)
}
