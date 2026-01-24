package main

import "fmt"

func main() {
	// 1. Deklarasi dengan var
	fmt.Println("1. Deklarasi dengan var:")
	var name string = "Akbar"
	var age int = 20
	var isStudent bool = true
	fmt.Printf("Nama: %s, Umur: %d, Mahasiswa: %t\n\n", name, age, isStudent)

	// 2. Deklarasi tanpa nilai awal (zero values)
	fmt.Println("2. Zero Values:")
	var username string // "" (empty string)
	var count int       // 0
	var isActive bool   // false
	var price float64   // 0.0
	fmt.Printf("username: '%s', count: %d, isActive: %t, price: %.2f\n\n", username, count, isActive, price)

	// 3. Short declaration dengan :=
	fmt.Println("3. Short Declaration (:=):")
	city := "Jogja"
	population := 10_000_000
	area := 664.01
	fmt.Printf("Kota: %s, Populasi: %d, Luas: %.2f km²\n\n", city, population, area)

	// 4. Multiple variable declaration
	fmt.Println("4. Multiple Variable Declaration:")
	var x, y, z int = 1, 2, 3
	fmt.Printf("x=%d, y=%d, z=%d\n", x, y, z)

	a, b, c := "Apple", "Banana", "Cherry"
	fmt.Printf("Fruits: %s, %s, %s\n\n", a, b, c)

	// 5. Group declaration
	fmt.Println("5. Group Declaration:")
	var (
		firstName string  = "Akbar"
		lastName  string  = "Wijaya"
		height    float64 = 170.5
		weight    float64 = 53.5
	)
	fmt.Printf("%s %s - Height: %.1f cm, Weight: %.1f kg\n\n", firstName, lastName, height, weight)

	// 6. Constants
	fmt.Println("6. Constants:")
	const pi = 3.14159
	const greeting = "Hello, World!"
	const maxUsers = 100
	fmt.Printf("Pi: %.5f, Greeting: %s, Max Users: %d\n\n", pi, greeting, maxUsers)

	// 7. Type conversion
	fmt.Println("7. Type Conversion:")
	var num int = 42
	var floatNum float64 = float64(num)
	var str string = fmt.Sprintf("%d", num)
	fmt.Printf("Int: %d, Float: %.2f, String: %s\n\n", num, floatNum, str)

	// 8. Different numeric types
	fmt.Println("8. Numeric Types:")
	var (
		int8Val    int8    = 127                 // -128 to 127
		uint8Val   uint8   = 255                 // 0 to 255
		int16Val   int16   = 32767               // -32768 to 32767
		int32Val   int32   = 2147483647          // -2^31 to 2^31-1
		int64Val   int64   = 9223372036854775807 // -2^63 to 2^63-1
		float32Val float32 = 3.14159
		float64Val float64 = 3.141592653589793
	)
	fmt.Printf("int8: %d, uint8: %d, int16: %d\n", int8Val, uint8Val, int16Val)
	fmt.Printf("int32: %d, int64: %d\n", int32Val, int64Val)
	fmt.Printf("float32: %f, float64: %f\n\n", float32Val, float64Val)

	// 9. String operations
	fmt.Println("9. String Operations:")
	firstName2 := "Jane"
	lastName2 := "Smith"
	fullName := firstName2 + " " + lastName2
	fmt.Printf("Full Name: %s\n", fullName)
	fmt.Printf("Length: %d characters\n\n", len(fullName))

	// 10. Boolean operations
	fmt.Println("10. Boolean Operations:")
	hasLicense := true
	hasExperience := false
	canDrive := hasLicense && hasExperience
	needsTraining := !hasExperience
	fmt.Printf("Has License: %t, Has Experience: %t\n", hasLicense, hasExperience)
	fmt.Printf("Can Drive: %t, Needs Training: %t\n\n", canDrive, needsTraining)

	// 11. Variable reassignment
	fmt.Println("11. Variable Reassignment:")
	score := 80
	fmt.Printf("Initial score: %d\n", score)
	score = 95
	fmt.Printf("Updated score: %d\n\n", score)

	// 12. Scope example
	fmt.Println("12. Variable Scope:")
	globalVar := "I'm in main function"
	fmt.Println(globalVar)

	if true {
		localVar := "I'm inside if block"
		fmt.Println(localVar)
		fmt.Println(globalVar) // Can access outer scope
	}
	// fmt.Println(localVar) // Error: localVar tidak bisa diakses di sini
}
