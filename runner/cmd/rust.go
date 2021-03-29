package cmd

import (
    "github.com/spf13/cobra"
)

func init() {
    rootCmd.AddCommand(rustLang)
}

var rustLang = &cobra.Command{
    Use:   "rust",
    Short: "rust 语言编译执行",
    Long:  `在 docker 中 编译 & 执行 Rust 语言`,
}
