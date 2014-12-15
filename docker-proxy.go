package main

import "fmt"
import "net"
import "flag"

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
    conn, err := net.Listen("tcp", ":" + port)
    defer conn.Close()

    so_buf := make([]byte, 1024*1024)
    check_strict(err)
    for {
        so, err := conn.Accept()
        check_warn(err)
        so_buf = make([]byte, 1024*1024)
        so_len, err := so.Read(so_buf)
        check_warn(err)
        doc_buf := make([]byte, 1024*1024)
        check_warn(err)
        if err == nil {
            doc_socket, err := net.Dial("unix", "/var/run/docker.sock")
            check_strict(err)
            _, err = doc_socket.Write(so_buf[:so_len])

            doc_len, err := doc_socket.Read(doc_buf)
            fmt.Printf("len is: %d\n", doc_len)
            fmt.Println("reply: "+string(doc_buf[:doc_len]))
            so.Write(doc_buf[:doc_len])

            check_warn(err)
            doc_socket.Close()
            fmt.Println(string(so_buf[:so_len]))
            so.Close()
        }
    } 
}

func main() {
    portPtr := flag.String("port", "9999", "listenting port")
    flag.Parse()
    server(*portPtr)
}
