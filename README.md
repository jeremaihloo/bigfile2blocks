# bigfile2blocks

A client to devide bigfile into small blocks.

## Installation

```bash
go get github.com/jeremaihloo/bigfile2blocks
```

## Usage

```bash
usage: bigfile2blocks --bigfile=BIGFILE --outputs=OUTPUTS [<flags>] <command> [<args> ...]

A command-line bigfile2blocks application.

Flags:
      --help                Show context-sensitive help (also try --help-long and --help-man).
      --bigfile=BIGFILE     Big file path.
      --outputs=OUTPUTS     Number of packets to send
      --hash                Big file hash to check
      --debug               Enable debug mode.
  -t, --timeout=5s          Timeout waiting for ping.
      --block-size=1048576  Block size
      --block-ext=".block"  Block extension name
      --version             Show application version.

Commands:
  help [<command>...]
    Show help.

  blocks
    Blocks a big file

  combine
    Combine blocks into a big file
```

## Api

```golang
import "github.com/jeremaihloo/bigfile2blocks/cores"
```

## TODO

- ~~Fix hash file feature~~
- Goroutine suport for fast blocks
- tests

## License

The MIT License (MIT)

Copyright (c) 2017 jeremaihloo

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.