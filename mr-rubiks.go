package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
)

type Diary struct {
  Walk int
  Sort int
  Clean int
}

func main() {
    file, err := os.Open(os.Args[1])
    if err != nil {
        fmt.Println(err)
    }

    defer file.Close()

    fileScanner := bufio.NewScanner(file)

    fileScanner.Split(bufio.ScanLines)

    for fileScanner.Scan() {
      line := fileScanner.Text()
      fmt.Println(strings.Split(line, " ")[1])
    }
}
