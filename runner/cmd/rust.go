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
    rootCmd.AddCommand(rustLang)
}

var rustLang = &cobra.Command{
    Use:   "rust",
    Short: "运行 " + values.RustOutFile + " 程序",
    Long: fmt.Sprintf(`在 docker 中 执行 rust  二进制程序

运行配置文件路径: %s [此文件必须存在]
执行的二进制文件路径: %s [此文件必须存在]
命令行参数文件路径: %s [此文件必须存在, 可以为空]
标准输入文件路径: %s [此文件必须存在, 可以为空]

输出结果文件路径: %s [此文件必须没有存在]
`, values.RustOutFile, values.ArgsFile, values.StdinFile, values.RunFile, values.OutFile),
    Run: func(cmd *cobra.Command, args []string) {

        c := run_helper.LoadRunConfig()

        runArgs := dt.RunArgs{
            Command:       values.RustOutFile,
            Args:          c.ArgsSplit(),
            Timeout:       c.Timeout,
            StdinData:     c.StdinData,
            StdoutMaxSize: c.StdoutMaxSize,
            StderrMaxSize: c.StderrMaxSize,
        }

        out := logic.RunBin(runArgs)

        file_helper.WriteJsonToFile(values.OutFile, out)
    },
}
