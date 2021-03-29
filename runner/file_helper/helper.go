package file_helper

import (
    "encoding/json"
    "io/ioutil"
)

/// 读取文件为字符串
func ReadToString(filename string) string {
    data, err := ioutil.ReadFile(filename)
    if err != nil {
        panic(err.Error())
    }
    return string(data)
}

/// 把 data 写入到文件中
func WriteJsonToFile(filename string, data interface{}) {
    bytes, err := json.Marshal(data)
    if err != nil {
        panic(err.Error())
    }

    if err := ioutil.WriteFile(filename, bytes, 0666); err != nil {
        panic(err.Error())
    }
}
