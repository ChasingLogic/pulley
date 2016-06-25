package pulley_test

import (
	"os"
	"os/user"
	"testing"

	"github.com/chasinglogic/pulley"
)

func getTestClient() *pulley.Client {
	var c *pulley.Client

	ptu := os.Getenv("PULLEY_TEST_USER")
	if ptu != "" {
		c = pulley.New(ptu)
	} else {
		u, _ := user.Current()
		c = pulley.New(u.Username)
	}

	pts := os.Getenv("PULLEY_TEST_SERVER")
	if pts != "" {
		c.HostName = pts
	}

	return c
}

func TestClientDefaults(t *testing.T) {
	currentUser, _ := user.Current()
	c := pulley.New(currentUser.Username)

	if c.User != currentUser.Username {
		t.Errorf("Expected: %s, Got: %s", currentUser.Username, c.User)
	}

	if c.HostName != "localhost" {
		t.Errorf("Expected: %s, Got: %s", "localhost", c.HostName)
	}

	if c.Port != "22" {
		t.Errorf("Expected: %s, Got: %s", "22", c.Port)
	}
}

func TestExec(t *testing.T) {
	c := getTestClient()
	err := c.Connect()
	if err != nil {
		t.Error("Failed to connect.", err)
	}

	result := c.Exec(`echo "Hello World"`)
	if result.Failure() {
		t.Error("Failed to exec echo.", result, result.Err())
	}

	if result.Success() == false {
		t.Error("Success returned false when it shouldn't have.", result)
	}

	if result.Err() != nil {
		t.Error("Err returned non-nil but success and failure both passed.",
			result.Err())
	}

	t.Log(result.String())
}

func TestExecAsync(t *testing.T) {
	c := getTestClient()
	err := c.Connect()
	if err != nil {
		t.Error("Failed to connect.", err)
	}

	resChannel := make(chan pulley.Result)

	c.ExecAsync(`echo "Hello World"`, resChannel)

	result := <-resChannel
	if result.Err() != nil {
		t.Error("Result was an error.", result, result.Err())
	}

	t.Log(result.String())
}
