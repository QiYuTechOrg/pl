package logic

import (
    "bytes"
    "context"
    "io"
    "runner/dt"
    "sync"
    "time"
)

import "os/exec"

/// 运行二进制程序
func RunBin(args dt.RunArgs) dt.RunRet {
    ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(args.Timeout))
    defer cancel()

    cmd := exec.CommandContext(ctx, args.Command, args.Args...)

    // 有数据 才需要重定向标准输入
    if len(args.StdinData) > 0 {
        if stdin, err := cmd.StdinPipe(); err != nil {
            return dt.RunRet{Execute: false, Message: "重定向 stdin(标准输入) 失败"}
        } else {
            if _, err = stdin.Write([]byte(args.StdinData)); err != nil {
                return dt.RunRet{Execute: false, Message: "写入 stdin(标准输入) 失败"}
            }
            if stdin.Close() != nil {
                return dt.RunRet{Execute: false, Message: "关闭 stdin(标准输入) 失败"}
            }
        }
    }

    stdout, err := cmd.StdoutPipe()
    if err != nil {
        return dt.RunRet{Execute: false, Message: "重定向 stdout(标准输出) 失败"}
    }
    stderr, err := cmd.StderrPipe()
    if err != nil {
        return dt.RunRet{Execute: false, Message: "重定向 stderr(标准错误) 失败"}
    }

    var wg sync.WaitGroup

    var stdoutString string
    var stdoutError string

    wg.Add(1)
    go func() {
        defer wg.Done()
        maxSize := args.StdoutMaxSize
        stdoutBuf := make([]byte, maxSize)
        buffer := bytes.NewBuffer(stdoutBuf)
        n, err := io.CopyN(buffer, stdout, int64(maxSize))
        stdoutString = string(stdoutBuf[:n])
        if n >= int64(maxSize) {
            stdoutError = "stdout(标准输出) 长度超过了限制"
            _ = cmd.Process.Kill()
        }

        if err == nil || err == io.EOF {
            return
        }
        stdoutError = err.Error()
    }()

    var stderrString string
    var stderrError string

    wg.Add(1)
    go func() {
        defer wg.Done()
        maxSize := args.StderrMaxSize
        stderrBuf := make([]byte, maxSize)
        buffer := bytes.NewBuffer(stderrBuf)
        n, err := io.CopyN(buffer, stderr, int64(maxSize))
        stdoutString = string(stderrBuf[:n])
        if n >= int64(maxSize) {
            stderrError = "stderr(标准错误) 长度超过了限制"
            _ = cmd.Process.Kill()
        }
        if err == nil || err == io.EOF {
            return
        }
        stderrError = err.Error()
    }()

    if err = cmd.Start(); err != nil {
        return dt.RunRet{
            Execute: false,
            Message: `启动子进程失败: ` + err.Error(),
        }
    }

    if err := cmd.Wait(); err != nil {
        exitError, ok := err.(*exec.ExitError)
        if ok {
            return dt.RunRet{
                ExitCode:   exitError.ExitCode(),
                StdoutData: stdoutString,
                StderrData: stderrString,

                StdoutError: stdoutError,
                StderrError: stderrError,
                Execute:     true,
                Message:     err.Error(),
            }
        }

        return dt.RunRet{
            ExitCode:   exitError.ExitCode(),
            StdoutData: stdoutString,
            StderrData: stderrString,

            StdoutError: stdoutError,
            StderrError: stderrError,
            Execute:     true,
            Message:     err.Error(),
        }
    }

    wg.Wait()

    return dt.RunRet{
        ExitCode:   0,
        StdoutData: stdoutString,
        StderrData: stderrString,

        StdoutError: stdoutError,
        StderrError: stderrError,
        Execute:     true,
        Message:     "",
    }
}
