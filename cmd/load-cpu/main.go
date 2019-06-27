package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func main() {
	flag1 := true
	sleeptime := 1000
	count := 0
	for flag1 {
		fmt.Printf("count %d sleeptime %d\n", count, sleeptime)
		count++
		time.Sleep(time.Duration(sleeptime) * time.Millisecond)
		dat, err := ioutil.ReadFile("./cmd/load-cpu/dat.txt")
		if err != nil {
			fmt.Println(err)
			return
		}
		datStr := string(dat)
		if strings.Contains(datStr, "s=") {
			sleeptime, err = strconv.Atoi(strings.Split(datStr, "=")[1])
			if err != nil {
				fmt.Println(err)
				return
			}
		}
		if strings.Contains(datStr, "e=") {
			fmt.Println("exiting as flag is set to false")
			flag1 = false
		}
	}
}
