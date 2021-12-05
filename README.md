# Prüfungsaufgaben zu Programmierung 1

## Aufbau der Aufgaben
Zu jeder Aufgabe gibt es mindestens zwei Go-Dateien:
In einer Datei wird die Aufgabe formuliert, in der anderen stehen Tests für die Aufgabe.
Die Test-Datei heißt genau wie die Aufgaben-Datei mit einem zusätzlich angehängten "_test".
Beispiel: `01_arrays.go` enthäkt die erste Aufgabe zu Arrays, `01_arrays_test.go` enthält
den zugehörigen Testcode.

Implementieren Sie Ihre Lösung in der Aufgaben-Datei und nutzen Sie die Funktion(en) aus der Testdatei,
um Ihre Lösung zu testen.
* Sie können die Tests automatisch ausführen lassen, indem Sie auf der Konsole
den Befehl `go test` verwenden. Dies führt alle Funktionen aus Test-Dateien aus,
die mit `Example_` beginnen, und vergleicht deren Ausgabe mit dem, was in der jeweiligen Funktion unten
unter der Zeile `// Output:` steht.
* Sie können die Tests natürlich auch manuell in der `main()`-Funktion in `main.go` ausführen.
