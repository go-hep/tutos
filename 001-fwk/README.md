fwk
===

[fwk](https://github.com/go-hep/fwk) is a HEP oriented concurrent framework in `Go`.

## Installation

`fwk`, like any pure-Go package, is `go get` able:

```sh
$ go get github.com/go-hep/fwk/...
```

(yes, with the ellipsis after the slash, to install all the "sub-packages")


## Documentation

The documentation is available on `godoc`:

 http://godoc.org/github.com/go-hep/fwk

## Examples

The `fwk` repository ships with a few examples:

- [tuto-1](github.com/go-hep/fwk/blob/master/examples/fwk-ex-tuto-1/main.go): runs 2 tasks exchanging integers,

- [tuto-2](github.com/go-hep/fwk/blob/master/examples/fwk-ex-tuto-2/main.go): reads integers from an ASCII file and massage them somehow,

- [tuto-3](github.com/go-hep/fwk/blob/master/examples/fwk-ex-tuto-3/main.go): reads integers from an ASCII file, massage them somehow and write them out into an output ASCII file,

- [tuto-4](github.com/go-hep/fwk/blob/master/examples/fwk-ex-tuto-4-write-histo/main.go): fills 3 histograms with random data and write them out into a `rio` file,

- [tuto-5](github.com/go-hep/fwk/blob/master/examples/fwk-ex-tuto-5-read-histo/main.go): reads histograms from a `rio` file,

- [tuto-6](github.com/go-hep/fwk/blob/master/examples/fwk-ex-tuto-6-write-data/main.go): writes event data (simple integers) into a `rio` file,

- [tuto-7](github.com/go-hep/fwk/blob/master/examples/fwk-ex-tuto-7-read-data/main.go): reads event data (simple integers) from a `rio` file and saves them back (after some massaging) into another `rio` file.
