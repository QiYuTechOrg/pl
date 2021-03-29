package cmd

import (
    "encoding/json"
    "github.com/spf13/cobra"
    "io/ioutil"
    "runner/dt"
    "runner/logic"
)

var jsonFile string

func init() {
    jsonCmd.Flags().StringVar(&jsonFile, "file", "", "指定要运行的 JSON 配置文件")
    rootCmd.AddCommand(jsonCmd)
}

/// JSON 执行
var jsonCmd = &cobra.Command{
    Use:   "json",
    Short: "使用 json 配置测试启动命令行",
    Long:  `加载 json 配置文件, 然后进行命令行的测试`,
    Run: func(cmd *cobra.Command, args []string) {
        if jsonFile == "" {
            return
        }

        data, err := ioutil.ReadFile(jsonFile)
        if err != nil {
            panic(err)
        }
        runArgs := dt.RunArgs{}
        if err := json.Unmarshal(data, &runArgs); err != nil {
            panic(err)
        }
        runRet := logic.RunBin(runArgs)
        if out, err := json.Marshal(runRet); err != nil {
            panic(err)
        } else {
            println(string(out))
        }
    },
}
