package main
import (
    "io"
    "log"
    "net"
    "fmt"
)

const listenAddr = "localhost:4000"

func main() {
    l, err := net.Listen("tcp", listenAddr)
    if err != nil {
        log.Fatal(err)
    }
    for {
        c, err := l.Accept()
        if err != nil {
            log.Fatal(err)
        }
        go match(c)
    }
}
var partner = make(chan io.ReadWriteCloser)

func match(c io.ReadWriteCloser) {
    fmt.Fprint(c, "Waiting for a chat Buddy")
    select {
    case partner <- c:
    case p := <-partner:
        chat(p, c)
    }
}
func chat(a, b io.ReadWriteCloser) {
    fmt.Fprintln(a, "Found one! Say hi.")
    fmt.Fprintln(b, "Found one! Say hi.")
    go io.Copy(a, b)
    io.Copy(b, a)
}

