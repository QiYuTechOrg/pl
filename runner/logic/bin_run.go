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

// RunBin 运行二进制程序
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

    // io 完成等待
    var ioDone sync.WaitGroup

    // 修正 file already closed problem
    // 必须准备好 stdout stderr IO 之后
    // 然后再启动进程
    //
    // 需要 ioPrepare 的原因为:
    // 在某些系统上，因为调度的原因
    // 如果 子进程已经执行完成 而本进程却没有 第一次读取操作 的时候
    // 这时候 进行读取会返回 文件已经关闭(file already closed)的错误
    var ioPrepare sync.WaitGroup

    stdoutData := new(string)
    stdoutError := new(string)

    ioDone.Add(1)
    ioPrepare.Add(1)
    go func() {
        defer ioDone.Done()
        maxSize := int64(args.StdoutMaxSize)
        buffer := bytes.NewBuffer(make([]byte, 0))
        ioPrepare.Done()
        n, err := io.CopyN(buffer, stdout, maxSize)
        *stdoutData = buffer.String()
        if n >= maxSize {
            *stdoutError = "stdout(标准输出) 长度超过了限制"
            _ = cmd.Process.Kill()
        }

        if err == nil || err == io.EOF {
            return
        }
        *stdoutError = err.Error()
    }()

    stderrData := new(string)
    stderrError := new(string)

    ioDone.Add(1)
    ioPrepare.Add(1)
    go func() {
        defer ioDone.Done()
        maxSize := int64(args.StderrMaxSize)
        buffer := bytes.NewBuffer(make([]byte, 0))
        ioPrepare.Done()
        n, err := io.CopyN(buffer, stderr, 0)
        *stderrData = buffer.String()
        if n >= maxSize {
            *stderrError = "stderr(标准错误) 长度超过了限制"
            _ = cmd.Process.Kill()
        }
        if err == nil || err == io.EOF {
            return
        }
        *stderrError = err.Error()
    }()

    ioPrepare.Wait()
    if err = cmd.Start(); err != nil {
        return dt.RunRet{
            Execute: false,
            Message: `启动子进程失败: ` + err.Error(),
        }
    }

    ioDone.Wait()

    if err := cmd.Wait(); err != nil {

        exitError, ok := err.(*exec.ExitError)
        if ok {
            return dt.RunRet{
                ExitCode:   exitError.ExitCode(),
                StdoutData: *stdoutData,
                StderrData: *stderrData,

                StdoutError: *stdoutError,
                StderrError: *stderrError,
                Execute:     true,
                Message:     err.Error(),
            }
        }

        return dt.RunRet{
            ExitCode:   exitError.ExitCode(),
            StdoutData: *stdoutData,
            StderrData: *stderrData,

            StdoutError: *stdoutError,
            StderrError: *stderrError,
            Execute:     true,
            Message:     err.Error(),
        }
    }

    return dt.RunRet{
        ExitCode:   0,
        StdoutData: *stdoutData,
        StderrData: *stderrData,

        StdoutError: *stdoutError,
        StderrError: *stderrError,
        Execute:     true,
        Message:     "",
    }
}
