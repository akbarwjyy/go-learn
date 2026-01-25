package main

import (
	"fmt"
	"strings"
)

func main() {
	// 1. Basic for loop (seperti C/Java)
	fmt.Println("1. Basic For Loop:")
	for i := 1; i <= 5; i++ {
		fmt.Printf("Iterasi ke-%d\n", i)
	}
	fmt.Println()

	// 2. For dengan condition saja (seperti while)
	fmt.Println("2. For dengan Condition (seperti while):")
	count := 1
	for count <= 5 {
		fmt.Printf("Count: %d\n", count)
		count++
	}
	fmt.Println()

	// 3. Infinite loop (dengan break)
	fmt.Println("3. Infinite Loop dengan Break:")
	num := 0
	for {
		num++
		fmt.Printf("Num: %d\n", num)
		if num >= 5 {
			break
		}
	}
	fmt.Println()

	// 4. For range dengan array
	fmt.Println("4. For Range dengan Array:")
	numbers := [5]int{10, 20, 30, 40, 50}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
	fmt.Println()

	// 5. For range dengan slice
	fmt.Println("5. For Range dengan Slice:")
	fruits := []string{"Apple", "Banana", "Cherry", "Durian"}
	for i, fruit := range fruits {
		fmt.Printf("%d. %s\n", i+1, fruit)
	}
	fmt.Println()

	// 6. For range - hanya index
	fmt.Println("6. For Range - Hanya Index:")
	colors := []string{"Red", "Green", "Blue"}
	for i := range colors {
		fmt.Printf("Index: %d\n", i)
	}
	fmt.Println()

	// 7. For range - hanya value
	fmt.Println("7. For Range - Hanya Value:")
	for _, color := range colors {
		fmt.Printf("Color: %s\n", color)
	}
	fmt.Println()

	// 8. For range dengan map
	fmt.Println("8. For Range dengan Map:")
	ages := map[string]int{
		"Alice":   25,
		"Bob":     30,
		"Charlie": 35,
	}
	for name, age := range ages {
		fmt.Printf("%s berumur %d tahun\n", name, age)
	}
	fmt.Println()

	// 9. For range dengan string
	fmt.Println("9. For Range dengan String:")
	text := "Hello"
	for index, char := range text {
		fmt.Printf("Index: %d, Char: %c, Unicode: %U\n", index, char, char)
	}
	fmt.Println()

	// 10. Continue statement
	fmt.Println("10. Continue Statement:")
	fmt.Println("Angka ganjil dari 1-10:")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue // Skip bilangan genap
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println("\n")

	// 11. Break statement
	fmt.Println("11. Break Statement:")
	fmt.Println("Cari angka 7:")
	for i := 1; i <= 10; i++ {
		if i == 7 {
			fmt.Printf("Ketemu! Angka %d\n", i)
			break
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// 12. Nested loops
	fmt.Println("12. Nested Loops - Tabel Perkalian:")
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			fmt.Printf("%3d ", i*j)
		}
		fmt.Println()
	}
	fmt.Println()

	// 13. Loop dengan decrement
	fmt.Println("13. Loop dengan Decrement:")
	fmt.Println("Countdown:")
	for i := 5; i >= 1; i-- {
		fmt.Printf("%d... ", i)
	}
	fmt.Println("Blast off! 🚀\n")

	// 14. Loop dengan increment lebih dari 1
	fmt.Println("14. Loop dengan Increment > 1:")
	for i := 0; i <= 20; i += 5 {
		fmt.Printf("%d ", i)
	}
	fmt.Println("\n")

	// 15. Loop dengan label dan break
	fmt.Println("15. Loop dengan Label:")
outer:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("i=%d, j=%d\n", i, j)
			if i == 2 && j == 2 {
				fmt.Println("Breaking outer loop!")
				break outer
			}
		}
	}
	fmt.Println()

	// 16. Real world example - Sum of numbers
	fmt.Println("16. Real World - Sum of Numbers:")
	nums := []int{5, 10, 15, 20, 25}
	sum := 0
	for _, n := range nums {
		sum += n
	}
	fmt.Printf("Numbers: %v\n", nums)
	fmt.Printf("Total: %d\n\n", sum)

	// 17. Real world example - Find max value
	fmt.Println("17. Real World - Find Max Value:")
	scores := []int{85, 92, 78, 95, 88}
	maxScore := scores[0]
	for _, score := range scores {
		if score > maxScore {
			maxScore = score
		}
	}
	fmt.Printf("Scores: %v\n", scores)
	fmt.Printf("Highest score: %d\n\n", maxScore)

	// 18. Real world example - Filter data
	fmt.Println("18. Real World - Filter Data:")
	allNumbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var evenNumbers []int
	for _, num := range allNumbers {
		if num%2 == 0 {
			evenNumbers = append(evenNumbers, num)
		}
	}
	fmt.Printf("All numbers: %v\n", allNumbers)
	fmt.Printf("Even numbers: %v\n\n", evenNumbers)

	// 19. Real world example - String manipulation
	fmt.Println("19. Real World - String Manipulation:")
	words := []string{"go", "is", "awesome"}
	var upperWords []string
	for _, word := range words {
		upperWords = append(upperWords, strings.ToUpper(word))
	}
	fmt.Printf("Original: %v\n", words)
	fmt.Printf("Uppercase: %v\n\n", upperWords)

	// 20. Real world example - FizzBuzz
	fmt.Println("20. Real World - FizzBuzz:")
	for i := 1; i <= 15; i++ {
		if i%15 == 0 {
			fmt.Print("FizzBuzz ")
		} else if i%3 == 0 {
			fmt.Print("Fizz ")
		} else if i%5 == 0 {
			fmt.Print("Buzz ")
		} else {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println("\n")

	// 21. Loop dengan multiple variables
	fmt.Println("21. Loop dengan Multiple Variables:")
	for i, j := 0, 10; i < j; i, j = i+1, j-1 {
		fmt.Printf("i=%d, j=%d\n", i, j)
	}
	fmt.Println()

	// 22. Print pattern - Triangle
	fmt.Println("22. Pattern - Triangle:")
	for i := 1; i <= 5; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
	fmt.Println()

	// 23. Reverse iteration
	fmt.Println("23. Reverse Iteration:")
	animals := []string{"Cat", "Dog", "Elephant", "Fish"}
	fmt.Println("Forward:")
	for _, animal := range animals {
		fmt.Printf("%s ", animal)
	}
	fmt.Println("\nReverse:")
	for i := len(animals) - 1; i >= 0; i-- {
		fmt.Printf("%s ", animals[i])
	}
	fmt.Println("\n")

	// 24. Loop dengan skip index
	fmt.Println("24. Loop dengan Skip Index:")
	items := []string{"A", "B", "C", "D", "E"}
	fmt.Println("Skip index 2:")
	for i, item := range items {
		if i == 2 {
			continue
		}
		fmt.Printf("%d: %s\n", i, item)
	}
	fmt.Println()
}
