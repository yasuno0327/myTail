# myTail

## Overview

## Usage

Install command

```shell
go get github.com/yasuno0327/myTail
```

Show last 10 lines of file.

```shell
myTail hogehoge.txt
```

### n option

Show last n lines of file.

```shell
myTail -n=2 hogehoge.txt
```

### w option

Monitor the file

```shell
myTail -f hogehoge.txt
```

### Output many files

Show last 2 lines of two file.

```
myTail -n=2 hogehoge.txt fugafuga.txt
```
