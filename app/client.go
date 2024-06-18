// +build client 

package main

import (
    "fmt"
    "net"
    "os"
)

func main() {
    conn, err := net.Dial("tcp", ":6379")
    if err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        os.Exit(1)
    }
    defer conn.Close()

    _, err = conn.Write([]byte("PING\r\n"))
    if err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        os.Exit(1)
    }

    buf := make([]byte, 128)
    _, err = conn.Read(buf)
    if err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        os.Exit(1)
    }

    fmt.Printf("response: %s\n", buf)
}
