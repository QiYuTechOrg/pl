package dt

import "github.com/google/shlex"

/// 运行配置
type RunConfig struct {
    Args          string `json:"args"`
    Timeout       int    `json:"timeout"`
    StdinData     string `json:"stdin_data"`
    StdoutMaxSize int    `json:"stdout_max_size"`
    StderrMaxSize int    `json:"stderr_max_size"`
}

func (c RunConfig) ArgsSplit() []string {
    if args, err := shlex.Split(c.Args); err != nil {
        panic(err.Error())
    } else {
        return args
    }
}

/// 运行参数
type RunArgs struct {
    Command       string   `json:"command"`         /// 需要运行的文件 使用绝对路径
    Args          []string `json:"args"`            /// 二进制的参数 [可以为空 表示没有]
    Timeout       int      `json:"timeout"`         /// 最大允许的执行时间
    StdinData     string   `json:"stdin_data"`      /// 标准输入的文件内容 todo 使用文件 如果允许输入的内容很多
    StdoutMaxSize int      `json:"stdout_max_size"` /// 标准输出 最大的大小
    StderrMaxSize int      `json:"stderr_max_size"` /// 标准错误 最大的大小
}

/// 运行结果
type RunRet struct {
    ExitCode   int    `json:"exit_code"`   // 程序退出的码
    StdoutData string `json:"stdout_data"` // 标准输出的数据
    StderrData string `json:"stderr_data"` // 标准错误的数据

    StdoutError string `json:"stdout_error"` // 标准输出(stdout)的错误
    StderrError string `json:"stderr_error"` // 标准错误(stderr)的错误
    Execute     bool   `json:"execute"`      // 是否执行二进制程序
    Message     string `json:"message"`      // 提示信息 解析 Stdout 失败 等
}
