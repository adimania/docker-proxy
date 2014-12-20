package main

import "fmt"
import "net"
import "flag"

func check_strict(e error) {
	if e != nil {
		panic(e)
	}
}

func check_warn(e error) bool {
	if e != nil {
		fmt.Println(e)
		return true
	}
	return false
}

func server(port string) {
	conn, err := net.Listen("tcp", ":"+port)
	defer conn.Close()
	check_strict(err)
	for {
		so, err := conn.Accept()
		defer so.Close()
		if check_warn(err) {
			continue
		}
		go proxyHandler(so)
	}
}

func proxyHandler(so net.Conn) {
	so_buf := make([]byte, 1024*1024)
	doc_buf := make([]byte, 1024*1024)
	so_len, err := so.Read(so_buf)
	check_strict(err)
	doc_socket, err := net.Dial("unix", "/var/run/docker.sock")
	defer doc_socket.Close()
	check_strict(err)
	_, err = doc_socket.Write(so_buf[:so_len])
	for {
		doc_len, err := doc_socket.Read(doc_buf)
		fmt.Printf("len is: %d\n", doc_len)
		fmt.Println("reply: " + string(doc_buf[:doc_len]))
		so.Write(doc_buf[:doc_len])
		check_strict(err)
		fmt.Println(string(so_buf[:so_len]))
	}
}

func main() {
	portPtr := flag.String("port", "9999", "listenting port")
	flag.Parse()
	server(*portPtr)
}
