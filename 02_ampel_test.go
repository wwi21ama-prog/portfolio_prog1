package main

import "fmt"

// Gibt Ergebnisse von colour() auf die Konsole aus.
// Der Kommentar unten gibt die erwarteten Ergebnisse an.
// Automatische Prüfung mittels des Befehls "go test" (statt "go run").
func Example_02_ampel() {
	fmt.Println(colour(0))
	fmt.Println(colour(9))
	fmt.Println(colour(10))
	fmt.Println(colour(20))
	fmt.Println(colour(30))
	fmt.Println(colour(40))

	// Output:
	// Rot
	// Rot
	// Rot-Gelb
	// Grün
	// Gelb
	// Rot
}
