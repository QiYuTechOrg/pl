package run_helper

import (
    "encoding/json"
    "github.com/google/shlex"
    "runner/dt"
    "runner/file_helper"
    "runner/values"
)

/// 读取参数
func ReadArgs() []string {
    content := file_helper.ReadToString(values.ArgsFile)
    if cmdArgs, err := shlex.Split(content); err != nil {
        panic(err.Error)
    } else {
        return cmdArgs
    }
}

/// 读取标准输入
func ReadStdin() string {
    return file_helper.ReadToString(values.StdinFile)
}

/// 加载运行时的配置
func LoadRunConfig() dt.RunConfig {
    var c dt.RunConfig

    data := file_helper.ReadToString(values.RunFile)
    if err := json.Unmarshal([]byte(data), &c); err != nil {
        panic(err.Error())
    }

    return c
}
