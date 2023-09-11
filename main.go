// You can edit this code!
// Click here and start typing.
package main

import (
	"flag"
	"fmt"
	"math/rand"
	"slices"
)

var notesSharp = []string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}
var notesFlat = []string{"C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab", "A", "Bb", "B"}

var nFlag = flag.Int("s", 1234, "help message for flag n")

func pickRootIndex() int {
	return rand.Intn(12)
}

func pickBeats() int {
	beats := []int{2, 3, 4, 6}
	return beats[rand.Intn(4)]
}
func Seed() {}

//	func pickScale() {
//		scales := []string{
//			"major",
//			"pentatonic",
//			"phyrgian dominant",
//			"hungarian minor",
//			"harmonic minor",
//			"melodic minor",
//			"blues",
//		}
//	}
func pickMode() string {
	beats := []string{"interval", "notes", "blank"}
	return beats[rand.Intn(3)]

}

func getIndex(elem string, list []string) int {
	for i := 0; i < cap(list); i++ {
		if elem == list[i] {
			return i
		}
	}
	return -1
}

func getNotesInPattern(root string, pattern string) []string {
	chromaticScale := []string{}

	index := getIndex(root, notesSharp)
	sliceOne := notesSharp[index:]
	sliceTwo := notesSharp[0 : index+1]
	chromaticScale = append(sliceOne, sliceTwo...)
	out := []string{}
	dict := map[string][]int{
		"major-triad": []int{0, 4, 7},
	}
	_pattern := dict[pattern]

	for _, e := range _pattern {
		out = append(out, chromaticScale[e])
	}

	return out
}

func getDiagramPart(scale []string, pattern string, root string) string {
	out := ""

	// TODO fill this in

	notesInPattern := getNotesInPattern(root, pattern)

	for _, e := range scale {
		if slices.Contains(notesInPattern, e) {
			for _, f := range notesInPattern {
				if e == f {
					newString := f
					if len(f) > 1 {
						out = out + newString + "|"

					} else {
						out = out + "-" + newString + "|"
					}
				}

			}

		} else {
			out = out + "--|"
		}

	}

	return out
}

func printFretboardDiagram(root string, pattern string) {
	// fretDefault := "-|--|--|--|--|--|--|--|--|--|--|--|--|"
	fretLegend := "---------------------------------------\n O| 1| 2| 3| 4| 5| 6| 7| 8| 9|10|11|12|"

	tuning := getTuning()

	strings := make([]string, 0)

	for _, s := range tuning {
		index := getIndex(s, notesSharp)
		sliceOne := notesSharp[index:]
		sliceTwo := notesSharp[0 : index+1]
		scale := append(sliceOne, sliceTwo...)
		fretDiagramPart := getDiagramPart(scale, pattern, root)
		strings = append(strings, fretDiagramPart)
	}

	for _, s := range strings {
		fmt.Println(s)
	}

	fmt.Println(fretLegend)
}

func getTuning() []string {
	return []string{"G", "D", "A", "E"}
}

func main() {
	/*
		todo:
		- return and pass seed as argument, read yaml file
		- fret modes: intervals (R, 3, 5) vs notes (C, E, G) vs blank
		- extend list of available patterns (
			inversions
			triads, sevenths, extendeds,
			modes,
			jazz, blues & pentatonic)
		- improve accuracy note names ( f vs e# )
		- [x] scale, arpeggio picking
		- [x] drawing fretboard

		- midi input?
		- walking bassline?
		- chord changes mode

	*/
	rootIndex := pickRootIndex()
	beats := pickBeats()
	// displayMode := pickMode()

	rootNote := notesSharp[rootIndex]
	patternType := "major-triad"

	notesSharp := []string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}
	// notesFlat := []string{"C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab", "A", "Bb", "B"}
	fmt.Println("root (sharp):  ", rootNote)
	// fmt.Println("root  (flat):  ", notesFlat[rootIndex])
	// fmt.Println("display mode:  ", displayMode)
	fmt.Println("beats: ", beats)
	printFretboardDiagram(notesSharp[rootIndex], patternType)
}
