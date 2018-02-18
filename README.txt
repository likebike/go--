Welcome to Go-- , a fork of Go that exports all symbols, so you can access EVERYTHING, like Python!
Remember, with great power comes great responsibility.  Please don't complain if you're using a private API and it changes without notice.

Go-- just consists of several single-line edits to the Go src/ tree.
Here is how you can easily see the diff:

    git diff go1.9.4 go--1.9.4 go/


Here is a quick summary of how to build Go--.  It's exactly the same process as building normal Go:

    # Note that, since Go is written in Go, you need to already have Go to build Go or Go--.
    # So you might find it convenient to download one of the binary distributions from https://golang.org/dl/
    # Use the GOROOT of your already-installed Go for the GOROOT_BOOTSTRAP value.  You can use "go env GOROOT" to find this value:
    export GOROOT_BOOTSTRAP="$(go env GOROOT)"
    export PATH="$GOROOT_BOOTSTRAP/bin:$PATH"
    which go   # Make sure we're using the bootstrap version of go.

    git clone https://github.com/likebike/go--
    cd go--
    git checkout go--1.9.4
    cd go/src
    ./make.bash

    # That's it!  The result will be located at ./go/bin/go .  Note that it's still called 'go' even though it's go-- .
    # I don't suggest that you re-name it to 'go--' because most tools expect it to be called 'go'.  Just set your PATH instead.

    # Here is the standard method to install the amazing monkey patch library:
    GOPATH=/path/to/your/GoPathDir go--/bin/go get github.com/bouk/monkey

    # We also have a copy of the monkey-patch library which matches our version of Go-- in case something changes or breaks.
    # It is located in the 'upstream' branch.  You'd install it like this:
    #
    #   cd go--
    #   git checkout upstream
    #   GOPATH=$(go env GOPATH)
    #   mkdir -p $GOPATH/src/github.com/bouk
    #   mv monkey $GOPATH/src/github.com/bouk/
    #   cd $GOPATH/src/github.com/bouk/monkey
    #   go install



Other Notes:
    # I'm in China, so I often need to proxy my connections:
    ALL_PROXY='socks5h://127.0.0.1:8080' curl -O 'https://dl.google.com/go/go1.9.4.src.tar.gz'

