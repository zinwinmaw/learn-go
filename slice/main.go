package main

import (
	"fmt"
	"regexp"
	"time"
)

func FavoriteCards() []int {
	return []int{2, 6, 9}
}

func Shuffle() []string {
	animals := []string{"ant", "beaver", "cat", "dog", "elephant", "fox", "giraffe", "hedghog"}

	// rand.Shuffle(len(animals), func(i, j int) {
	// 	animals[i], animals[j] = animals[j], animals[i]
	// })
	return animals
}

func main() {
	// cards := FavoriteCards()
	// fmt.Println(cards)

	// fmt.Println(Shuffle())
	// fmt.Println(Schedule("7/25/2019 13:45:00"))

	// fmt.Println(AnniversaryDate())

	// fmt.Println(Units())

	a := 42
	p := &a // p is a pointer to a

	fmt.Println(a)  // Output: 42
	fmt.Println(p)  // Output: memory address of a
	fmt.Println(*p) // Output: 42, dereferencing the pointer

	// myString := "❗hello❗"
	// stringLength := len(myString)                     // count by byte of UTF-8 encoding
	// numberOfRunes := utf8.RuneCountInString(myString) // count by unicode code point

	// fmt.Printf("myString - Length: %d - Runes: %d\n", stringLength, numberOfRunes)

	// lines := []string{
	// 	`[INF] passWord`, // contains 'password' but not surrounded by quotation marks
	// 	`"passWord"`,     // count this one
	// 	`[INF] User saw error message "Unexpected Error" on page load.`,          // does not contain 'password'
	// 	`[INF] The message "Please reset your password" was ignored by the user`, // count this one
	// }
	// fmt.Println(CountQuotedPasswords(lines))

	lines := []string{
		"[WRN] User James123 has exceeded storage space.",
		"[WRN] Host down. User   Michelle4 lost connection.",
		"[INF] Users can login again after 23:00.",
		"[DBG] We need to check that user names are at least 6 chars long.",
	}
	fmt.Println(TagWithUserName(lines))
}

// func TagWithUserName(lines []string) []string {

// 	re := regexp.MustCompile(`User\s+\b([a-zA-Z0-9]+)\b\s+`)

// 	for i, line := range lines {
// 		if re.MatchString(line) {
// 			words := re.FindStringSubmatch(line)
// 			lines[i] = "[USR] " + words[1] + " " + line
// 		}
// 	}
// 	return lines
// }

func TagWithUserName(lines []string) []string {

	re := regexp.MustCompile(`User\s+\b([a-zA-Z0-9]+)\b\s+`)
	var tags []string

	for i, line := range lines {
		if re.MatchString(line) {
			words := re.FindStringSubmatch(line)
			tags[i] = "[USR] " + words[1] + " " + line
		} else {
			tags[i] = line
		}
	}
	return tags
}

func CountQuotedPasswords(lines []string) int {
	// re := regexp.MustCompile(`"(?i)password"`)
	re := regexp.MustCompile(`"(?i:.*password.*)"`)
	count := 0
	for _, line := range lines {
		if re.MatchString(line) {
			count++
		}
	}
	return count
}

func Schedule(date string) time.Time {

	inputFormat := "1/02/2006 15:04:05"
	parseTime, err := time.Parse(inputFormat, date)
	if err != nil {
		panic(err)
	}
	return parseTime
}

func AnniversaryDate() time.Time {
	parseTime, _ := time.Parse("2006-01-2", fmt.Sprintf("%d-09-15", time.Now().Year()))
	return parseTime
}

func Units() map[string]int {
	units := map[string]int{}
	units["quarter_of_a_dozen"] = 3
	units["half_of_a_dozen"] = 6
	units["dozen"] = 12
	units["small_gross"] = 120
	units["gross"] = 144
	units["great_gross"] = 1728

	return units
}
