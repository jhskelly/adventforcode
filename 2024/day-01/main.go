package main

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.
import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
func main() {
	var lines []string = readfile("/Users/jkelly/Projects/adventforcode/2024/day-01/small-input.txt")
	length := len(lines)
	left := []int{}
	right := []int{}

	for i := 0; i < length; i++ {
		fmt.Println(lines[i])
		split := strings.Split(lines[i], "   ")
		leftint, err := strconv.Atoi(split[0])
		check(err)
		rightint, err := strconv.Atoi(split[1])
		check(err)
		left = append(left, leftint)
		right = append(right, rightint)

	}
	sort.Ints(left)
	sort.Ints(right)
	total := 0
	for i := 0; i < length; i++ {
		difference := (left[i] - right[i])
		if difference < 0 {
			difference = difference * -1
		}
		total += difference
	}
	fmt.Println(total)
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
