go-installation
===============

The installation of a `Go` toolchain is better explained at [golang.org](https://golang.org/doc/install)

## Installing `Go` manually

For people in a hurry, here are a couple of copy-paste instructions,
for a `Linux` based environment:

```sh
$ mkdir -p $HOME/sdk
$ cd $HOME/sdk
$ curl -L https://golang.org/dl/go1.12.linux-amd64.tar.gz | tar zxf -
$ export GOROOT=$HOME/sdk/go
$ export PATH=$GOROOT/bin:$PATH

$ which go
~/sdk/go/bin/go
```

At the time of writing, the latest version of Go is `1.12`.
You should always try to download and install the latest one.

## Setting up the work environment

Like `python` and its `$PYTHONPATH` environment variable, `Go` uses
`$GOPATH` to locate packages' source trees.
You can choose whatever you like (obviously a directory under which
you have read/write access, though.)
In the following, we'll assume you chose `$HOME/go`:

```sh
$ mkdir -p $HOME/go
$ export GOPATH=$HOME/go
$ export PATH=$GOPATH/bin:$PATH
```

_Note: `$GOPATH` will probably go away with Go1.13._

Make sure the `go` tool is correctly setup:

```sh
$ go env
GOARCH="amd64"
GOBIN=""
GOCHAR="6"
GOEXE=""
GOHOSTARCH="amd64"
GOHOSTOS="linux"
GOOS="linux"
GOPATH="$HOME/go"
GORACE=""
GOROOT="$HOME/sdk/go"
GOTOOLDIR="$HOME/sdk/go/pkg/tool/linux_amd64"
CC="gcc"
GOGCCFLAGS="-fPIC -m64 -pthread -fmessage-length=0"
CXX="g++"
CGO_ENABLED="1"
```

(on other platforms/architectures, the output might differ slightly.
The import environment variables are `GOPATH` and `GOROOT`.)

## Testing `go get`

Now that the `go` tool is correctly setup, let's try to fetch some
code.
For this part, you'll need `git` (or `mercurial` for some
repositories) to actually retrieve the code from the remote
repositories.

Without further ado:

```sh
$ go get -v github.com/go-hep/tutos/000-go-installation/tuto-hello
github.com/go-hep/tutos (download)
github.com/go-hep/tutos/000-go-installation/tuto-hello
```

`go get` downloaded (cloned, in `git` speak) the whole
`github.com/go-hep/tutos` repository (under `$GOPATH/src`) and
compiled the `tuto-hello` command.
As the compilation was successful, it also installed the `tuto-hello`
command under `$GOPATH/bin`.

The `tuto-hello` command is now available from your shell:

```sh
$ tuto-hello
Hello from go-hep/tutos!

$ tuto-hello you
Hello you!
```

## Setting up your favorite editor

Extensive documentation on how to setup your editor (for code
highlighting, code completion, ...) is available here:

 https://github.com/golang/go/wiki/IDEsAndTextEditorPlugins
 
At the very least, you should try to install and setup `goimports` as
explained here:

 https://godoc.org/golang.org/x/tools/cmd/goimports

`goimports` provides automatic code formating as well as automated
insertion/deletion of used/unused packages (in your `import` package
statements.)

## Documentation

The `Go` programming language is quite new (released in 2009) but
ships already with quite a fair amount of documentation.
Here are a few pointers:

- https://golang.org/doc/code.html
- https://tour.golang.org
- https://golang.org/doc/effective_go.html
- https://dave.cheney.net/resources-for-new-go-programmers
- https://dave.cheney.net/practical-go/presentations/qcon-china.html
- https://gobyexample.com

For more advanced topics:

- https://talks.golang.org
- https://blog.golang.org
- https://groups.google.com/d/forum/golang-nuts (`Go` community forum)
- the internets. typical queries are `go something` or `golang something`

