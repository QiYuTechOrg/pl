package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "os"
)

var rootCmd = &cobra.Command{
    Short: "runner 在 docker 中运行多种语言",
    Long:  `runner 是一个在 docker 中运行多种语言代码的包装工具`,
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        _, _ = fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}
