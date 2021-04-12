package cmd

import (
    "github.com/spf13/cobra"
)

func init() {
    rootCmd.AddCommand(version)
}

var version = &cobra.Command{
    Use:   "version",
    Short: "显示当前版本",
    Long:  "显示软件的当前版本",
    Run: func(cmd *cobra.Command, args []string) {
        println("v0.1.1")
    },
}
