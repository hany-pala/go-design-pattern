package adapter

import "fmt"

type Mac struct{}

func (w *Mac) InsertIntoLightningPort() {
	fmt.Println("USB connector is plugged into windows machine.")
}
