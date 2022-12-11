package days

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func Day01() string {
	totalCalories, err := totalCalories()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("TotalCalories for the top three elfs are: %v\n", totalCalories)
	return totalCalories
}

// totalCalories started out looking for the elf with the most calories worth of startfruit. Instead of refactoring
// I just updated this function to account for part two and return the top three. Need to break into separate funcs
// TODO: Refactor, cleanup, and tests
func totalCalories() (totalCalories string, err error) {
	file, err := os.Open("internal/testdata/days/day01.txt")
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	defer file.Close()

	// TODO: maybe create a map and increment the key at the linebreak (ln 39) or make some types/structs
	i := 0
	overallTotal := 0
	tmp := []int{}
	content := [][]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		calorie, err := strconv.Atoi(line)
		if err != nil {
			// log.Printf("Unable to parse: %v. Assuming end of elf's starfruit.\n", line)
			content = append(content, tmp)
			tmp = []int{}
			continue
		}
		tmp = append(tmp, calorie)
		overallTotal += calorie
		// log.Printf("line[%v]: %v\n", i, line)
		i++
	}
	log.Printf("overallTotal: %v\n", overallTotal)

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	// log.Printf("content: %v\n", content)

	first, second, third := 0, 0, 0
	for _, group := range content {
		elfTotal := 0
		for _, fruit := range group {
			elfTotal += fruit
		}
		if elfTotal > first {
			if first > second {
				if second > third {
					third = second
				}
				second = first
			}
			first = elfTotal
			continue
		}
		if elfTotal > second {
			if second > third {
				third = second
			}
			second = elfTotal
			continue
		}
		if elfTotal > third {
			third = elfTotal
			continue
		}
	}
	topThree := first + second + third
	log.Printf("\n\tfirst: %v\n\tsecond: %v\n\tthird: %v\n\ttopThree: %v\n", first, second, third, topThree)
	return strconv.Itoa(topThree), nil
}
