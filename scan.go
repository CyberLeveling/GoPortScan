package main

import (
        "fmt"
        "net"
        "sync"
        "time"
        "os"
)

var wg sync.WaitGroup

func portScan(ip string, port int, timeout time.Duration) {
        defer wg.Done()

        target := fmt.Sprintf("%s:%d", ip, port)
        conn, err := net.DialTimeout("tcp", target, timeout)
        if err != nil {
                if err, ok := err.(net.Error); ok && err.Timeout() {
                        fmt.Println("Timed out")
                }
                return
        }
        conn.Close()
        fmt.Printf("Port %d is open\n", port)
}

func main() {
        if len(os.Args) != 2 {
                fmt.Println("Invalid argument.")
                fmt.Println("Syntax: go run scanner.go <ip>")
                return
        }

        ip := os.Args[1]
        timeout := time.Second

        for i := 1; i <= 65536; i++ {
                wg.Add(1)
                go portScan(ip, i, timeout)
        }
        wg.Wait()
}
