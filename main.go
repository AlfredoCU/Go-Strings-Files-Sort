package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	// Declaration.
	var (
		length int
		str	   string
		data1  []string
		data2  []string
	)

	// For strings.
	scanner := bufio.NewScanner(os.Stdin)

	// Added strings.
	fmt.Println("******* REGISTER STRINGS *******")
	fmt.Print("\nHow many registrations do you want to make? ")
	fmt.Scan(&length)

	for i := 0; i < length; i++ {
		fmt.Print("Add String [", i + 1, "]: ")
		scanner.Scan()
		str = scanner.Text()
		data1 = append(data1, str)
		data2 = append(data2, str)
	}

	// Sort strings.
	sort.Strings(data1)
	sort.Sort(sort.Reverse(sort.StringSlice(data2)))

	// Create and write files.
	createFile("ascendant.txt", data1)
	createFile("descendant.txt", data2)

	// Open and read files.
	openFile("ascendant.txt")
	openFile("descendant.txt")
}

// Create file function.
func createFile(name string, data []string) {
	// Create file.
	file, err := os.Create(name)
	errorFile(err)
	defer file.Close()

	// Write file.
	writeFile(file, data)
}

// Write file function.
func writeFile(file *os.File, data []string) {
	file.WriteString(strings.Join(data, "\n"))
}

// Open file function.
func openFile(name string) {
	// Open file.
	file, err := os.Open(name)
	errorFile(err)
	defer file.Close()

	// Read file.
	ReadFile(file)
}


// Read and information of the file function.
func ReadFile(file *os.File) {
	// File information.
	stat, err := file.Stat()
	errorFile(err)

	total := stat.Size() // Total of the bytes.
	bytes := make([]byte, total) // Array of the bytes with length of total.

	// Read file.
	count, err := file.Read(bytes)
	errorFile(err)

	str := string(bytes) // Convert bytes to string.
	fmt.Println(str) // Word.
	fmt.Println("Count:", count) // Total = Count.
}

// Error file function.
func errorFile(e error) {
	if e != nil {
		fmt.Println("\nERROR:", e)
		return
	}
}