package latex

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

type Point struct {
	Label string
	X, Y  float64
}

func (pt *Point) ToString() string {
	return fmt.Sprintf("(%f,%f)", pt.X, pt.Y)
}

type Entry struct {
	MathMode bool
	Value    string
}

func (entry *Entry) ToString() string {
	if entry.MathMode {
		return "$" + entry.Value + "$"
	}
	return entry.Value
}

type Table struct {
	CornerValue      string
	HorizontalLookup []string
	VerticalLookup   []string
	Entries          [][]Entry
}

func MakeTable(fileName string, entries [][]Entry) {
	output := bytes.NewBufferString("")
	for _, row := range entries {
		str := ""
		for colNumber, col := range row {
			if colNumber == 0 {
				str += col.ToString() + " "
			} else {
				str += "& " + col.ToString() + " "
			}
		}
		output.WriteString(str + "\\\\ \\hline\n")
	}
	ioutil.WriteFile(fileName, output.Bytes(), 0777)
}

func MakeTableWithLookup(fileName string, tbl Table) {
	output := bytes.NewBufferString(tbl.CornerValue + " ")
	for _, topVal := range tbl.HorizontalLookup {
		output.WriteString("& " + topVal + " ")
	}
	output.WriteString("\\\\ \\hline\n")
	for rowNumber, row := range tbl.Entries {
		str := tbl.VerticalLookup[rowNumber] + " "
		for _, col := range row {
			str += "& " + col.ToString() + " "
		}
		output.WriteString(str + "\\\\ \\hline\n")
	}
	ioutil.WriteFile(fileName, output.Bytes(), 0777)
}

func MakeColoredPointGraph(fileName string, labelToMark map[string]string, points []Point) {
	output := bytes.NewBufferString("")
	for label, mark := range labelToMark {
		output.WriteString("\\addplot+[only marks, mark=" + mark + "] coordinates{\n")
		for _, p := range points {
			if p.Label == label {
				output.WriteString(p.ToString() + "\n")
			}
		}
		output.WriteString("};\n")
	}
	ioutil.WriteFile(fileName, output.Bytes(), 0777)
}

func MakeLineGraph(fileName string, points []Point) {
	output := bytes.NewBufferString("\\addplot [black] coordinates {\n")
	for _, p := range points {
		output.WriteString(p.ToString() + "\n")
	}
	output.WriteString("};")
	ioutil.WriteFile(fileName, output.Bytes(), 0777)
}

func SelectPoints(pts []Point, label string) []Point {
	newpts := []Point{}
	for _, p := range pts {
		if p.Label == label {
			newpts = append(newpts, p)
		}
	}
	return newpts
}
