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

var verbose bool
var logging bool
var logFile string
var verboseLog bool
var quiet bool

func main() {
	lester.Execute()
}

func init() {
	lester.AddCommand()

	lester.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	lester.PersistentFlags().BoolVar(&quiet, "quiet", false, "run in quiet mode")
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
