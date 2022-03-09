package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "strings"
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
    total_hrs := 0
    total_mins := 0
    days := 0 

    for fileScanner.Scan() {
        days = days+1
        line := fileScanner.Text()
        fmt.Println(line)
        am_hour := strings.Split(line, " ")[0]
        am_hour_int, _ := strconv.Atoi(strings.Split(am_hour, ":")[0])

        am_or_pm_to_bed := strings.Split(line, " ")[1]
        am_or_pm_to_rise := strings.Split(line, " ")[3]
        fmt.Println(am_or_pm_to_bed)
        fmt.Println(am_or_pm_to_rise)

        pm_hour := strings.Split(line, " ")[2]
        pm_hour_int, _ := strconv.Atoi(strings.Split(pm_hour, ":")[0])
        am_mins, _ := strconv.Atoi(strings.Split(am_hour, ":")[1])
        pm_mins, _ := strconv.Atoi(strings.Split(pm_hour, ":")[1])
        if strings.Split(am_hour, ":")[1] != "00" {
            am_hour_int = am_hour_int+1
            if am_hour_int >= 12 {
              am_hour_int = am_hour_int-12
            }
            total_mins = total_mins+(60-am_mins)
            total_mins = total_mins+(pm_mins)
        }

        if am_or_pm_to_bed == "PM," && am_or_pm_to_rise == "PM" {
            total_hrs = total_hrs+pm_hour_int-am_hour_int
        }

        if am_or_pm_to_bed == "PM," && am_or_pm_to_rise == "AM" {
            total_hrs = total_hrs+(12-am_hour_int)
            total_hrs = total_hrs+pm_hour_int
        }

        if am_or_pm_to_bed == "AM," && am_or_pm_to_rise == "AM" {
            total_hrs = total_hrs+(pm_hour_int-am_hour_int)
        }

        if am_or_pm_to_bed == "AM," && am_or_pm_to_rise == "PM" {
            total_hrs = total_hrs+(pm_hour_int-am_hour_int)
            total_hrs = total_hrs+(12-am_hour_int+pm_hour_int)
        }

        fmt.Println(am_hour_int)
        fmt.Println(pm_hour_int)
        fmt.Println(total_hrs)
        fmt.Println(total_mins)
        var total_time float64
        total_time = float64((total_hrs*60+total_mins))/float64(60)
        total_time = total_time/float64(days)
        fmt.Println(total_time)
        fmt.Printf("%.1f hours", total_time)

    }
}
