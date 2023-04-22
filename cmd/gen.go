/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var template map[string]map[string]string = map[string]map[string]string{
	"default": {
		"url": "",
	},
	"react":    {},
	"react-ts": {},
}

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "gen - Generates the web app",
	Long: `the 'gen' commands handle the generation according to what template
	you specified which is usually one of these 'default', 'react', 'react-ts'`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println("False input, use it like this 'webgen gen TEMPLATE APP_NAME")
			fmt.Println("TEMPLATE can be one of these(more coming soon):")
			fmt.Println("	x default - normal legacy web app (HTML/CSS/JS)")
			fmt.Println("	x react - a JavaScript react web app")
			fmt.Println("	x react-ts - a TypeScript react web app")
			return
		}

		template := args[0]
		app_name := args[1]

	},
}

func init() {
	rootCmd.AddCommand(genCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// genCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// genCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
