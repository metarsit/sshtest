package examples

import (
	"testing"

	"github.com/metarsit/sshtest"
	"gotest.tools/v3/assert"
)

func TestRunCmdOverSSH(t *testing.T) {
	const (
		fakeUsername = "fake_username"
		fakePassword = "super_secret_password"
		fakeCmd      = "fake_command"
		addr         = "localhost:2222"

		returnMsg = "something was run over there"
	)

	hp := sshtest.NewHoneyPot(addr)

	// Start the server in the background
	go func() {
		hp.ListenAndServe()
	}()
	t.Cleanup(func() {
		hp.Close()
	})

	hp.SetReturnString(returnMsg)

	result, err := RunCmdOverSSH(addr, fakeUsername, fakePassword, fakeCmd)
	assert.NilError(t, err)
	assert.Equal(t, string(result), returnMsg)
}
