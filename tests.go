package main

import (
	"./latex"
)

func main() {
	entries := [][]latex.Entry{{{false, "a"}, {false, "b"}, {false, "c"}, {false, "d"}, {false, "e"}},
		{{true, "f"}, {true, "g"}, {true, "h"}, {true, "i"}, {true, "j"}}}
	latex.MakeTable("simple_table.tex", entries)
	cornerValue := "Mode/Letter"
	table := latex.Table{cornerValue, []string{"Letter 1", "Letter 1", "Letter 1", "Letter 1", "Letter 1"}, []string{"No Math Mode", "Math Mode"}, entries}
	latex.MakeTableWithLookup("table_with_lookup.tex", table)

	linePts := []latex.Point{{"", 0., 0.}, {"", 1., 1.}, {"", 2., 2.}, {"", 3., 3.}}
	coloredPts := []latex.Point{{"1", 0., 0.}, {"1", 1., 1.}, {"-1", 2., 2.}, {"-1", 3., 3.}}

	labelToMark := make(map[string]string)
	labelToMark["-1"] = "o"
	labelToMark["1"] = "x"

	latex.MakeLineGraph("line_graph.tex", linePts)
	latex.MakeColoredPointGraph("point_graph.tex", labelToMark, coloredPts)
}
