/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/Ajlow2000/repo-manager/app"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
    debug = false
    repo = "";
    urlPrefix = "";
    destination = "";
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add URL_PATH",
    DisableFlagsInUseLine: true,
    Args: cobra.ExactArgs(1),
	Short: "Clone a git repo into a managed directory",
	Long: "Clones a git repo into a specified directory as " +
    "well as perform some basic setup like registering the " +
    "new path with zoxide. This utility also supports a " +
    "default url prefix to make cloning personal repos more " + 
    "convenient. Ex: only specifying 'repo-manager' and auto " +
    "prefixing 'git@github:Ajlow2000'",
	Run: func(cmd *cobra.Command, args []string) {
        if len(args) != 1 {
            cmd.Help()
        } else {
            repo = args[0]
            if (urlPrefix == "") {
                urlPrefix = viper.GetString("urlPrefix")
            }

            if (destination == "") {
                destination = viper.GetString(("managedDir"))
            }
		    app.Add(urlPrefix, repo, destination)
        }
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

    addCmd.Flags().StringVar(&urlPrefix, "urlPrefix", "", "The url prefix for the repo name (Ex: git@github:myusername/)")
    addCmd.Flags().StringVar(&destination, "destination", "", "Filepath to clone the specified repo into")

}
