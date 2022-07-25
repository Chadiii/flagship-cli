/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/Chadiii/flagship/cmd/authorization"
	"github.com/Chadiii/flagship/cmd/campaign"
	"github.com/Chadiii/flagship/cmd/panic"
	"github.com/Chadiii/flagship/cmd/project"
	"github.com/Chadiii/flagship/cmd/user"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile                string
	Token                  string
	Account_id             string
	Account_environment_id string
	Output_format          string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "flagship",
	Short: "Cli to manage your usecases, project, users etc...",
	Long: `
	The goal of the cli is to give the user the ability to manage his account
	`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Root().Help()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func addSubCommandPalettes() {
	rootCmd.AddCommand(campaign.CampaignCmd)
	rootCmd.AddCommand(project.ProjectCmd)
	rootCmd.AddCommand(authorization.ConfigureCmd)
	rootCmd.AddCommand(authorization.AuthenticateCmd)
	rootCmd.AddCommand(panic.PanicCmd)
	rootCmd.AddCommand(user.UserCmd)
}
func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&Token, "token", "t", "", "authorization token")
	rootCmd.PersistentFlags().StringVarP(&Account_id, "account-id", "a", "", "account id")
	rootCmd.PersistentFlags().StringVarP(&Account_environment_id, "account-environment-id", "e", "", "account env id")
	rootCmd.PersistentFlags().StringVarP(&Output_format, "output-format", "f", "table", "output format")
	viper.BindPFlag("token", rootCmd.PersistentFlags().Lookup("token"))
	viper.BindPFlag("account_id", rootCmd.PersistentFlags().Lookup("account-id"))
	viper.BindPFlag("account_environment_id", rootCmd.PersistentFlags().Lookup("account-environment-id"))
	viper.BindPFlag("output_format", rootCmd.PersistentFlags().Lookup("output-format"))

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.flagship/credentials.yaml)")

	addSubCommandPalettes()
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		homeDir, err := os.UserHomeDir()
		cobra.CheckErr(err)
		// Search config in home directory with name ".flagship" (without extension).
		viper.SetConfigFile(homeDir + "/.flagship/credentials.yaml")
	}

	// read in environment variables that match
	// If a config file is found, read it in.
	viper.MergeInConfig()
}
