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
    path = "";
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Clone a git repo into a managed directory",
	Long: "Clones a git repo into a specified directory as " +
    "well as perform some basic setup like registering the " +
    "new path with zoxide. This utility also supports a " +
    "default url prefix to make cloning personal repos more " + 
    "convenient. Ex: only specifying 'repo-manager' and auto " +
    "prefixing 'git@github:Ajlow2000'",
	Run: func(cmd *cobra.Command, args []string) {
        if repo == "" {
            cmd.Help()
        } else {
            if (urlPrefix == "") {
                urlPrefix = viper.GetString("urlPrefix")
            }

            if (path == "") {
                path = viper.GetString(("managedDir"))
            }
		    app.Add(urlPrefix, repo, path)
        }
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

    addCmd.Flags().StringVar(&repo, "repo", "", "The url pointing at a git repository or just the project name (fails if no urlPrefix is provided)")
    addCmd.Flags().StringVar(&urlPrefix, "urlPrefix", "", "The url prefix for the repo name (Ex: git@github:myusername/)")
    addCmd.Flags().StringVar(&path, "path", "", "The path to clone the specified repo into")

}
