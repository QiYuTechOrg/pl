package run_helper

import (
    "encoding/json"
    "runner/dt"
    "runner/file_helper"
    "runner/values"
)

/// 加载运行时的配置
func LoadRunConfig() dt.RunConfig {
    var c dt.RunConfig

    data := file_helper.ReadToString(values.RunFile)
    if err := json.Unmarshal([]byte(data), &c); err != nil {
        panic(err.Error())
    }

    return c
}
