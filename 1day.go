package main

import (
    "fmt"
    "log"
	"strings"
	"strconv"
	"os"
	"sort"
)

func main() {

    content, err := os.ReadFile("1day.txt")

     if err != nil {
          log.Fatal(err)
     }

	 calories := string(content)
	 splittedCalories := strings.Split(calories, "\n")

	 fmt.Println(splittedCalories)
	var merged []int

	tmp := 0

	 for _, calory := range splittedCalories {
		if calory != "" {
			i, _ := strconv.Atoi(calory)
			tmp += i
		} else {
			merged = append(merged, tmp)
			tmp = 0
		}
	}

	

	max := 0
	sort.Ints(merged)

	fmt.Println(merged)
}