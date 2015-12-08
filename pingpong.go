package main

import "fmt"

func main() {
    for out := range PingPong(200) {
        fmt.Println(out)
    }
}

func PingPong(limit int) <-chan string {

    out := make(chan string, limit)

    go func() {
        for i := 1; i <= limit; i++ {
            result := ""
            if i%3 == 0 { result += "Ping" }
            if i%4 == 0 { result += "Pong" }
            if result == "" { result = fmt.Sprintf("%v", i) }
            out <- result
        }
        close(out)
    }()

    return out
}
