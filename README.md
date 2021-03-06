# SSH Test Library

SSH Testing in GO can, at time, be troublesome. This is a helper library for users to test functions that do SSH to another host

## Installation

sshtest is compatible with modern Go releases in module mode, with Go installed:

```bash
go get github.com/metarsit/sshtest
```

## Usage
Below is an example of the usage of the library

```go
import (
	"testing"

	"github.com/metarsit/sshtest"
	"gotest.tools/v3/assert"
)

func TestYourSSHFunction(t *testing.T) {
    // Example of an IP address for localhost
    const (
        addr      = "127.0.0.1:2222"
        returnMsg = "some_return"
    )
    
    hp := ssh.NewHoneyPot(addr)
    
    // Close the Honeypot when goes out of scope
    t.Cleanup(func() {
        hp.Close()
    })
    
    // Run Honeypot in the background
    go func() {
        hp.ListenAndServe()
    }()

    // Set output
    hp.SetReturnString(returnMsg)

    // Call the function that talks to the addr specified
    someReturn, someError := YourSSHFunction(addr)

    // Assert the returns
    assert.NilError(t, someError)
    assert.Equal(t, someReturn, returnMsg)
}
```

### Authentication
No password is required to access honeypot

## Pre-requisite
- docker
- make
- go

## Credit
Special thanks to `gliderlabs` for creating and providing us with [ssh library](https://github.com/gliderlabs/ssh)
