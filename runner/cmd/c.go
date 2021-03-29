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
    rootCmd.AddCommand(c)
}

/// 编译 & 执行 C 语言的
var c = &cobra.Command{
    Use:   "c",
    Short: "运行 " + values.CLangADotOutFile + " 程序",
    Long: fmt.Sprintf(`在 docker 中 执行 a.out  二进制程序

运行配置文件路径: %s [此文件必须存在]
执行的 a.out 文件路径: %s [此文件必须存在]
命令行参数文件路径: %s [此文件必须存在, 可以为空]
标准输入文件路径: %s [此文件必须存在, 可以为空]

输出结果文件路径: %s [此文件必须没有存在]
`, values.CLangADotOutFile, values.ArgsFile, values.StdinFile, values.RunFile, values.OutFile),
    Run: func(cmd *cobra.Command, args []string) {

        c := run_helper.LoadRunConfig()

        runArgs := dt.RunArgs{
            Command:       values.CLangADotOutFile,
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
