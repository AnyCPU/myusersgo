[![Build Status](https://travis-ci.org/AnyCPU/myusersgo.svg?branch=master)](https://travis-ci.org/AnyCPU/myusersgo)
# My users go

## Version 1.0.0

## Description
The `My users go` utility helps to back your MySql users up. The utility was tested on both MySql 5.6.* and MySql 5.7.* versions, also may work with the other versions.

## Installation
* Install [Go Language](https://golang.org/)
* Set `GOPATH` environment variable if use Go 1.7 and below
* Install [MySql driver](https://github.com/go-sql-driver/mysql) - `go get github.com/go-sql-driver/mysql`
* Run `build.bat` or `go build -ldflags "-s -w" -o myusersgo.exe myusers.go`
* Dump your users `./myusersgo.exe <params>` (on Unix) or `.\myusersgo.exe <params>` (on Windows)

## Command line
* `-u <name>` - user name (default "root")
* `-p <password>` - user password (default "12345")
* `-s <ip>` - db server ip (default "127.0.0.1")
* `--port <port>` - db server port (default "3306")
* `-d <name>` - db name with users (default "mysql")
* `-v` - show version information and exit
* `-h` - show help and exit

## License
```
BSD 2-Clause License

Copyright (c) 2017, M1xA
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

* Redistributions of source code must retain the above copyright notice, this
  list of conditions and the following disclaimer.

* Redistributions in binary form must reproduce the above copyright notice,
  this list of conditions and the following disclaimer in the documentation
  and/or other materials provided with the distribution.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
```

## Changelog
* __1.0.0__ - 2017.02.08
  * First release
