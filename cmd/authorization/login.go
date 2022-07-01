/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package authorization

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	loginClientId     string
	loginClientSecret string
	cfgFile           string
)

func login(loginClientId, loginClientSecret string) string {
	return "login with id: " + loginClientId + ", secret: " + loginClientSecret
}

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "this authorization login",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if loginClientId == "" {
			loginClientId = viper.GetViper().GetString("client_id")
		}
		if loginClientSecret == "" {
			loginClientSecret = viper.GetViper().GetString("client_secret")
		}
		fmt.Println(login(loginClientId, loginClientSecret))

	},
}

func init() {

	loginCmd.Flags().StringVarP(&loginClientId, "client_id", "i", "", "the client id")
	loginCmd.Flags().StringVarP(&loginClientSecret, "client_secret", "s", "", "the client secret")

	if err := loginCmd.MarkFlagRequired("client_id"); err != nil {
		fmt.Println(err)
	}

	if err := loginCmd.MarkFlagRequired("client_secret"); err != nil {
		fmt.Println(err)
	}
	// Here you will define your flags and configuration settings.
	AuthorizationCmd.AddCommand(loginCmd)
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}