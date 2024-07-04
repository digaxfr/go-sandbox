package main

import (
	"fmt"
)

type SshCmd struct {
	Hostname string `help:"Host/IP of remote host." required`
	Port int `help:"Port of remote host." default:"22"`
	PrivateKey string `help:"Path to private key." required`
	PrivateKeyPasswordFile string `help:"Path to file containing passphrase of the private key. Can export PRIVATE_KEY_PASSWORD alternatively."`
	User string `help:"User of remote host" required`

	Ping SshPingCmd `help:"Ping the remote host." cmd:""`
}

type SshPingCmd struct {}

func (cmd *SshCmd) Run() error {
	fmt.Println("Inside SshCmd Run()")
	return nil
}
