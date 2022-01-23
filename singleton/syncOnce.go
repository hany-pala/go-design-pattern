package singleton

import (
	"fmt"
	"sync"
)

var once sync.Once

type SingleCaseTwo struct {
}

var SingleCaseTwoInstance *SingleCaseTwo

func GetInstanceCaseTwo() *SingleCaseTwo {
	if SingleCaseTwoInstance == nil {
		once.Do(
			func() {
				fmt.Println("creating single instance now")
				SingleCaseTwoInstance = &SingleCaseTwo{}
			})
	} else {
		fmt.Println("single instance already created@@")
	}
	return SingleCaseTwoInstance
}
