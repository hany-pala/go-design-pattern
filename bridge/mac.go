package bridge

import "fmt"

type Mac struct {
	Printer Printer
}

func (m *Mac) Print() {
	fmt.Println("Print request for mac")
	m.Printer.PrintFile()
}

func (m *Mac) SetPrinter(p Printer) {
	m.Printer = p
}
