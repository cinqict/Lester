package main

import (
	"fmt"
	"io/ioutil"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh"
)

//

func compare(cmd *cobra.Command, args []string) {
	// read in the parameters

	//compFile := args[0]

	// setup a connection to on of the systems

}

//create a ssh client configuration
func createSSHClient(u string, f string, p string) (*ssh.ClientConfig, error) {

	sshConfig := &ssh.ClientConfig{
		User: u,
		Auth: []ssh.AuthMethod{
			ssh.Password(p),
			publicKeyFile(f),
		},
	}

	return sshConfig, nil
}

// parse the public key file
func publicKeyFile(file string) ssh.AuthMethod {
	buffer, err := ioutil.ReadFile(file)
	if err != nil {
		return nil
	}

	key, err := ssh.ParsePrivateKey(buffer)
	if err != nil {
		return nil
	}
	return ssh.PublicKeys(key)
}

//get a new ssh connection
func sshSession(h, p string, s *ssh.ClientConfig) (*ssh.Session, error) {

	connection, err := ssh.Dial("tcp", h+":"+p, s)

	if err != nil {
		return nil, fmt.Errorf("Failed to dial: %s", err)
	}

	session, err := connection.NewSession()
	if err != nil {
		return nil, fmt.Errorf("Failed to create session: %s", err)
	}

	return session, nil
}
