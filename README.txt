Welcome to Go-- , a fork of Go that exports all symbols, so you can access EVERYTHING, like Python!
Remember, with great power comes great responsibility.  Please don't complain if you're using a private API and it changes without notice.

Go-- just consists of several single-line edits to the Go src/ tree.
Here is how you can easily see the diff:

    git diff go1.10.2 go--1.10.2 go/


Here is a quick summary of how to build Go--.  It's exactly the same process as building normal Go:

    # Note that, since Go is written in Go, you need to already have Go to build Go or Go--.
    # So you might find it convenient to download one of the binary distributions from https://golang.org/dl/
    # Use the GOROOT of your already-installed Go for the GOROOT_BOOTSTRAP value.  You can use "go env GOROOT" to find this value:
    export GOROOT_BOOTSTRAP="$(go env GOROOT)"
    export PATH="$GOROOT_BOOTSTRAP/bin:$PATH"
    which go   # Make sure we're using the bootstrap version of go.

    git clone https://github.com/likebike/go--
    cd go--
    git checkout go--1.10.2
    cd go/src
    ./make.bash

    # That's it!  The result will be located at ./go/bin/go .  Note that it's still called 'go' even though it's go-- .
    # I don't suggest that you re-name it to 'go--' because most tools expect it to be called 'go'.  Just set your PATH instead.

    # Here is the standard method to install the amazing monkey patch library:
    GOPATH=/path/to/your/GoPathDir go--/bin/go get github.com/bouk/monkey

    # We also have a copy of the monkey-patch library which matches our version of Go-- in case something changes or breaks.
    # You'd install it like this:
    #
    #   cd go--
    #   GOPATH=$(go env GOPATH)
    #   mkdir -p $GOPATH/src/github.com/bouk
    #   mv monkey $GOPATH/src/github.com/bouk/
    #   cd $GOPATH/src/github.com/bouk/monkey
    #   go install


Here is how we upgrade when new versions of Go are released:

    OLD_VERSION=1.10
    NEW_VERSION=1.10.2

    # First, generate the current Go-- diff:
    git checkout go--
    git diff upstream go/ >diffs/v$OLD_VERSION
    git add diffs/v$OLD_VERSION
    git commit -m"Generate diff of v$OLD_VERSION before upgrade to v$NEW_VERSION"

    # Next, upgrade the 'upstream' branch:
    git checkout upstream
    rm -r go
    curl "https://dl.google.com/go/go${NEW_VERSION}.src.tar.gz" | tar x     # Creates the new 'go' directory.
    git diff                            # Review changes briefly, just to get an idea of their scope.
    git add go                          # Add new files.
    git status                          # Make sure nothing is forgotten.
    git commit -m"Upgrade to v$NEW_VERSION"

    # Re-apply Go-- edits.  You need to re-think each one while applying it:
    git checkout go--
    git merge upstream
    vim diff/v$OLD_VERSION   # Use a vsplit and Manually review/re-apply edits from diff.
    git diff                 # Review
    git diff upstream        # Review
    git diff upstream go/ >diffs/v$NEW_VERSION
    git add diffs/v$NEW_VERSION
    git commit -a -m"Upgrade to v$NEW_VERSION"
    git push --all
    git tag go--$NEW_VERSION
    git push --tags


Other Notes:
    # I'm in China, so I often need to proxy my connections:
    ALL_PROXY='socks5h://127.0.0.1:1080' curl -O 'https://dl.google.com/go/go1.10.2.src.tar.gz'

