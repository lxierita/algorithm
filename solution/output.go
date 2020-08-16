package solution

import (
	"fmt"
	"os"
	"regexp"
)

//CreateText writes the results to a textfile
func CreateText() (err error) {
	file, err := os.Create("rita_result.txt")
	if err != nil {
		fmt.Printf("Could not create text file: %s\n", err)
	}

	var result string
	for _, set := range Solution() {
		result = fmt.Sprintf("%v\n%v", removeBrackets(set), result)
	}
	defer file.Close()
	_, err = file.WriteString(result)
	if err != nil {
		fmt.Printf("Could not write to file: %s\n", err)
	}
	return err
}

func removeBrackets(set []int) string {
	str := fmt.Sprintf("%v", set)
	re := regexp.MustCompile(`[\[\]]+`)
	str = re.ReplaceAllString(str, "")
	return str
}
