package main

import (
	"io/ioutil"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh"
)

//
var compareCommand = &cobra.Command{
	Use:   "compare",
	Short: "compares files from two remote systems",
	Long:  "", //TODO document the long output,
	Run:   compare,
}

var privateKeyFile string
var hosts string

func getCompareCommand() *cobra.Command {
	// this flag controls the output destination for our run
	compareCommand.Flags().StringVarP(&privateKeyFile, "privateKey", "P", "~/.ssh/id_rsa", "private keyfile to use to connect to remote machines (defaults to ~/.ssh/id_rsa)")
	compareCommand.Flags().StringVarP(&hosts, "host_list", "h", "localhost", "comma seperated list of hosts (2max for now)")

	//collect the commands in the package
	return compareCommand
}

func compare(cmd *cobra.Command, args []string) {
	// read in the parameters

	compFile := args[0]

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
