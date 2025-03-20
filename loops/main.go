package main

import "fmt"

func main() {
	// 	// Basic for loop
	// 	fmt.Println("Basic for loop:")
	// 	for i := 1; i < 5; i++ {
	// 		fmt.Printf("%d ", i)
	// 	}
	// 	fmt.Println()

	// 	// While-style loop
	// 	fmt.Println("\nWhile-style loop:")
	// 	i := 1
	// 	for i < 5 {
	// 		fmt.Printf("%d ", i)
	// 		i++
	// 	}
	// 	fmt.Println()

	// // Infinite loop with break
	// fmt.Println("\nInfinite loop with break:")
	// count := 1
	// for {
	// 	if count > 4 {
	// 		break
	// 	}
	// 	fmt.Printf("%d ", count)
	// 	count++
	// }
	// fmt.Println()

	// Continue with label
	fmt.Println("\nContinue with label:")
here:
	for i := 0; i < 2; i++ {
		for j := i + 1; j < 3; j++ {
			if i == 0 {
				continue here
			}
			fmt.Printf("i=%d, j=%d\n", i, j)
			if j == 2 {
				break
			}
		}
	}

	fmt.Println()

	// Break with label
	fmt.Println("Break with label example:")
there:
	for i := 0; i < 2; i++ { // outer loop
		for j := i + 1; j < 3; j++ { // inner loop
			if j == 1 {
				continue // skips current iteration of inner loop
			}
			fmt.Printf("i=%d, j=%d\n", i, j)
			if j == 2 {
				break there // exits both loops completely
			}
		}
	}
}
