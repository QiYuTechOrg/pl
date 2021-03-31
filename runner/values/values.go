package values

const (
    /// 全局工作目录
    WorkDir = "/data/"

    /// 全局文件目录
    ArgsFile  = WorkDir + "args.txt"
    StdinFile = WorkDir + "stdin.txt"
    RunFile   = WorkDir + "run.json"
    OutFile   = WorkDir + "out.json"

    /// C 二进制文件
    CLangADotOutFile = "/data/a.out"

    /// C++ 二进制文件
    CppLangADotOutFile = "/data/a.out"

    /// Rust 二进制文件
    RustOutFile = "/data/rust"
)
