package writer

import (
	"bytes"
	"strings"
)

const endOfLine = "\r\n"

type Register interface {
	ToArray() []string
	Type() string
}

type Writer struct {
	buf           *bytes.Buffer
	rows          []Register
	cleanRows     []Register
	linesReg      map[string]int
	linesBlock    map[string]int
	totalLines    int
	ArrayToString func([]string) string
}

func NewWriter(buf *bytes.Buffer, arrayToString func([]string) string) *Writer {

	return &Writer{
		buf:           buf,
		linesReg:      make(map[string]int),
		linesBlock:    make(map[string]int),
		ArrayToString: arrayToString,
	}
}

func (w *Writer) AddRegister(r Register) {
	if w.notExist(r) {
		w.cleanRows = append(w.cleanRows, r)
	}
	w.rows = append(w.rows, r)

	w.totalLines++
	//w.linesBlock[r.Block()]++
	w.linesReg[r.Type()]++
}
func (w *Writer) notExist(r Register) bool {
	for _, r1 := range w.cleanRows {
		if r.Type() == r1.Type() {
			return false
		}
	}
	return true
}

func (w *Writer) WriteRegister() {
	for _, r := range w.rows {
		// this return error, but the error will come always as nil, since the WriteString, panic if the buf become too large
		w.buf.WriteString(w.ArrayToString(r.ToArray()))
	}
}

func ArrayToStringLine(s []string) string {
	return "|" + strings.Join(s, "|") + "|" + endOfLine
}

func ArrayToStringCSV(s []string) string {
	return strings.Join(s, ";") + endOfLine
}
