package main // Package main is special. It defines a standalone executable program, not a library.

import "fmt" // Package fmt implements formatted I/O with functions analogous to C's printf and scanf. The format 'verbs' are derived from C's but are simpler.

func main(){
 fmt.Println("Hello, World")// Println formats using the default formats for its operands and writes to standard output. Spaces are always added between operands and a newline is appended. It returns the number of bytes written and any write error encountered.
}
