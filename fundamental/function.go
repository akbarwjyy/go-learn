package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("=== Belajar Parameter & Return Value di Go ===\n")

	// 1. Function dengan single parameter
	fmt.Println("1. Single Parameter:")
	greet("Budi")
	greet("Ani")
	fmt.Println()

	// 2. Multiple parameters
	fmt.Println("2. Multiple Parameters:")
	sum := add(10, 20)
	fmt.Printf("10 + 20 = %d\n", sum)

	result := multiply(5, 6)
	fmt.Printf("5 x 6 = %d\n\n", result)

	// 3. Parameters dengan tipe yang sama
	fmt.Println("3. Parameters dengan Tipe yang Sama:")
	total := add3Numbers(10, 20, 30)
	fmt.Printf("10 + 20 + 30 = %d\n\n", total)

	// 4. Multiple return values
	fmt.Println("4. Multiple Return Values:")
	quot, rem := divide(17, 5)
	fmt.Printf("17 / 5 = %d sisa %d\n\n", quot, rem)

	// 5. Named return values
	fmt.Println("5. Named Return Values:")
	area, perimeter := rectangleInfo(10, 5)
	fmt.Printf("Persegi panjang 10x5:\n")
	fmt.Printf("  Luas: %d\n", area)
	fmt.Printf("  Keliling: %d\n\n", perimeter)

	// 6. Return value dengan error handling
	fmt.Println("6. Return dengan Error Handling:")
	res1, err1 := safeDivide(10, 2)
	if err1 != nil {
		fmt.Println("Error:", err1)
	} else {
		fmt.Printf("10 / 2 = %.2f\n", res1)
	}

	res2, err2 := safeDivide(10, 0)
	if err2 != nil {
		fmt.Println("Error:", err2)
	} else {
		fmt.Printf("10 / 0 = %.2f\n", res2)
	}
	fmt.Println()

	// 7. Variadic functions
	fmt.Println("7. Variadic Functions:")
	total1 := sumAll(1, 2, 3)
	fmt.Printf("Sum of 1, 2, 3 = %d\n", total1)

	total2 := sumAll(5, 10, 15, 20, 25)
	fmt.Printf("Sum of 5, 10, 15, 20, 25 = %d\n\n", total2)

	// 8. Variadic dengan slice
	fmt.Println("8. Variadic dengan Slice:")
	nums := []int{10, 20, 30, 40}
	total3 := sumAll(nums...)
	fmt.Printf("Sum of %v = %d\n\n", nums, total3)

	// 9. Mixed parameters (normal + variadic)
	fmt.Println("9. Mixed Parameters:")
	message := formatMessage("Info", "Server started", "on port 8080", "successfully")
	fmt.Println(message)
	fmt.Println()

	// 10. Function sebagai parameter (callback)
	fmt.Println("10. Function sebagai Parameter:")
	numbers := []int{1, 2, 3, 4, 5}
	doubled := applyFunction(numbers, double)
	fmt.Printf("Original: %v\n", numbers)
	fmt.Printf("Doubled: %v\n", doubled)

	squared := applyFunction(numbers, square)
	fmt.Printf("Squared: %v\n\n", squared)

	// 11. Anonymous function
	fmt.Println("11. Anonymous Function:")
	subtract := func(a, b int) int {
		return a - b
	}
	fmt.Printf("50 - 20 = %d\n\n", subtract(50, 20))

	// 12. Closure
	fmt.Println("12. Closure:")
	counter := makeCounter()
	fmt.Println("Counter:", counter()) // 1
	fmt.Println("Counter:", counter()) // 2
	fmt.Println("Counter:", counter()) // 3
	fmt.Println()

	// 13. Pass by value
	fmt.Println("13. Pass by Value:")
	x := 10
	fmt.Printf("Before: x = %d\n", x)
	changeValue(x)
	fmt.Printf("After: x = %d (tidak berubah)\n\n", x)

	// 14. Pass by reference (pointer)
	fmt.Println("14. Pass by Reference:")
	y := 10
	fmt.Printf("Before: y = %d\n", y)
	changeValueByReference(&y)
	fmt.Printf("After: y = %d (berubah)\n\n", y)

	// 15. Slice sebagai parameter (reference)
	fmt.Println("15. Slice sebagai Parameter:")
	scores := []int{10, 20, 30}
	fmt.Printf("Before: %v\n", scores)
	modifySlice(scores)
	fmt.Printf("After: %v (berubah)\n\n", scores)

	// 16. Return function
	fmt.Println("16. Return Function:")
	addFive := makeAdder(5)
	fmt.Printf("5 + 10 = %d\n", addFive(10))
	fmt.Printf("5 + 20 = %d\n\n", addFive(20))

	// 17. Real world - Validator
	fmt.Println("17. Real World - Validator:")
	validateAndPrint("john@example.com", isValidEmail)
	validateAndPrint("invalid-email", isValidEmail)
	fmt.Println()

	// 18. Real world - Calculator
	fmt.Println("18. Real World - Calculator:")
	calc := calculator(10, 5)
	fmt.Printf("Sum: %d\n", calc.sum)
	fmt.Printf("Difference: %d\n", calc.diff)
	fmt.Printf("Product: %d\n", calc.product)
	fmt.Printf("Quotient: %.2f\n\n", calc.quotient)

	// 19. Real world - String operations
	fmt.Println("19. Real World - String Operations:")
	text := "hello world"
	upper, lower, capitalized := stringTransform(text)
	fmt.Printf("Original: %s\n", text)
	fmt.Printf("Upper: %s\n", upper)
	fmt.Printf("Lower: %s\n", lower)
	fmt.Printf("Capitalized: %s\n\n", capitalized)

	// 20. Defer with return
	fmt.Println("20. Defer with Return:")
	result1 := processWithDefer(true)
	fmt.Printf("Result: %d\n", result1)
	result2 := processWithDefer(false)
	fmt.Printf("Result: %d\n\n", result2)

	fmt.Println("=== Selesai ===")
}

