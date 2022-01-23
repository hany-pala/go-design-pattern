package singleton

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type Single struct {
}

var SingleInstance *Single

func GetInstance() *Single {
	if SingleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if SingleInstance == nil {
			fmt.Println("creating single instance now!")
			SingleInstance = &Single{}
		} else {
			fmt.Println("single instance already created!")
		}
	} else {
		fmt.Println("single instance already created!!!!!")
	}
	return SingleInstance
}
