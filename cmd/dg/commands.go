package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/digaxfr/go-sandbox/internal/dg/ssh"
)

const envVarPrivateKeyPassword = "PRIVATE_KEY_PASSWORD"

type SshCmd struct {
	Hostname               string `help:"Host/IP of remote host." required:""`
	IgnoreHostKey          bool   `help:"Ignore checking host SSH key/fingerprint." default:"false"`
	Port                   int    `help:"Port of remote host." default:"22"`
	PrivateKey             string `help:"Path to private key." required:""`
	PrivateKeyPasswordFile string `help:"Path to file containing passphrase of the private key. Can export PRIVATE_KEY_PASSWORD alternatively."`
	User                   string `help:"User of remote host" required:""`

	Ping SshPingCmd `help:"Ping the remote host." cmd:""`
}

type SshPingCmd struct{}

func (cmd *SshCmd) Run() error {
	// Figure out SSH key password from CLI, envvars, or empty.
	var sshKeyPassword string

	// Check env vars as it will take precedence.
	sshKeyPasswordEnvVar := os.Getenv(envVarPrivateKeyPassword)
	if sshKeyPasswordEnvVar == "" {
		// No env var. Use flag.
		sshKeyPasswordFile, err := os.ReadFile(cmd.PrivateKeyPasswordFile)
		if err != nil {
			fmt.Println("Error opening ssh key password file.")
			fmt.Println(err)
			os.Exit(1)
		}
		sshKeyPassword = string(sshKeyPasswordFile)
	} else {
		sshKeyPassword = sshKeyPasswordEnvVar
	}

	// Fix things up
	sshKeyPassword = strings.TrimSuffix(sshKeyPassword, "\n")

	ssh.NewSshSession(cmd.User, cmd.Hostname, cmd.Port, cmd.PrivateKey, sshKeyPassword, cmd.IgnoreHostKey)
	return nil
}
