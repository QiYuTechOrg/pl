package cmd

import (
    "github.com/spf13/cobra"
    "runner/dt"
    "runner/file_helper"
    "runner/logic"
    "runner/run_helper"
)

func init() {
    rootCmd.AddCommand(c)
}

/// 编译 & 执行 C 语言的
var c = &cobra.Command{
    Use:   "c",
    Short: "运行 /data/a.out 程序",
    Long:  `在 docker 中 编译 & 执行 C 语言`,
    Run: func(cmd *cobra.Command, args []string) {

        c := run_helper.LoadRunConfig()

        runArgs := dt.RunArgs{
            Command:       "/data/a.out", // 这是一个固定值
            Args:          c.ArgsSplit(),
            Timeout:       c.Timeout,
            StdinData:     c.StdinData,
            StdoutMaxSize: c.StdoutMaxSize,
            StderrMaxSize: c.StderrMaxSize,
        }

        out := logic.RunBin(runArgs)

        file_helper.WriteJsonToFile("out.json", out)
    },
}
