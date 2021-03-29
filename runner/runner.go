package main

import "runner/cmd"

/// Runner 代理
/// 主要目的是为了更好、更方便的运行指定的程序
func main() {
	cmd.Execute()
}
