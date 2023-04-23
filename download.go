package main

import (
    "fmt"
    "io"
    "net/http"
    "os"

    pb "gopkg.in/cheggaaa/pb.v1"
)

func main() {
    url := os.Args[1]
    path := os.Args[2]

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

    bar := pb.StartNew(int(resp.ContentLength))

    _, err = io.Copy(out, io.TeeReader(resp.Body, bar))
    if err != nil {
        fmt.Println("Error downloading file:", err)
        return
    }

    bar.FinishPrint("Download has been completed.")
}
