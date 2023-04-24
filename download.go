package main

import (
    "fmt"
    "io"
    "net"
    "net/http"
    "os"
    "strings"

    pb "github.com/cheggaaa/pb/v3"
)

func main() {
    if len(os.Args) < 3 {
        fmt.Println("Usage:", os.Args[0], "url", "path")
        return
    }

    url := os.Args[1]
    path := os.Args[2]

    // DNS Resolver
    addrs, err := net.LookupHost(strings.Split(url, "/")[2])
    if err != nil {
        fmt.Println("DNS lookup failed:", err.Error())
        return
    }

    // Select the first IP from the list
    ip := addrs[0]

    // Replace hostname with IP in the URL
    url = strings.Replace(url, strings.Split(url, "/")[2], ip, 1)

    resp, err := http.Get(url)
    if err != nil {
        fmt.Println("Error downloading file:", err)
        return
    }
    defer resp.Body.Close()

    out, err := os.Create(path)
    if err != nil {
        fmt.Println("Error creating file:", err)
        return
    }
    defer out.Close()

    bar := pb.Full.Start64(resp.ContentLength)

    _, err = io.Copy(out, bar.NewProxyReader(resp.Body))
    if err != nil {
        fmt.Println("Error downloading file:", err)
        return
    }

    bar.Finish()
    fmt.Println("Download has been completed.")
}
