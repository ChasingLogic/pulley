package pulley_test

import (
	"os/user"
	"testing"

	"github.com/chasinglogic/pulley"
)

func TestClientDefaults(t *testing.T) {
	c := pulley.New()

	currentUser, err := user.Current()
	if err != nil {
		t.Error("Failed to get current user.")
	}

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
	c := pulley.New()

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

	t.Log(result)
}

func TestExecAsync(t *testing.T) {
	c := pulley.New()

	resChannel := make(chan pulley.Result)

	c.ExecAsync(`echo "Hello World"`, resChannel)

	result := <-resChannel
	if result.Err() != nil {
		t.Error("Result was an error.", result, result.Err())
	}

	t.Log(result)
}