// 1. Single parameter
func greet(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

// 2. Multiple parameters
func add(a int, b int) int {
	return a + b
}

func multiply(a, b int) int { // shorthand untuk tipe yang sama
	return a * b
}

// 3. Parameters dengan tipe yang sama
func add3Numbers(a, b, c int) int {
	return a + b + c
}

// 4. Multiple return values
func divide(a, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

// 5. Named return values
func rectangleInfo(length, width int) (area int, perimeter int) {
	area = length * width
	perimeter = 2 * (length + width)
	return // naked return
}

// 6. Return dengan error
func safeDivide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("tidak bisa membagi dengan nol")
	}
	return a / b, nil
}

// 7. Variadic function
func sumAll(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// 9. Mixed parameters
func formatMessage(level string, parts ...string) string {
	message := "[" + level + "] " + strings.Join(parts, " ")
	return message
}

// 10. Function sebagai parameter
func applyFunction(numbers []int, fn func(int) int) []int {
	result := make([]int, len(numbers))
	for i, num := range numbers {
		result[i] = fn(num)
	}
	return result
}

func double(n int) int {
	return n * 2
}

func square(n int) int {
	return n * n
}

// 12. Closure
func makeCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// 13. Pass by value
func changeValue(n int) {
	n = 100
	fmt.Printf("Inside function: n = %d\n", n)
}

// 14. Pass by reference
func changeValueByReference(n *int) {
	*n = 100
	fmt.Printf("Inside function: n = %d\n", *n)
}

// 15. Slice sebagai parameter
func modifySlice(s []int) {
	for i := range s {
		s[i] = s[i] * 2
	}
}

// 16. Return function
func makeAdder(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

// 17. Validator
func isValidEmail(email string) bool {
	return strings.Contains(email, "@") && strings.Contains(email, ".")
}

func validateAndPrint(value string, validator func(string) bool) {
	if validator(value) {
		fmt.Printf("✓ '%s' is valid\n", value)
	} else {
		fmt.Printf("✗ '%s' is invalid\n", value)
	}
}

// 18. Calculator struct untuk multiple returns
type CalculatorResult struct {
	sum      int
	diff     int
	product  int
	quotient float64
}

func calculator(a, b int) CalculatorResult {
	return CalculatorResult{
		sum:      a + b,
		diff:     a - b,
		product:  a * b,
		quotient: float64(a) / float64(b),
	}
}

// 19. String operations
func stringTransform(s string) (upper string, lower string, capitalized string) {
	upper = strings.ToUpper(s)
	lower = strings.ToLower(s)
	capitalized = strings.Title(s)
	return
}

// 20. Defer with return
func processWithDefer(success bool) int {
	defer fmt.Println("  Cleanup: Closing resources...")

	if success {
		fmt.Println("  Processing: Success")
		return 1
	}

	fmt.Println("  Processing: Failed")
	return 0
}
