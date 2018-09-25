pidpath - retrieves the absolute path to the executable of a running process

Usage:

	pidpath <pid>

Example:

	ps aux | grep bash
	vagrant   3249  0.0  0.2  23156  5220 pts/0    Ss   22:57   0:00 -bash

	pidpath 3249
	/bin/bash

Install:

	go get github.com/cirocosta/pidpath

Know more:

	https://ops.tips/blog/macos-pid-absolute-path-and-procfs-exploration/#a-golang-binary-that-suits-linux-and-macos
