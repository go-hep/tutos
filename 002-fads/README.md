fads
====

[fads](https://go-hep.org/x/hep/fads) is a `FAst Detector Simulation` toolkit.

## Installation

Installing `fads` is done like for any other `Go` based application or
package:

```sh
$ go get -u -v go-hep.org/x/hep/fads
```

Alternatively, adding the ellipsis `...` instructs the `go get` tool
to also compile and install all the applications and packages *under*
`go-hep/fads`:

```sh
$ go get -u -v go-hep.org/x/hep/fads/...
```

`fads` is a toolkit; all the code directly under `/fads` is a `Go`
package for implementing a fast simulation for HEP.

An example application `fads-app` is provided: [cmd/fads-app/main.go](https://github.com/go-hep/hep/blob/master/fads/cmd/fads-app/main.go)

Installing & compiling this example application can thus be done like
so:

```sh
$ go get go-hep.org/x/hep/fads/cmd/fads-app
```

The binary should then be available from the shell prompt:

```sh
$ fads-app -help
Usage: fads-app [options] <hepmc-input-file>

ex:
 $ fads-app -l=INFO -evtmax=-1 ./testdata/hepmc.data

options:
  -cpu-prof=false: enable CPU profiling
  -evtmax=-1: number of events to process
  -l="INFO": log level (DEBUG|INFO|WARN|ERROR)
  -nprocs=0: number of concurrent events to process
```

A test `HepMC` file is also shipped with the `go-hep/fads` repository,
under the `testdata` directory:

```sh
$ fads-app $GOPATH/src/go-hep.org/x/hep/fads/testdata/hepmc.data
::: fads-app...
app                  INFO cpu: 1.948031994s
app                  INFO mem: alloc:           3189 kB
app                  INFO mem: tot-alloc:      26807 kB
app                  INFO mem: n-mallocs:      53300
app                  INFO mem: n-frees:        52632
app                  INFO mem: gc-pauses:         51 ms
::: fads-app... [done] (time=1.9540843s)
```

Running with all the available processors (2 on this test machine), on
5 events:

```sh
$ fads-app -nprocs=-1 -evtmax=5 $GOPATH/src/go-hep.org/x/hep/fads/testdata/hepmc.data
::: fads-app...
app                  INFO workers done: 1/2
app                  INFO workers done: 2/2
app                  INFO cpu: 1.15986402s
app                  INFO mem: alloc:           4280 kB
app                  INFO mem: tot-alloc:      26827 kB
app                  INFO mem: n-mallocs:      53380
app                  INFO mem: n-frees:        45305
app                  INFO mem: gc-pauses:         32 ms
::: fads-app... [done] (time=1.165603508s)
```
