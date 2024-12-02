package main

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.
import (
	"bufio"
	"fmt"
	"os"
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
func check(err error) {
	if err != nil {
		panic(err)
	}
}

func abs(value int) int {
	if value < 0 {
		return value * -1
	}
	return value
}
func stringArrayToIntArray(values []string) []int {
	var result []int
	for _, i2 := range values {
		atoi, _ := strconv.Atoi(i2)
		result = append(result, atoi)
	}
	return result

}

func main() {
	var lines []string = readfile("/Users/jkelly/Projects/adventforcode/2024/day-02-part-01/large-input.txt")

	total := 0
	for i := 0; i < len(lines); i++ {
		fmt.Println(lines[i])
		split_text := strings.Split(lines[i], " ")
		split_ints := stringArrayToIntArray(split_text)
		if valid(split_ints) {
			total++
		}
	}
	fmt.Println("-----")
	fmt.Println(total)
}

func valid(values []int) bool {
	return validDifferByOneOrTwo(values) && validIsIncreasingOrDecreasing(values)
}
func validIsIncreasingOrDecreasing(values []int) bool {
	multiplier := 1
	if values[0] < values[1] {
		multiplier = -1
	}
	current := values[0]
	for i := 1; i < len(values); i++ {
		if ((current - values[i]) * multiplier) < 0 {
			return false
		}
		current = values[i]
	}

	return true
}
func validDifferByOneOrTwo(values []int) bool {
	current := values[0]
	for i := 1; i < len(values); i++ {
		var diff = abs(current - values[i])
		fmt.Println(current)
		fmt.Println(values[i])
		fmt.Println("---------")
		if !((diff == 1) || (diff == 2) || (diff == 3)) {
			return false
		}
		current = values[i]
	}
	return true
}
