package main

import "fmt"

// Gibt Ergebnisse von sums() auf die Konsole aus.
// Der Kommentar unten gibt die erwarteten Ergebnisse an.
// Automatische Prüfung mittels des Befehls "go test" (statt "go run").
func Example_05_arrays() {
	fmt.Println(sums([]int{1, 3, 5, 7}))  // Soll 1 4 9 16 ausgeben.
	fmt.Println(sums([]int{1, 1, 2, 80})) // Soll 1 2 4 84 ausgeben.
	fmt.Println(sums([]int{7, 3, 1, 2}))  // Soll 7 10 11 13 ausgeben.
	fmt.Println(sums([]int{3, 3, 0, 2}))  // Soll 3 6 6 8 ausgeben.

	// Output:
	// [1 4 9 16]
	// [1 2 4 84]
	// [7 10 11 13]
	// [3 6 6 8]
}
