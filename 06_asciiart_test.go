package main

// Gibt Ergebnisse von asciiart() auf die Konsole aus.
// Der Kommentar unten gibt die erwarteten Ergebnisse an.
// Automatische Prüfung mittels des Befehls "go test" (statt "go run").
func Example_06_asciiart() {
	asciiart("hallo welt")

	// Output:
	// #     #   ##   #      #       ####     #  #  # ###### #      #####
	// #     #  #  #  #      #      #    #    #  #  # #      #        #
	// ####### #    # #      #      #    #    #  #  # #####  #        #
	// #     # ###### #      #      #    #    #  #  # #      #        #
	// #     # #    # #      #      #    #    #  #  # #      #        #
	// #     # #    # ###### ######  ####      ## ##  ###### ######   #
}
