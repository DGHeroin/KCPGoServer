package main

import (
    "github.com/xtaci/kcp-go"
    "log"
    "net"
)

func main()  {
    listener, err :=  kcp.Listen(":6789")
    if err != nil {
        log.Panic(err)
    }
    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Println(err)
            break
        }
        log.Println("接收新链接")
        go serveConn( conn )
    }
}

func serveConn (conn net.Conn)  {
    defer func() {
        log.Printf("连接关闭: %v\n", conn)
    }()
    buf := make([]byte, 2048)
    for {
        n, err := conn.Read(buf)
        if err != nil {
            log.Printf("读取错误: %v\n", err)
            break
        }
        msg := buf[:n]
        log.Printf("recv: %s\n", msg)
        conn.Write(msg)
    }
}