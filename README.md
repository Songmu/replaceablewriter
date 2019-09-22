replaceablewriter
=======

[![Build Status](https://travis-ci.org/Songmu/replaceablewriter.svg?branch=master)][travis]
[![Coverage Status](https://coveralls.io/repos/Songmu/replaceablewriter/badge.svg?branch=master)][coveralls]
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)][license]
[![GoDoc](https://godoc.org/github.com/Songmu/replaceablewriter?status.svg)][godoc]

[travis]: https://travis-ci.org/Songmu/replaceablewriter
[coveralls]: https://coveralls.io/r/Songmu/replaceablewriter?branch=master
[license]: https://github.com/Songmu/replaceablewriter/blob/master/LICENSE
[godoc]: https://godoc.org/github.com/Songmu/replaceablewriter

replaceablewriter replace internal io.Writer

## Synopsis

```go
f, _ := os.Open("...")
w := replaceablewriter.New(f)
defer w.Close()
w.Write(...)

// replace
f2, _ := os.Open("...")
w.Replace(f2)
w.Write(...)
```

## Description

replaceablewriter is a library that keeps io.Wrter inside, provides io.WriteCloser interface,
and allows us to replace internal io.Writer safely.

It is useful for log rotation, etc.

## Installation

```console
% go get github.com/Songmu/replaceablewriter
```

## Author

[Songmu](https://github.com/Songmu)
