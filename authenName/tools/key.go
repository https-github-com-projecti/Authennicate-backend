package tools

import (
	repo "authenName/repository"
	"fmt"
	"math/rand"
	"time"
)

func GenerateKey() int64 {
	var randomNum int64
	for {
		rand.Seed(time.Now().UnixNano())
		randomNum = int64(Random(100000, 999999))
		fmt.Printf("Random Num: %d\n", randomNum)
		_, err := repo.FindKey(randomNum)
		if err != nil {
			break
		}
	}
	return randomNum
}

func Random(min int, max int) int {
	return rand.Intn(max-min) + min
}
