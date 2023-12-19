# staticweb
Standalone static web server for testing purposes.

This web server shares everything from current directory.

This is written in Go.

## Install

```shell
go install github.com/SeppoTakalo/staticweb
```

## Usage

```shell
cd <directory I want to share>
staticweb -p <port>
```

## Example

```shell
$ cd html
$ staticweb -p 3000
2023/12/19 15:05:08 Listening on :3000
```
