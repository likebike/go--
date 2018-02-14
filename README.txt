THIS README IS UNDER CONSTRUCTION




# Find the latest Go Source URL from https://golang.org/dl/
# I'm in China, so I need to proxy my request:
ALL_PROXY='socks5h://127.0.0.1:8080' curl 'https://dl.google.com/go/go1.9.4.src.tar.gz' >go1.9.4.src.tar.gz
tar xf go1.9.4.src.tar.gz 
cd go/src












# Installation:
ROOT=$HOME/koin

# Install Go:
mkdir -p $ROOT/env/.src
cd $ROOT/env/.src/
wget 'https://storage.googleapis.com/golang/go1.5.1.linux-amd64.tar.gz'
tar xf go1.5.1.linux-amd64.tar.gz
mv go ../

# Install Go Libs (Note, github.com is blocked in China, so you might need to use a VPN or put an entry in /etc/hosts to run these):
GOROOT=$ROOT/env/go GOPATH=$ROOT/env/gopath $ROOT/env/go/bin/go get github.com/mattn/go-sqlite3 






####### Building a custom version of Go that allows you to access unexported symbols (i.e. Go--) (Go-- Build Instructions):
#
#  If your libraries are well-designed and they give you what you need, then life is great.
#  But if the library is good but they things you need are un-exported, you're out of luck.
#  I had too much of a hassle using net/http.  The only way I was able to access the
#  data I needed was by making a copy-paste of the entire net/http directory and
#  placing it into my own workspace, and making edits to expose the required data.
#  This is a very intricate and error-prone process.  I really don't think i could
#  do it again perfectly 6 months from now... let alone anybody else being able to do this right.
#  Also, it's a maintenance headache,,, soooo many diffs!
#  So, that's why i'm interested increating a hacked version of Go that just
#  gives me access to everything, whether it is exposed or not.
#
#  Here are the basic instructions for building Go from source:
#      https://golang.org/doc/install/source
#
#  Go is surprisingly simple to build:
#      export GOROOT_BOOTSTRAP="$ROOT/env/go"
#      export PATH="$GOROOT_BOOTSTRAP/bin:$PATH"
#      # Fetch the git repo as described in the above documentation link.
#      cd go/src
#      ./make.bash  # OR ./all.bash if you want to run tests.
#
#  Unfortunately, the 'Export' logic is scattered throughout the codebase
#  because the logic is very simple (uppercase first letter).  So we need
#  to track down the important places:
#      grep -ri export . | less    ### Produces lots of noise, but still necessary.
#      grep -ri unicode.IsUpper .  ### Really good list.
#
#  As of Go-1.9, I found these important files:
#      go/ast/ast.go
#      cmd/compile/internal/gc/export.go
#      cmd/compile/internal/gc/reflect.go
#      cmd/link/internal/ld/deadcode.go
#      encoding/json/encode.go
#
#  There are also places in the code where we definitely do NOT want to affect exports:
#      encoding/gob/type.go
#      net/rpc/server.go
#      cmd/doc/main.go
#      cmd/compile/internal/syntax/dumper.go
#      reflect/value.go
#      reflect/type.go
#      go/internal/gcimporter/bimport.go
#      cmd/compile/internal/gc/bexport.go
#
#  In the important files, find the function that deals with exports, something like this:
#      func IsExported(name string) bool {
#          ch, _ := utf8.DecodeRuneInString(name)
#          return unicode.IsUpper(ch)
#      }
#
#  ...And short-circuit the logic so that it always just returns 'true'.
#  In cmd/compile/internal/gc/export.go you instead need to short-circuit exportname() with:
#      return !strings.Contains(s, ".")
#
#  Test the resulting compiler with something like this:
#      package main
#      import "fmt"; import "net/http"
#      func main() { fmt.Printf("Hello Unexported World! %#v %#v\n", http.NotFound, http.conn{}) }








####################################################################################
Here is how I set up our production server:

--- Env Setup ---

We need some tools before we can build our project:
    Go:
        ROOT=$HOME/koin
        mkdir -p $ROOT/env/.src
        cd $ROOT/env/.src/
        wget 'https://storage.googleapis.com/golang/go1.9.1.linux-amd64.tar.gz'
        tar xf go1.9*
        mv go ../go_1.9.1
        cd ..
        ln -s go_1.9.1 go
        mkdir gopath_1.9.1
        ln -s gopath_1.9.1 gopath
        GOPATH=$ROOT/env/gopath $ROOT/env/go/bin/go get github.com/mattn/go-sqlite3   # Also install the sqlite bindings.

    Go--:
        cd $ROOT/env/.src/
        git clone https://go.googlesource.com/go     # Note the potential directory name conflict with the above 'Go' instructions!
        cd go
        git checkout go1.9.1
        cd src
        # Edit the src files as described in the above "Go-- Build Instructions".  Specifically:
        #   go/ast/ast.go: func IsExported(name string) bool { return true // Edited by Christopher Sebastian
        #   cmd/compile/internal/gc/export.go: func exportname(s string) bool { return !strings.Contains(s, ".") // Edited by Christopher Sebastian.  You also need to "import strings" for this one.
        #   cmd/compile/internal/gc/reflect.go: func isExportedField(ft *types.Field) (bool, *types.Pkg) { return true, nil // Edited by Christopher Sebastian
        #   cmd/link/internal/ld/deadcode.go: func (m methodref) isExported() bool { return true // Edited by Christopher Sebastian
        #   encoding/json/encode.go: func isExported(id string) bool { return true // Edited by Christopher Sebastian
        vim go/ast/ast.go cmd/compile/internal/gc/export.go cmd/compile/internal/gc/reflect.go cmd/link/internal/ld/deadcode.go encoding/json/encode.go
        export GOROOT_BOOTSTRAP="$ROOT/env/go"
        export PATH="$GOROOT_BOOTSTRAP/bin:$PATH"
        which go   # Make sure we're using the right 'go'.
        ./make.bash
        cd $ROOT/env/.src/
        mv go ../go--_1.9.1
        cd ..
        ln -s go--_1.9.1 go--
        GOPATH=$ROOT/env/gopath $ROOT/env/go--/bin/go get github.com/bouk/monkey   # Also install monkey patching support.
        GOPATH=$ROOT/env/gopath $ROOT/env/go--/bin/go get github.com/NYTimes/gziphandler    # GZip Compression Middleware.






