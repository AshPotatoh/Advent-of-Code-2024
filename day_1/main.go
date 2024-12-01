package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

// I imagine this can be neater, but it's my first time using Go lol
// maybe returning multi-dimensional arrays intead
func read_data() ([]int, []int, []string, []string) {
	var list_one = []int{}
	var list_two = []int{}
	var string_one = []string{}
	var string_two = []string{}
	filedata, err := os.Open("input.txt")
	defer filedata.Close()
	r := bufio.NewReader(filedata)
	if err != nil {
		fmt.Println("Error reading file")
	}
	for {
		line, _, err := r.ReadLine()
		if len(line) > 0 {
			result := strings.Split(string(line), "   ")
			string_one = append(string_one, result[0])
			string_two = append(string_two, result[1])

			//Converting to Int to calculate distance later
			int_result_one, _ := strconv.Atoi(result[0])
			int_result_two, _ := strconv.Atoi(result[1])
			list_one = append(list_one, int_result_one)
			list_two = append(list_two, int_result_two)
			slices.Sort(list_one)
			slices.Sort(list_two)
		}
		if err != nil {
			break
		}
	}
	return list_one, list_two, string_one, string_two
}

func calculate_distance(o []int, t []int) int {
	var distance int
	for i := 0; i < len(o); i++ {
		if o[i] < t[i] {
			distance = distance + (t[i] - o[i])
			fmt.Println("Total Distance: ", distance)
		} else {
			distance = distance + (o[i] - t[i])
			fmt.Println("Total Distance: ", distance)
		}
	}
	return distance
}

func calculate_similarity(o []string, t []string) int {
	var similarity int
	joined_t := strings.Join(t, " ")
	for i := 0; i < len(o); i++ {

		// We can use Regex to quickly find the copies of a string in the other Slice.
		// Then Multiply the queary by the length of the result
		r, _ := regexp.Compile(o[i])
		regex_amount := r.FindAllString(joined_t, -1)
		len_regex := len(regex_amount)

		if len_regex > 0 {
			int_o, _ := strconv.Atoi(o[i])
			similarity = similarity + (int_o * len_regex)
		}

	}
	return similarity
}

func main() {
	list_one, list_two, string_one, string_two := read_data()
	distance := calculate_distance(list_one, list_two)
	similarity := calculate_similarity(string_one, string_two)
	fmt.Println("Distance: ", distance, "\nSimilarity: ", similarity)
}
