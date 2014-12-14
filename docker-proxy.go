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
        buf = make([]byte, 1024)
        so.Read(buf)
        if err != nil {
            check_warn(err)
        } else {
            n := bytes.Index(buf, []byte{0})         
            doc_socket, err := net.Dial("unix", "/var/run/docker.sock")
            check_strict(err)
            _, err = doc_socket.Write(buf[:n])
            check_warn(err)
            doc_socket.Close()
            fmt.Println(string(buf[:n]))
            so.Close()
        }
    } 
}

func main() {
    portPtr := flag.String("port", "9999", "listenting port")
    flag.Parse()
    server(*portPtr)
}
