package clients

import (
	"bytes"
	"context"
	"errors"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// STACKITCLIFlow invoke the STACKIT CLI from PATH to get the access token.
// If successful, then token is passed to clients.TokenFlow.
type STACKITCLIFlow struct {
	TokenFlow
}

// STACKITCLIFlowConfig is the flow config
type STACKITCLIFlowConfig struct{}

// GetConfig returns the flow configuration
func (c *STACKITCLIFlow) GetConfig() STACKITCLIFlowConfig {
	return STACKITCLIFlowConfig{}
}

func (c *STACKITCLIFlow) Init(_ *STACKITCLIFlowConfig) error {
	token, err := c.getTokenFromCLI()
	if err != nil {
		return err
	}

	c.config = &TokenFlowConfig{
		ServiceAccountToken: strings.TrimSpace(token),
	}

	c.configureHTTPClient()
	return c.validate()
}

func (c *STACKITCLIFlow) getTokenFromCLI() (string, error) {
	return runSTACKITCLICommand(context.TODO(), "stackit auth get-access-token")
}

// runSTACKITCLICommand executes the command line and returns the output.
func runSTACKITCLICommand(ctx context.Context, commandLine string) (string, error) {
	var cliCmd *exec.Cmd
	if runtime.GOOS == "windows" {
		dir := os.Getenv("SYSTEMROOT")
		if dir == "" {
			return "", errors.New("environment variable 'SYSTEMROOT' has no value")
		}
		cliCmd = exec.CommandContext(ctx, "cmd.exe", "/c", commandLine)
		cliCmd.Dir = dir
	} else {
		cliCmd = exec.CommandContext(ctx, "/bin/sh", "-c", commandLine)
		cliCmd.Dir = "/bin"
	}
	cliCmd.Env = os.Environ()
	var stderr bytes.Buffer
	cliCmd.Stderr = &stderr

	output, err := cliCmd.Output()
	if err != nil {
		msg := stderr.String()
		var exErr *exec.ExitError
		if errors.As(err, &exErr) && exErr.ExitCode() == 127 || strings.HasPrefix(msg, "'stackit' is not recognized") {
			msg = "STACKIT CLI not found on path"
		}
		if msg == "" {
			msg = err.Error()
		}
		return "", errors.New(msg)
	}

	return string(output), nil
}
