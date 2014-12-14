package main

import "fmt"
import "net"
import "flag"
import "bytes"

func check_strict(e error) {
    if e != nil {
        panic(e)
    }
}

func check_warn(e error) {
    if e != nil {
        fmt.Println(e)
    }
}

func server(port string) {
    conn, err := net.Listen("tcp", ":"+port)
    defer conn.Close()
    buf := make([]byte, 1024)
    check_strict(err)
    for {
        so, err := conn.Accept()
        check_warn(err)
        for {
            buf = make([]byte, 1024)
            so.Read(buf)
            if err != nil {
                check_warn(err)
            } else {
                n := bytes.Index(buf, []byte{0})         
                fmt.Println(string(buf[:n]))
            }
        }
        so.Close()
    } 
}



func main() {
    portPtr := flag.String("port", "9999", "listenting port")
    flag.Parse()
    server(*portPtr)
}
