package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func readfile(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
func main() {
	lines := readfile("/Users/jkelly/Projects/adventforcode/2024/day-02-part-02/large-input.txt")
	total := 0
	for _, line := range lines {
		total += parse(line)
	}
	fmt.Println(total)
}
func parse(parse string) int {
	//r := regexp.MustCompile(`(?P<Year>\d{4})-(?P<Month>\d{2})-(?P<Day>\d{2})`)
	//fmt.Printf("%#v\n", r.FindStringSubmatch(`2015-05-27`))
	//fmt.Printf("%#v\n", r.SubexpNames())
	r := regexp.MustCompile("mul\\((?P<left>[0-9]+),(?P<right>[0-9]+)\\)")
	parsed_result := r.FindAllStringSubmatch(parse, -1)
	total := 0
	for i := 0; i < len(parsed_result); i++ {
		multiplier := parsed_result[i]
		left, _ := strconv.Atoi(multiplier[1])
		right, _ := strconv.Atoi(multiplier[2])
		total += left * right
	}
	return total
}
