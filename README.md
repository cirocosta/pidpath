# pidpath

Retrieve the absolute path to the executable of a running process.

## Usage
```
pidpath <pid>
```

## Example

```sh
ps aux | grep bash
vagrant   3249  0.0  0.2  23156  5220 pts/0    Ss   22:57   0:00 -bash

pidpath 3249
/bin/bash
```

## Install

To compile and install the `pidpath` CLI, run:

```sh
go install github.com/cirocosta/pidpath/cmd/pidpath@latest
```

To add `pidpath` to your Go project, import `github.com/cirocosta/pidpath` or run:

```sh
go get github.com/cirocosta/pidpath
```
