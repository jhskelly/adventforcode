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

func fixmap() map[string]int {
	maps := make(map[string]int)
	maps["5 6 7 10 13 16 13"] = 1
	maps["19 21 24 27 28 28"] = 1
	maps["16 18 20 21 23 25 29"] = 1
	maps["44 46 48 49 52 55 56 62"] = 1
	maps["31 32 32 33 36"] = 1
	maps["67 66 68 70 71 74 76"] = 1
	maps["9 6 12 13 15"] = 1
	maps["80 80 81 84 86"] = 1
	maps["20 24 27 28 30 33 35"] = 1
	maps["39 43 40 43 44 46"] = 1
	maps["9 14 17 18 19 21 22 23"] = 1
	maps["34 32 29 27 24 22 20 23"] = 1
	maps["99 97 95 92 90 89 86 86"] = 1
	maps["23 22 21 20 17 16 13 9"] = 1
	maps["77 74 73 71 68 65 64 59"] = 1
	maps["31 29 27 24 25 23"] = 1
	maps["12 11 8 8 5 2 1"] = 1
	maps["40 41 38 37 35 33 32"] = 1
	maps["28 31 26 23 21 18"] = 1
	maps["42 42 41 39 37 35 34 31"] = 1
	maps["20 16 13 12 9 8"] = 1
	maps["70 65 63 62 61 58 56 53"] = 1
	maps["52 55 57 60 63 66 63"] = 1
	maps["10 11 12 13 14 16 16"] = 1
	maps["84 87 88 90 91 95"] = 1
	maps["12 14 17 19 20 21 26"] = 1
	maps["46 49 48 51 54"] = 1
	maps["13 14 17 17 20"] = 1
	maps["34 33 34 35 38 41 42 45"] = 1
	maps["31 31 33 36 38 40 42 45"] = 1
	maps["6 10 11 13 16"] = 1
	maps["72 77 79 80 81 83"] = 1
	maps["35 40 37 38 41"] = 1
	maps["16 13 12 9 8 11"] = 1
	maps["68 67 66 65 62 62"] = 1
	maps["46 45 44 42 40 36"] = 1
	maps["49 48 47 46 43 42 41 34"] = 1
	maps["93 92 91 88 90 88"] = 1
	maps["21 20 20 19 16 14 13"] = 1
	maps["19 17 16 14 10 13"] = 1
	maps["63 66 63 60 57 55"] = 1
	maps["81 82 78 77 75 72 69 68"] = 1
	maps["65 68 62 60 57"] = 1
	maps["42 42 41 39 36 35 34 31"] = 1
	maps["43 39 37 36 34 33 32 29"] = 1
	maps["61 54 52 49 46 45 43"] = 1
	maps["48 43 45 43 42"] = 1
	maps["76 73 72 70 70 68 66 63"] = 1
	maps["47 47 45 44 41 39 38 37"] = 1
	maps["36 38 40 43 46 51"] = 1
	maps["67 66 67 70 72 74 77 80"] = 1
	maps["67 70 71 72 73 75 78 78"] = 1
	maps["87 84 83 81 78 77 76 76"] = 1
	maps["55 53 57 60 63 64"] = 1
	maps["47 49 50 52 53 56 59 63"] = 1
	maps["62 59 57 55 54 51 50 44"] = 1
	maps["12 16 15 17 19"] = 1
	maps["48 43 42 41 40 37"] = 1
	maps["23 20 17 15 13 11 10 6"] = 1
	maps["57 60 63 66 67 69 71 68"] = 1
	maps["49 52 49 46 43 41"] = 1
	maps["22 27 28 31 32 33"] = 1
	maps["51 55 58 59 60 63 66"] = 1
	maps["56 55 52 51 50 48 44 45"] = 1
	maps["68 64 61 58 55 52 50"] = 1
	maps["15 12 10 7 5"] = 1
	maps["61 62 65 66 69 71 73 76"] = 1
	maps["75 76 77 80 82 85 87 90"] = 1
	maps["83 82 81 78 76"] = 1
	maps["52 50 48 46 43 42"] = 1
	maps["34 32 31 30 29 27"] = 1
	maps["53 55 57 58 61 64 65 68"] = 1
	maps["23 22 19 18 16 15"] = 1
	maps["36 35 33 30 29 26 23 21"] = 1
	maps["43 41 38 35 32 31 28"] = 1
	maps["96 94 91 89 88 86 84 83"] = 1
	maps["94 93 92 90 87 85 84 81"] = 1
	maps["83 86 89 90 93"] = 1
	maps["42 45 48 51 52 53"] = 1
	maps["3 4 6 7 10"] = 1
	maps["83 81 80 78 77 74 71 69"] = 1
	maps["69 66 63 61 58"] = 1
	maps["17 16 14 12 9"] = 1
	maps["74 73 70 68 65 63"] = 1
	maps["41 44 46 47 48 51"] = 1
	maps["81 79 78 77 74"] = 1
	maps["11 14 16 17 18 21 23"] = 1
	maps["79 82 83 85 86 89 92 95"] = 1
	maps["88 87 86 83 82 79 76"] = 1
	maps["28 31 33 36 37 38 40"] = 1
	maps["55 56 59 60 61 63 65"] = 1
	maps["57 60 62 64 67 70 73"] = 1
	maps["58 60 63 66 69 70 73 74"] = 1
	maps["24 22 21 20 18"] = 1
	maps["76 79 82 85 87 90 93"] = 1
	maps["29 27 25 23 22"] = 1
	maps["31 32 33 35 37 39 40"] = 1
	maps["14 12 11 8 7 5 4"] = 1
	maps["80 77 74 71 68 67 66 63"] = 1
	maps["57 59 60 63 65 66 69"] = 1
	maps["96 93 92 90 87 85"] = 1
	maps["76 77 80 82 83 86"] = 1
	maps["20 23 25 27 30 32"] = 1
	maps["26 29 32 34 37 38 41"] = 1
	maps["31 30 29 27 24"] = 1
	maps["31 34 35 36 37 38 41"] = 1
	maps["78 77 75 72 71 69 66"] = 1
	maps["49 47 45 42 40 37 34"] = 1
	maps["94 93 92 91 88"] = 1
	maps["81 78 76 75 74 72 71"] = 1
	maps["7 10 13 16 17 20 23"] = 1
	maps["63 64 66 67 68"] = 1
	maps["9 7 5 4 3"] = 1
	maps["75 73 71 68 66 64 61"] = 1
	maps["64 67 68 69 71 74"] = 1
	maps["60 62 64 67 70 71 73"] = 1
	maps["63 64 67 70 73 75"] = 1
	maps["95 93 92 89 88"] = 1
	maps["40 39 38 36 35"] = 1
	maps["30 28 26 24 23"] = 1
	maps["26 24 21 20 18 17 15"] = 1
	maps["90 91 92 93 96 97"] = 1
	maps["4 5 6 8 10 12 14"] = 1
	maps["36 33 32 30 28"] = 1
	maps["67 66 64 62 60 58 56"] = 1
	maps["52 51 49 48 46 45"] = 1
	maps["16 17 18 21 23 24 27 28"] = 1
	maps["26 28 31 33 35 37 39 40"] = 1
	maps["79 81 82 85 88"] = 1
	maps["36 39 41 44 46"] = 1
	maps["43 45 46 48 51 52 53"] = 1
	maps["49 46 45 43 41 39"] = 1
	maps["24 27 29 32 34 36"] = 1
	maps["55 54 52 49 48"] = 1
	maps["69 72 75 77 80 81 82"] = 1
	maps["13 12 10 7 6"] = 1
	maps["9 10 12 15 16 18 20"] = 1
	maps["76 77 78 80 81 82 84"] = 1
	maps["68 71 72 73 74"] = 1
	maps["26 24 23 22 20 18"] = 1
	maps["38 37 34 33 30 29"] = 1
	maps["53 52 49 48 47 45 42"] = 1
	maps["32 30 27 24 22 20 18"] = 1
	maps["34 35 37 38 41 44 46"] = 1
	maps["5 6 8 10 13 16 17 19"] = 1
	maps["40 38 37 34 31 28 25 23"] = 1
	maps["4 6 7 9 12 15 18"] = 1
	maps["9 11 13 14 16 19 22 23"] = 1
	maps["69 67 66 65 63 62 59 57"] = 1
	maps["81 84 85 87 90 91 94 97"] = 1
	maps["54 56 59 62 65 67 68 70"] = 1
	maps["53 55 58 61 64"] = 1
	maps["65 66 69 71 74 76 77"] = 1
	maps["93 92 90 88 87 85 83 81"] = 1
	maps["83 80 77 75 74"] = 1
	maps["77 76 75 74 73"] = 1
	maps["90 88 85 84 82 80 77 74"] = 1
	maps["58 60 61 62 65"] = 1
	maps["2 5 7 8 9 10 12"] = 1
	maps["46 45 44 41 40"] = 1
	maps["85 83 82 80 79 77 75 72"] = 1
	maps["99 97 95 93 92 89"] = 1
	maps["67 69 72 75 78"] = 1
	maps["79 77 76 74 72"] = 1
	maps["43 45 48 51 54 57 60"] = 1
	maps["98 95 92 91 90 88"] = 1
	maps["41 44 47 49 52 55 56 57"] = 1
	maps["68 66 64 62 61 59 56 55"] = 1
	maps["66 69 72 74 76 78 80 81"] = 1
	maps["72 75 76 79 81 84 86"] = 1
	maps["99 97 94 91 89"] = 1
	maps["52 55 58 59 61 63"] = 1
	maps["80 81 84 86 88 91 93"] = 1
	maps["8 9 10 12 15 17 18"] = 1
	maps["90 91 93 96 99"] = 1
	maps["30 31 34 36 38"] = 1
	maps["90 88 85 82 79 76 74 72"] = 1
	maps["82 84 87 88 89 92"] = 1
	maps["12 9 7 6 3"] = 1
	maps["83 84 87 89 91 94 97"] = 1
	maps["33 36 39 41 42 45 48"] = 1
	maps["7 8 11 13 14 15"] = 1
	maps["48 49 50 51 54 56"] = 1
	maps["92 89 86 83 80 78"] = 1
	maps["60 62 63 66 67 70 71"] = 1
	maps["25 24 21 19 16 13"] = 1
	maps["29 26 25 23 21"] = 1
	maps["71 68 67 64 63"] = 1
	maps["83 85 88 89 90 91 93 95"] = 1
	maps["2 4 7 10 12 13"] = 1
	maps["54 53 50 47 45"] = 1
	maps["8 10 12 13 15 18 19 20"] = 1
	maps["83 82 81 80 77 76 74 71"] = 1
	maps["93 90 88 87 85 84 81 79"] = 1
	maps["76 77 80 81 82 85 86"] = 1
	maps["42 39 37 36 34 32 29 26"] = 1
	maps["48 47 46 44 41 39 36"] = 1
	maps["39 38 37 36 33 32 30"] = 1
	maps["12 15 17 19 20 21 22"] = 1
	maps["30 29 28 27 24 22"] = 1
	maps["37 40 42 43 46 48 50"] = 1
	maps["35 38 39 42 45 48 50"] = 1
	maps["55 58 59 60 62"] = 1
	maps["7 8 11 13 15 16 18"] = 1
	maps["68 67 66 64 61 60"] = 1
	maps["77 79 81 83 84 87 88"] = 1
	maps["60 59 56 54 52 51 50 49"] = 1
	maps["14 16 17 20 22 25 28 30"] = 1
	maps["60 63 66 68 69 70 72"] = 1
	maps["19 16 14 11 9 6 5 3"] = 1
	maps["57 58 61 64 65 66 68"] = 1
	maps["44 42 41 38 37"] = 1
	maps["72 74 75 78 79 82 84"] = 1
	maps["51 54 56 59 60"] = 1
	maps["69 67 66 64 62"] = 1
	maps["33 30 27 25 24"] = 1
	maps["56 54 53 52 50 48"] = 1
	maps["76 78 81 84 86 89 90 93"] = 1
	maps["77 74 71 68 66 64 61 59"] = 1
	maps["41 39 36 35 32 31"] = 1
	maps["75 73 71 69 68 67 64"] = 1
	maps["75 76 78 81 82 83"] = 1
	maps["78 81 83 86 87 88 91 94"] = 1
	maps["33 30 27 25 24 21 19 16"] = 1
	maps["46 47 48 50 51 54 56 59"] = 1
	maps["76 78 79 82 84 87"] = 1
	maps["71 73 75 77 78 79 80"] = 1
	maps["52 54 56 59 61 63 65"] = 1
	maps["7 10 12 13 15 18"] = 1
	maps["3 6 8 10 12 13"] = 1
	maps["78 80 82 84 86"] = 1
	maps["58 55 53 50 47 45"] = 1
	maps["88 85 82 79 78 77 74"] = 1
	maps["8 7 6 5 3 2"] = 1
	maps["38 36 33 30 29 28"] = 1
	maps["59 62 63 66 68 70"] = 1
	maps["83 86 88 89 91 93 96 99"] = 1
	maps["19 21 23 26 28 31"] = 1
	maps["14 16 19 22 25 27"] = 1
	maps["91 90 88 87 86 85 84 83"] = 1
	maps["5 8 11 14 17 18 20"] = 1
	maps["87 88 91 92 94 96 97"] = 1
	maps["4 5 6 8 9"] = 1
	maps["32 30 27 24 23 20"] = 1
	maps["56 55 54 52 49 48"] = 1
	maps["31 30 28 27 24 22 19"] = 1
	maps["79 82 85 86 88 89 90 91"] = 1
	maps["92 89 87 84 83 80"] = 1
	maps["75 72 71 69 67"] = 1
	maps["54 51 49 46 45 44 41 38"] = 1
	maps["38 36 35 34 31"] = 1
	maps["68 66 63 62 59 58"] = 1
	maps["90 89 87 86 84 81 78"] = 1
	maps["51 52 54 57 59 60 61 62"] = 1
	maps["18 19 20 21 23 25"] = 1
	maps["52 50 47 45 44"] = 1
	maps["71 73 74 76 79 82 85"] = 1
	maps["19 21 23 25 26"] = 1
	maps["47 46 43 41 40 38 36 33"] = 1
	maps["32 35 37 40 42 43 45"] = 1
	maps["57 54 52 50 47 46 44 42"] = 1
	maps["33 35 37 40 43 44"] = 1
	maps["38 35 33 32 29 26 24 23"] = 1
	maps["63 61 60 57 56"] = 1
	maps["90 87 84 83 81 80 79 76"] = 1
	maps["11 13 14 15 18 19 21"] = 1
	maps["59 56 53 51 48 47 44"] = 1
	maps["54 56 57 60 61 63"] = 1
	maps["26 27 29 31 34 35 38"] = 1
	maps["64 63 61 59 58 55 54"] = 1
	maps["17 15 14 12 10 7"] = 1
	maps["85 86 88 91 94"] = 1
	maps["23 26 29 30 32 33 36 37"] = 1
	maps["31 32 33 34 36 38 39"] = 1
	maps["88 89 92 94 97"] = 1
	maps["74 72 70 67 64 61"] = 1
	maps["18 17 14 11 10 8 5 2"] = 1
	maps["66 63 60 57 54 51 48"] = 1
	maps["49 52 55 56 58 59 62"] = 1
	maps["55 52 49 48 47 46 45"] = 1
	maps["59 57 56 55 54 53 50"] = 1
	maps["25 26 28 29 31 32 33"] = 1
	maps["52 50 48 47 46 45 44 41"] = 1
	maps["48 47 46 44 41"] = 1
	maps["16 14 13 12 9"] = 1
	maps["18 19 20 21 24"] = 1
	maps["34 37 40 43 44"] = 1
	maps["46 49 51 54 56"] = 1
	maps["81 79 78 77 75 72 69 68"] = 1
	maps["26 24 22 19 17 16 13"] = 1
	maps["19 21 24 26 27 30"] = 1
	maps["13 14 16 18 19 22 24 26"] = 1
	maps["41 43 45 46 47 48"] = 1
	maps["92 89 87 85 84"] = 1
	maps["34 35 38 39 40 41 44"] = 1
	maps["96 95 93 90 89"] = 1
	maps["66 68 69 70 71 72"] = 1
	maps["82 79 78 77 74"] = 1
	maps["40 38 37 34 33 31 28"] = 1
	maps["71 72 74 76 78 80 81 83"] = 1
	maps["37 36 34 32 31 28 27 25"] = 1
	maps["40 42 44 47 49 51 52"] = 1
	maps["84 81 78 77 76 75"] = 1
	maps["22 23 25 26 29 31 32 35"] = 1
	maps["35 34 33 31 29 28 25"] = 1
	maps["34 36 38 41 44 46 47 49"] = 1
	maps["51 48 46 43 42 41 38 36"] = 1
	maps["1 2 4 5 7"] = 1
	maps["70 68 65 64 62"] = 1
	maps["34 31 30 29 27 26"] = 1
	maps["35 34 31 29 26"] = 1
	maps["33 35 38 39 42 45"] = 1
	maps["65 64 61 59 56 53"] = 1
	maps["65 63 60 57 55"] = 1
	maps["38 36 33 32 29"] = 1
	maps["59 61 64 65 66"] = 1
	maps["93 92 89 86 84 83 82"] = 1
	maps["14 13 12 9 7"] = 1
	maps["63 66 69 70 72 75 76 79"] = 1
	maps["31 32 33 36 37 39"] = 1
	maps["37 39 40 42 45 47"] = 1
	maps["20 22 23 25 26"] = 1
	maps["55 52 49 46 45 42"] = 1
	maps["19 21 23 24 27 30 33"] = 1
	maps["2 4 6 7 9 12 15 16"] = 1
	maps["76 79 80 82 83 84 87"] = 1
	maps["49 51 52 53 55 57"] = 1
	maps["54 55 57 59 61"] = 1
	maps["3 5 6 9 12 14 15"] = 1
	maps["95 92 91 88 87 84 83 80"] = 1
	maps["54 56 57 58 61 62 63"] = 1
	maps["1 3 5 6 8 10"] = 1
	maps["76 79 80 82 85"] = 1
	maps["21 23 24 27 28 31 33"] = 1
	maps["52 49 46 43 41 40 38"] = 1
	maps["69 71 74 76 77"] = 1
	maps["57 55 52 49 48"] = 1
	maps["22 19 16 14 11"] = 1
	maps["67 68 70 72 73 74 75 76"] = 1
	maps["17 14 13 10 7 6 3 2"] = 1
	maps["76 73 72 71 69 67 66"] = 1
	maps["70 72 73 75 78 79"] = 1
	maps["56 53 50 47 44 43 42 39"] = 1
	maps["72 70 68 66 64 61"] = 1
	maps["80 81 84 86 87 90"] = 1
	maps["35 32 29 28 26 25"] = 1
	maps["30 33 35 36 39 40 41 43"] = 1
	maps["62 59 57 56 53"] = 1
	maps["51 48 46 44 43"] = 1
	maps["32 31 29 28 25 23 21 19"] = 1
	maps["90 89 87 85 84 81 80 79"] = 1
	maps["66 64 62 59 56 55 54 53"] = 1
	maps["70 71 72 73 74 76 79 81"] = 1
	maps["84 86 88 89 92 94 95 98"] = 1
	maps["61 60 58 56 54"] = 1
	maps["79 81 83 86 87 89 92"] = 1
	maps["23 20 17 16 15 14 11 8"] = 1
	maps["39 38 37 34 33 32"] = 1
	maps["14 13 10 8 7 6"] = 1
	maps["29 31 34 37 38 41 43"] = 1
	maps["22 24 26 28 29 30"] = 1
	maps["31 28 27 25 24"] = 1
	maps["51 50 49 48 46 44 43"] = 1
	maps["69 70 71 74 75 77"] = 1
	return maps
}

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
	var lines []string = readfile("/Users/jkelly/Projects/adventforcode/2024/day-02-part-02/large-input.txt")
	//missing := fixmap()
	total := 0
	for i := 0; i < len(lines); i++ {
		split_text := strings.Split(lines[i], " ")
		split_ints := stringArrayToIntArray(split_text)
		if valid(split_ints) {
			total++
		} else {
			for j := 0; j < len(split_ints); j++ {
				fmt.Println(split_ints, j)
				if valid(removeIndex(split_ints, j)) {
					total++
					break
				}
			}

		}
	}
	fmt.Println("-----")
	fmt.Println(total)
}

func removeIndex(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}
func diffInValid(first, second int) bool {
	var diff = abs(first - second)
	return !((diff > 0) && (diff < 4))
}
func valid(values []int) bool {
	multiplier := 1
	if values[0] < values[1] {
		multiplier = -1
	}
	current := values[0]
	i := 1
	for i < len(values) {
		fmt.Println("Compare", current, values[i])
		if ((current-values[i])*multiplier) < 0 || diffInValid(current, values[i]) {
			return false
		}
		current = values[i]
		i++
	}
	return true
}
