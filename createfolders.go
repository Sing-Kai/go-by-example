package main

/*
script that generates folders
*/
import (
	"bufio"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	data, err := os.Open("foldernames.txt")

	check(err)

	defer data.Close()

	scanner := bufio.NewScanner(data)

	names := make([]string, 0)

	for scanner.Scan() {
		name := scanner.Text()
		names = append(names, name)

	}

	newNames := createPrefixNames(names)
	createFolder(newNames)
	createGoFile(newNames)
}

func createGoFile(names []string) {
	path := "GoByExample/"

	for i := 0; i < len(names); i++ {
		newPath := path + names[i] + "/" + names[i] + ".go"

		if _, err := os.Stat(newPath); os.IsNotExist(err) {
			_, errDir := os.OpenFile(newPath, os.O_RDONLY|os.O_CREATE, 0666)

			if errDir != nil {
				panic(errDir)
			}
		}

	}
}

func delete(names []string) {

	path := "GoByExample/"

	for i := 0; i < len(names); i++ {
		newPath := path + names[i] + "/" + names[i] + ".go"

		if _, err := os.Stat(newPath); !os.IsNotExist(err) {
			errDir := os.Remove(newPath)

			if errDir != nil {
				panic(errDir)
			}
		}

	}

}

func createFolder(names []string) {
	path := "GoByExample/"

	for i := 0; i < len(names); i++ {
		newPath := path + names[i]

		if _, err := os.Stat(newPath); os.IsNotExist(err) {
			errDir := os.MkdirAll(newPath, 0755)

			if errDir != nil {
				panic(errDir)
			}
		}

	}

}

func createPrefixNames(folderNames []string) []string {

	names := make([]string, 0)

	for i := 0; i < len(folderNames); i++ {
		prefix := createPrefix(i)
		newFolderName := prefix + folderNames[i]
		names = append(names, newFolderName)
	}

	return names
}

func createPrefix(val int) string {

	prefix := ""

	val += 1
	max := 3
	strNum := strconv.Itoa(val)
	strLen := len(strNum)

	for i := 0; strLen < max; i++ {
		prefix += "0"
		strLen = len(prefix + strNum)
	}

	prefix = prefix + strNum

	return prefix
}
