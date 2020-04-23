package cmd

import (
	"fmt"
	"cmd2yaml/pkg"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var cfgFile string
var input string

var cmd = &cobra.Command{
	Use:   "cmd2yaml",
	Short: "docker cmd to k8s yaml for pod",
	Long:  `docker cmd to k8s yaml for pod`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if len(input) == 0 {
			cmd.Help()
			return
		}
		pkg.Cmd2yaml(input)
	},
}

func Execute() {
	cmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	//rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.demo.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	cmd.Flags().StringVarP(&input, "input", "i", "", "input docker command ")
	cmd.Flags().StringVarP(&input, "filepath", "f", "", "the file include input docker command ")
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".demo" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".cmd2yaml")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
