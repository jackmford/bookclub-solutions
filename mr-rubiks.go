package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
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
    diary := Diary {
      Walk: 0,
      Sort: 0,
      Clean: 0,
    }

    for fileScanner.Scan() {
      line := fileScanner.Text()
      task := strings.Split(line, " ")[1]
      value, _ := strconv.Atoi(strings.Split(line, " ")[3])
      if task == "Clean" {
        fmt.Println(value)
        diary.Clean += value
      }

    }
}
