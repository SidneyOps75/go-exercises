// package main

// import (
// 	"fmt"

// 	"os"
// )

// func WriteFile() {
// 	file, err := os.Create("test.txt")
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	}
// 	defer file.Close()
// 	_, err = file.WriteString("Heloo")
// 	if err != nil {
// 		fmt.Println("Error", err)
// 	}
// }
// func ReadFile() {
// 	data, err := os.ReadFile("test.txt")
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	}
// 	fmt.Println(string(data))
// }
// func main() {
// 	WriteFile()
// 	ReadFile()

// }
