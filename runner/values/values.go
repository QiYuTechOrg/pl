package values

const (
    // 全局工作目录
    WorkDir = "/data/"

    // 全局文件目录
    RunFile = WorkDir + "run.json"
    OutFile = WorkDir + "out.json"

    // C 二进制文件
    CLangADotOutFile = "/data/a.out"

    // C++ 二进制文件
    CppLangADotOutFile = "/data/a.out"

    // Rust 二进制文件
    RustBinFile = "/data/rust"

    // PythonBinary Python 脚本文件
    PythonBinary = "python"
    PythonInFile = "/data/main.py"

    // node 脚本文件
    JavaScriptBinary = "node"
    JavaScriptInFile = "/data/main.js"

    // PHP 脚本文件
    PHPBinary = "php"
    PHPInFile = "/data/main.php"
)
