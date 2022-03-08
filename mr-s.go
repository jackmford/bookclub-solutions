package main

import (
  "bufio"
  "fmt"
  "os"
)

func main() {
    file, err := os.Open(os.Args[1])
    if err != nil {
        fmt.Println(err)
    }

    defer file.Close()


    fileScanner := bufio.NewScanner(file)
    fileScanner.Split(bufio.ScanLines)

    fileScanner.Scan()

    for fileScanner.Scan() {
        line := fileScanner.Text()
        fmt.Println(line)
    }
}
