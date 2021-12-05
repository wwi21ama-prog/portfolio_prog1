package main

import "fmt"

// Gibt Ergebnisse von power2() auf die Konsole aus.
// Der Kommentar unten gibt die erwarteten Ergebnisse an.
// Automatische Prüfung mittels des Befehls "go test" (statt "go run").
func Example_04_rekursion() {

	fmt.Println(power2(2))
	fmt.Println(power2(5))
	fmt.Println(power2(0))

	// Output:
	// 4
	// 32
	// 1
}
