# Prüfungsaufgaben zu Programmierung 1

[![Run on Repl.it](https://repl.it/badge/github/wwi21ama-prog/portfolio_prog1)](https://repl.it/github/wwi21ama-prog/portfolio_prog1)

Dieses Repo enthält Prüfungsaufgaben für die Vorlesung "Programmierung 1".

## Aufbau der Aufgaben
Zu jeder Aufgabe (außer der letzten) gibt es zwei Go-Dateien:
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

## Prüfungsmodalitäten

### Abgabe
* **Termin:** 24.12.2021, 00 Uhr
* Abgabe über **Moodle**
* Arbeit/Abgabe in Gruppen erlaubt (1-3 Personen pro "Gruppe")
    * Bitte unbedingt angeben, wie die Gruppen zusammengesetzt sind.

### Prüfungsgespräche
* Prüfungsgespräch mit jeder/jedem erforderlich.
    * Sollten bis zur Abgabe stattgefunden haben.
    * Ziele: Prüfung der Eigenständigkeit und Feedback/Hilfestellung
* Aufgaben müssen bis dahin nicht endgültig erledigt sein, es kann auch über Teillösungen geredet werden.
* Prüfungsgespräche dienen nicht primär der Notenfindung, können aber herangezogen werden.

## Bewertungskriterien

Für die Aufgaben 1-5 werden in erster Linie die Tests ausschlaggebend sein.
D.h. die Tests müssen laufen, die Signaturen der Funktionen sollen auch nicht verändert werden.
Neben den Tests aus diesem Repo wird es weitere Tests für jede der Funktionen geben, mit denen
andere Werte, Sonderfälle etc. geprüft werden.

Wer die Aufgaben 1-5 vollständig bearbeitet, die Tests aus diesem Repo zum Laufen bringt
und die Lösung im Prüfungsgespräch erklären kann, besteht die Prüfung.

Für eine gute Note müssen auch die Aufgaben 6 und 7 bearbeitet werden.
Die genaueren Bewertungskriterien für diese beiden Aufgaben stehen in den jeweiligen Go-Dateien.
Bei allen Aufgaben spielt für die Bewertung neben der Korrektheit auch der Programmierstil eine Rolle.
D.h. Sie sollten auf gute Lesbarkeit achten, z.B. sprechende Funktions- und Variablennamen verwenden
und Kommentare zur Erklärung einbauen, wo der Code nicht selbsterklärend ist.
