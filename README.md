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

```sh
go get github.com/echocrow/pidpath
```

## Credit

This is a fork of [pidpath by cirocosta](https://github.com/cirocosta/pidpath), updated for go modules.
