package connection

import (
	"bytes"
	"fmt"

	"golang.org/x/crypto/ssh"
)

type SSHClient struct {
	client *ssh.Client
}

func NewSSHClient(host, user string, port int, privateKey []byte) (*SSHClient, error) {
	signer, err := ssh.ParsePrivateKey(privateKey)
	if err != nil {
		return nil, err
	}

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // For simplicity; use proper host key verification in production
	}

	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", host, port), config)
	if err != nil {
		return nil, err
	}

	return &SSHClient{client: client}, nil
}

func (c *SSHClient) RunCommand(cmd string, sudo bool) (string, error) {
	session, err := c.client.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()

	if sudo {
		cmd = fmt.Sprintf("sudo %s", cmd)
	}

	var stdout, stderr bytes.Buffer
	session.Stdout = &stdout
	session.Stderr = &stderr

	if err := session.Run(cmd); err != nil {
		return stderr.String(), err
	}

	return stdout.String(), nil
}

func (c *SSHClient) Close() {
	c.client.Close()
}
