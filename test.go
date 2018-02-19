package main

import (
    "fmt"
    "net/http"
    "os"  // Unused imports now generate Warnings instead of Errors.
)

func main() {
    fmt.Println("Welcome to Go--")

    http.Header{}.clone()
    fmt.Println("You can access un-exported symbols.")

    unused:=5
    fmt.Println("Un-used variables and imports now generate Warnings instead of Errors.")

//I figured out how to enable this (edit go/src/cmd/compile/internal/syntax/parser.go),
//but it's prohibitively complicated.  I'd need to re-implement a major part of the
//parser infrastructure to support 'unnext()' instead of just the current 'ungetr()'.
//
//  if false { }
//  else { fmt.Println("You can now begin a line with 'else'.") }

    fmt.Println("Have fun!")
}
