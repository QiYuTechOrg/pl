package logic

import (
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

    go func() {
        wg.Add(1)
        defer wg.Done()

        var buf = make([]byte, 2048)
        for {
            n, err := stdout.Read(buf)
            if n == 0 {
                break
            }
            stdoutString += string(buf[:n])
            if len(stdoutString) > args.StdoutMaxSize {
                stdoutError = "stdout(标准输出) 长度超过了限制"
                _ = cmd.Process.Kill()
                break
            }
            if err == io.EOF {
                break
            }
            if err != nil {
                stdoutError = err.Error()
                break
            }
        }
    }()

    var stderrString string
    var stderrError string

    go func() {
        wg.Add(1)
        defer wg.Done()
        var buf = make([]byte, 2048)
        for {
            n, err := stderr.Read(buf)
            if n == 0 {
                break
            }
            s := string(buf[:n])
            stderrString += s
            if len(stderrString) > args.StderrMaxSize {
                stdoutError = "stderr(标准错误) 长度超过了限制"
                _ = cmd.Process.Kill()
                break
            }
            if err == io.EOF {
                break
            }
            if err != nil {
                stderrError = err.Error()
                break
            }
        }
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
