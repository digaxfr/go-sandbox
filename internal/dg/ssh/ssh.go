package ssh

import (
	"bytes"
	"fmt"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
	"os"
)

func Func1() string {
	return "Hello world from ssh"
}

// https://pkg.go.dev/golang.org/x/crypto/ssh#example-Dial
func NewSshSession(sshUser string, sshHost string, sshPort int, sshKeyFile string, sshKeyPassword string, ignoreHostKey bool) {
	// Read in SSH key.
	sshKey, err := os.ReadFile(sshKeyFile)
	if err != nil {
		fmt.Println("Error opening ssh key.")
		fmt.Println(err)
		os.Exit(1)
	}

	// Parse file
	var signer ssh.Signer
	if sshKeyPassword == "" {
		signer, err = ssh.ParsePrivateKey(sshKey)
		if err != nil {
			fmt.Println("Error parsing ssh key.")
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		signer, err = ssh.ParsePrivateKeyWithPassphrase(sshKey, []byte(sshKeyPassword))
		if err != nil {
			fmt.Println("Error parsing ssh key.")
			fmt.Println(err)
			os.Exit(1)
		}
	}

	var hostKeyCallback ssh.HostKeyCallback
	if ignoreHostKey {
		hostKeyCallback = ssh.InsecureIgnoreHostKey()
	} else {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Error getting home directory.")
			fmt.Println(err)
			os.Exit(1)
		}

		knownHostsFile := homeDir + "/.ssh/known_hosts"
		hostKeyCallback, err = knownhosts.New(knownHostsFile)
		if err != nil {
			fmt.Println("Error generating knownhosts.")
			fmt.Println(err)
			os.Exit(1)
		}
	}

	config := &ssh.ClientConfig{
		User: sshUser,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: hostKeyCallback,
	}

	client, err := ssh.Dial("tcp", fmt.Sprintf("%v:%v", sshHost, sshPort), config)
	if err != nil {
		fmt.Println("Error connecting to host.")
		fmt.Println(err)
		os.Exit(1)
	}
	defer client.Close()

	// Multiple sessions over a clientconn.
	session, err := client.NewSession()
	if err != nil {
		fmt.Println("Error creating new session.")
		fmt.Println(err)
		os.Exit(1)
	}
	defer session.Close()

	// Once a Session is created, you can execute a single command on
	// the remote side using the Run method.
	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run("/usr/bin/whoami"); err != nil {
		fmt.Println("Error running command over ssh.")
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(b.String())
}
