#!go/bin/go run
//^^^^^^^^^^^^^ "shebang" support enables you to run Go programs as scripts.
//              I usually like to say something like "#!/usr/bin/env go run",
//              but on Linux a shebang command can only have ONE argument, so
//              the kernel tries to run "go run" as a command, which fails.
//              (BSD and MacOSX don't have this problem.)
//              We *could* create a "go run" binary as part of Go-- to enable this
//              in a cross-platform sort of way, but for now I just use this
//              relative path.  You'll need to run from this directory.

package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "os"             // Unused imports now generate Warnings instead of Errors.
)

type T1 struct { n1,n2 int }

func                       // I'm not saying that this is actually a good way to style your function definitions...
(t T1)                     // I'm only saying that you should be allowed to do it if you want to.
crazy_func_style
(in1,in2 int)
(out1,out2 int)
{
    return 3,4
}

func main() {
    fmt.Println("\nWelcome to Go--.  Don't worry about those warnings that you see above -- they are a *feature*, not a bug!")
    fmt.Println("Go-- gives you some extra abilities, compared to normal Go:\n")

    unused:=5
    fmt.Println("    * Unused variables and imports now generate Warnings (like you see above) instead of Errors.\n")

    http.Header{}.clone(); _=http.response{}.status; http.htmlEscape("We can access <em>everything</em>!")
    fmt.Println("    * You can access unexported symbols.\n")

    var a,b T1; a.n1,a.n2=8,9
    bs,err:=json.Marshal(a); if err!=nil { panic(err) }
    err=json.Unmarshal(bs, &b); if err!=nil { panic(err) }
    if b.n1!=8 || b.n2!=9 { panic("json was unable to read/set unexported field!") }
    fmt.Println("    * encoding/json can un/marshal unexported fields.\n")

    if 0>1 { }  // Comments here are fine too.
    else { fmt.Println("    * More flexible whitespace in 'if' statements; `if cond {...} \\n else {...}` is now a valid syntax.\n") }

    three,four:=a.crazy_func_style(1,2); if three!=3 || four!=4 { panic("Incorrect crazy_func_style outputs!") }
    fmt.Println("    * More flexible whitespace in function definitions.\n")

    fmt.Println("    * You can use a 'shebang' line (#!...) to run your Go-- programs as scripts.\n")

    fmt.Println("Remember: with great power comes great responsiblity.  Have fun!\n")
}

