package main

import (
	"fmt"
	"github.com/inspiration33/design/bridge"
)

func main() {
	hpPrinter := &bridge.Hp{}
	epsonPrinter := &bridge.Epson{}

	macComputer := &bridge.Mac{}

	macComputer.SetPrinter(hpPrinter)
	macComputer.Print()
	fmt.Println()

	macComputer.SetPrinter(epsonPrinter)
	macComputer.Print()
	fmt.Println()

	winComputer := &bridge.Windows{}

	winComputer.SetPrinter(hpPrinter)
	winComputer.Print()
	fmt.Println()

	winComputer.SetPrinter(epsonPrinter)
	winComputer.Print()
	fmt.Println()
}

//func main() {
//	client := &adapter.Client{}
//	mac := &adapter.Mac{}
//
//	client.InsertLightningConnectorIntoComputer(mac)
//
//	windowsMachine := &adapter.Windows{}
//	windowsMachineAdapter := &adapter.WindowsAdapter{
//		WindowMachine: windowsMachine,
//	}
//
//	client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)
//}

//func main() {
//	normalBuilder := builder.GetBuilder("normal")
//	iglooBuilder := builder.GetBuilder("igloo")
//
//	director := builder.NewDirector(normalBuilder)
//	normalHouse := director.BuildHouse()
//
//	fmt.Printf("normal House Door Type: %s\n", normalHouse.DoorType)
//	fmt.Printf("normal House Window Type: %s\n", normalHouse.WindowType)
//	fmt.Printf("normal House Num Floor: %d\n", normalHouse.Floor)
//
//	director.SetBuilder(iglooBuilder)
//	iglooHouse := director.BuildHouse()
//
//	fmt.Printf("normal House Door Type: %s\n", iglooHouse.DoorType)
//	fmt.Printf("normal House Window Type: %s\n", iglooHouse.WindowType)
//	fmt.Printf("normal House Num Floor: %d\n", iglooHouse.Floor)
//}

//func main() {
//	adidasFactory, _ := abstract_factory.GetSportsFactory("adidas")
//	nikeFactory, _ := abstract_factory.GetSportsFactory("nike")
//
//	nikeShoe := nikeFactory.MakeShoe()
//	nikeShirt := nikeFactory.MakeShirt()
//
//	adidasShoe := adidasFactory.MakeShoe()
//	adidasShirt := adidasFactory.MakeShirt()
//
//	printShoeDetails(nikeShoe)
//	printShirtDetails(nikeShirt)
//
//	printShoeDetails(adidasShoe)
//	printShirtDetails(adidasShirt)
//}
//
//func printShoeDetails(s abstract_factory.IShoe) {
//	fmt.Printf("Logo: %s", s.GetLogo())
//	fmt.Println()
//	fmt.Printf("Size: %d", s.GetSize())
//	fmt.Println()
//}
//
//func printShirtDetails(s abstract_factory.IShirt) {
//	fmt.Printf("Logo: %s", s.GetLogo())
//	fmt.Println()
//	fmt.Printf("Size: %d", s.GetSize())
//	fmt.Println()
//}

//func main() {
//
//	for i := 0; i < 30; i++ {
//		//go singleton.GetInstance()
//		//go singleton.GetInstanceCaseTwo()
//	}
//
//	// Scanln is similar to Scan, but stops scanning at a newline and
//	// after the final item there must be a newline or EOF.
//	fmt.Scanln()
//}

//func main() {
//	file1 := &prototype.File{Name: "File1"}
//	file2 := &prototype.File{Name: "File2"}
//	file3 := &prototype.File{Name: "File3"}
//
//	folder1 := &prototype.Folder{
//		Children: []prototype.Inode{file1},
//		Name:     "Folder1",
//	}
//
//	folder2 := &prototype.Folder{
//		Children: []prototype.Inode{folder1, file2, file3},
//		Name:     "Folder2",
//	}
//
//	fmt.Println("\nPrinting hierarchy for Folder2")
//	folder2.Print(" ")
//
//	cloneFolder := folder2.Clone()
//	fmt.Println("\nPrinting hierarchy for clone Folder")
//	cloneFolder.Print(" ")
//}

//func printDetails(g factory.IGun) {
//	fmt.Printf("Gun: %s", g.GetName())
//	fmt.Println()
//	fmt.Printf("Power: %d", g.GetPower())
//	fmt.Println()
//}
