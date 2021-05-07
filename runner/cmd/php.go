package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "runner/dt"
    "runner/file_helper"
    "runner/logic"
    "runner/run_helper"
    "runner/values"
)

func init() {
    rootCmd.AddCommand(phpLang)
}

var phpLang = &cobra.Command{
    Use:   "php",
    Short: "运行 PHP 脚本 [文件路径:" + values.PHPInFile + "]",
    Long: fmt.Sprintf(`在 docker 中 执行 PHP 脚本

运行配置文件路径: %s [此文件必须存在]
执行的脚本文件路径: %s [此文件必须存在]

输出结果文件路径: %s [存在会被重写]
`, values.RunFile, values.PHPInFile, values.OutFile),
    Run: func(cmd *cobra.Command, args []string) {

        c := run_helper.LoadRunConfig()

        runArgs := dt.RunArgs{
            Command:       values.PHPBinary,
            Args:          append([]string{values.PHPInFile}, c.ArgsSplit()...),
            Timeout:       c.Timeout,
            StdinData:     c.StdinData,
            StdoutMaxSize: c.StdoutMaxSize,
            StderrMaxSize: c.StderrMaxSize,
        }

        out := logic.RunBin(runArgs)

        file_helper.WriteJsonToFile(values.OutFile, out)
    },
}
