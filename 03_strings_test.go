package main

import "fmt"

// Gibt Ergebnisse von zip() auf die Konsole aus.
// Der Kommentar unten gibt die erwarteten Ergebnisse an.
// Automatische Prüfung mittels des Befehls "go test" (statt "go run").
func Example_03_strings() {

	fmt.Println(zip("abc", "def"))
	fmt.Println(zip("bab", "abacc"))
	fmt.Println(zip("", ""))

	// Output:
	// adbecf
	// baabbacc
	//
}
