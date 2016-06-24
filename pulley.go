package pulley

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"golang.org/x/crypto/ssh"
)

// Client is the client for the current ssh session.
type Client struct {
	config     *ssh.ClientConfig
	connection *ssh.Client
}

// New creates a default SSHClient you can use globally, using an import such as
// ssh import "github.com/chasinglogic/pulley"
func New() *Client {
	return &Client{}
}

// Session will return a new session for the current connection or an error if
// there was one. You only need to use this if you're doing something advanced.
func (s *Client) Session() (*ssh.Session, error) {
	return s.connection.NewSession()
}

// Connect will connect the client to the given hostname
func (s *Client) Connect(hostname, port string) error {
	var err error

	s.connection, err = ssh.Dial("tcp",
		fmt.Sprintf("%s:%s", hostname, port),
		s.config)

	return err
}

// LoadDefaultKey will load the ssh key at $HOME/.ssh/id_rsa
func (s *Client) LoadDefaultKey() error {
	keyFile := filepath.Join(os.Getenv("HOME"), ".ssh", "id_rsa")

	key, err := ioutil.ReadFile(keyFile)
	if err != nil {
		return err
	}

	return s.LoadKey(key)
}

// LoadKey will load the key given in a []byte
func (s *Client) LoadKey(key []byte) error {
	parsedKey, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return err
	}

	s.config.Auth = append(s.config.Auth, ssh.PublicKeys(parsedKey))
	return nil
}

// TODO: Not right
type Result string

// Exec runs the command on the server that's connected to by this client, if
// It will handle sessions automatically.
func (s *Client) Exec(cmd string) Result {
	return Result{}
}
