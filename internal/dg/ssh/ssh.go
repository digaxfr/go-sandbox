package ssh

import (
	"fmt"
	"golang.org/x/crypto/ssh"
)

func Func1() string {
	return "Hello world from ssh"
}

func NewSshSession() {
	var hostKey ssh.PublicKey

	fmt.Printf("%v\n", hostKey)
}
