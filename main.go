package main

import (
	"fmt"
	"time"
)

func main() {
    t := time.Now()
    fmt.Println(t)
    fmt.Println(t.Format(time.RFC3339))
}