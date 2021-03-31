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
    rootCmd.AddCommand(pythonLang)
}

var pythonLang = &cobra.Command{
    Use:   "python",
    Short: "运行 Python 脚本 [文件路径:" + values.PythonInFile + "]",
    Long: fmt.Sprintf(`在 docker 中 执行 Python 脚本

运行配置文件路径: %s [此文件必须存在]
执行的脚本文件路径: %s [此文件必须存在]

输出结果文件路径: %s [存在会被重写]
`, values.RunFile, values.PythonInFile, values.OutFile),
    Run: func(cmd *cobra.Command, args []string) {

        c := run_helper.LoadRunConfig()

        runArgs := dt.RunArgs{
            Command:       values.PythonBinary,
            Args:          append([]string{values.PythonInFile}, c.ArgsSplit()...),
            Timeout:       c.Timeout,
            StdinData:     c.StdinData,
            StdoutMaxSize: c.StdoutMaxSize,
            StderrMaxSize: c.StderrMaxSize,
        }

        out := logic.RunBin(runArgs)

        file_helper.WriteJsonToFile(values.OutFile, out)
    },
}
