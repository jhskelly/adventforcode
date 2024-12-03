package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
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
	//total += parse("xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))")

	all := ""
	for _, line := range lines {
		all += line
	}
	fmt.Println(parse(all))
}
func parse(parse string) int {
	r := regexp.MustCompile("mul\\(|don\\'t\\(\\)|do\\(\\)")
	extract_value_regex := regexp.MustCompile("^mul\\((?P<left>[0-9]+),(?P<right>[0-9]+)\\)")
	parsed_result := r.FindAllStringSubmatchIndex(parse, -1)
	var take = true
	total := 0
	for i := 0; i < len(parsed_result); i++ {
		current := parse[parsed_result[i][0]:len(parse)]
		if strings.HasPrefix(current, "mul(") && take {
			parsed_result := extract_value_regex.FindAllStringSubmatch(current, -1)
			if len(parsed_result) == 1 && len(parsed_result[0]) == 3 {
				left, _ := strconv.Atoi(parsed_result[0][1])
				right, _ := strconv.Atoi(parsed_result[0][2])
				fmt.Println(left, right)
				total += (left * right)
			}
		}
		if strings.HasPrefix(current, "don't()") {
			take = false
		}

		if strings.HasPrefix(current, "do()") {
			take = true
		}
	}
	return total
}
