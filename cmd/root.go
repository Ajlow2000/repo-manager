package cmd

import (
	"os"

	"github.com/Ajlow2000/add-repo/app"
	"github.com/spf13/cobra"
)

var (
    debug = false
    url = "";
    path = "";
)


var rootCmd = &cobra.Command{
	Use:   "add-repo",
	Short: "",
    Long: "",
    Version: "", 
    Run: func(cmd *cobra.Command, args []string) {
        if url == "" {
            cmd.Help()
        } else {
		    app.Main(url, path)
        }
    },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
    rootCmd.Flags().StringVar(&url, "url", "", "The url pointing at a git repository")
    rootCmd.Flags().StringVar(&path, "path", "$HOME/repos", "The path to clone the specified url into")
}
