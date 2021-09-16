package main

import (
	"fmt"
	"sync"
	"time"
)

type AccessToken struct {
	Token     string
	CreatedAt time.Time
}

type AccessTokens map[string]AccessToken

var accessTokensLock sync.Mutex

type Values struct {
	Value int
	Lock  sync.Mutex
}

func main() {
	var wg sync.WaitGroup
	v := Values{
		Value: 0,
	}

	wg.Add(1)
	for i := 1; i <= 100; i++ {
		//go func(v *Values, index int) {
		//	wg.Wait()
		//
		//	v.Lock.Lock()
		//	defer v.Lock.Unlock()
		//
		//	fmt.Println(index, "Current value:", v.Value)
		//	fmt.Println(index, "Changing value to:", v.Value +  index)
		//
		//	v.Value = v.Value + index
		//
		//	time.Sleep(1 * time.Second)
		//
		//}(&v, i)
		go addValue(&wg, &v, i)
	}

	wg.Done()

	time.Sleep(30 * time.Second)
}

func addValue(wg *sync.WaitGroup, v *Values, i int) {
	wg.Wait()
	v.Lock.Lock()
	defer v.Lock.Unlock()
	fmt.Println(i, "Current value:", v.Value)
	fmt.Println(i, "Changing value to:", v.Value+i)
	v.Value = v.Value + i
	time.Sleep(1 * time.Second)
}
