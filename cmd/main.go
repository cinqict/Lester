package main

import

//external libraries

(
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var lester = &cobra.Command{
	Use:   "lester",
	Short: "lester provides a simple interface to compare files on two different machines over ssh",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		initializeConfig()
	},
}

var compareCommand = &cobra.Command{
	Use:   "compare",
	Short: "compares files from two remote systems",
	Long:  "", //TODO document the long output,
	Run:   compare,
}

var verbose bool
var quiet bool

var userName string
var privateKeyFile string
var hosts string

func main() {
	lester.Execute()
}

func init() {
	// global flags for lester
	lester.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	lester.PersistentFlags().BoolVar(&quiet, "quiet", false, "run in quiet mode")

	// flags controlling the compare command
	compareCommand.Flags().StringVarP(&privateKeyFile, "privateKey", "P", "~/.ssh/id_rsa", "private keyfile to use to connect to remote machines (defaults to ~/.ssh/id_rsa)")
	compareCommand.Flags().StringVarP(&userName, "user", "u", "", "username to use when connecting to remote machines")
	compareCommand.Flags().StringVarP(&hosts, "host_list", "h", "localhost", "comma seperated list of hosts (2max for now)")

	lester.AddCommand(compareCommand)

}

func initializeConfig() {
	// get input from config files

	// configfile name is lester
	viper.SetConfigName("lester")

	// add the filepaths that will be used
	viper.AddConfigPath("/etc/lester/")
	viper.AddConfigPath("$HOME/.lester")
	viper.AddConfigPath(".")

	// Handle errors reading the config file
	err := viper.ReadInConfig()
	if err != nil {
		//debugInfo := viper.Debug()
		fmt.Println("error loading config")
	}
}
