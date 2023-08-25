package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Ini Sentence", "Sen"))
	fmt.Println(strings.Split("Ini Sentence", " ")[0])
	fmt.Println(strings.ToLower("Ini Sentence"))
	fmt.Println(strings.ToUpper("Ini Sentence"))
	fmt.Println(strings.Trim("   Ini Sentence   ", " "))
	fmt.Println(strings.ReplaceAll("Ini Sentence", "Sen", "SEN"))

	parseBoolResult, _ := strconv.ParseBool("true")
	fmt.Println(parseBoolResult)

	parseIntResult, _ := strconv.ParseInt("100", 10, 64)
	fmt.Println(parseIntResult)

	parseFloatResult, _ := strconv.ParseFloat("3.14", 64)
	fmt.Println(parseFloatResult)

	fmt.Println(strconv.FormatBool(true))
	fmt.Println(strconv.FormatInt(100, 10))
	fmt.Println(strconv.FormatFloat(3.14, 'f', 2, 64))

}
