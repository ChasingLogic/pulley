// Package pulley is a wrapper around the golang.org/x/crypto/ssh package
// providing a suckless experience.
package pulley

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"golang.org/x/crypto/ssh"
)

// Client is the client for the ssh connection and the primary way to interact
// with pulley.
type Client struct {
	HostName   string
	Port       string
	User       string
	config     *ssh.ClientConfig
	connection *ssh.Client
}

// New creates a default Client. Modify this client via it's public members,
// HostName, Port, and User. These default to localhost, 22, and the current
// user running the process.
func New(user string) *Client {
	return &Client{
		HostName: "localhost",
		Port:     "22",
		User:     user,
		config: &ssh.ClientConfig{
			User: user,
		},
	}
}

// Session will return a new session for the current connection or an error if
// there was one. You only need to use this if you're doing something advanced.
func (s *Client) Session() (*ssh.Session, error) {
	return s.connection.NewSession()
}

// Connect will connect the client to it's hostname and port, if LoadKey or
// LoadDefaultKey have not been called yet this will call LoadDefaultKey
func (s *Client) Connect() error {
	var err error

	if s.config.Auth == nil {
		kerr := s.LoadDefaultKey()
		if kerr != nil {
			return kerr
		}
	}

	s.connection, err = ssh.Dial("tcp",
		fmt.Sprintf("%s:%s", s.HostName, s.Port),
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

// LoadKey will load the given key
func (s *Client) LoadKey(key []byte) error {
	parsedKey, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return err
	}

	s.config.Auth = append(s.config.Auth, ssh.PublicKeys(parsedKey))
	return nil
}

// Exec runs the command on the server that's connected to by this client, if
// It will handle sessions automatically and return a pulley.Result
func (s *Client) Exec(cmd string) Result {
	sess, serr := s.Session()
	if serr != nil {
		return Result{err: serr}
	}
	defer sess.Close()

	var r Result

	r.Output, r.err = sess.Output(cmd)
	return r
}

// ExecErr is the same as exec however the result's output will have both
// stdout and stderr.
func (s *Client) ExecErr(cmd string) Result {
	sess, serr := s.Session()
	if serr != nil {
		return Result{err: serr}
	}
	defer sess.Close()

	var r Result

	r.Output, r.err = sess.CombinedOutput(cmd)
	return r
}

// ExecAsync is the same as exec however will execute in a go routine and takes
// a channel which it will send the result over.
func (s *Client) ExecAsync(cmd string, rc chan Result) {
	go func() {
		rc <- s.Exec(cmd)
	}()
}

// ExecAsyncErr is the same as ExecAsync however the result's output will have
// both stderr and stdout.
func (s *Client) ExecAsyncErr(cmd string, rc chan Result) {
	go func() {
		rc <- s.ExecErr(cmd)
	}()
}

// Ugly will return the underlying ssh.Client and ssh.ClientConfig in case you
// need those structs directly.
func (s *Client) Ugly() (*ssh.Client, *ssh.ClientConfig) {
	return s.connection, s.config
}
