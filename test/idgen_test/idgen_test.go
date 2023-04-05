package idgen_test

import (
	"fmt"
	"github.com/yitter/idgenerator-go/idgen"
	"strconv"
	"sync"
	"testing"
)

func TestIdGen(t *testing.T) {
	options := idgen.NewIdGeneratorOptions(1)
	idgen.SetIdGenerator(options)
	var wg sync.WaitGroup
	var syncMap sync.Map
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			id := idgen.NextId()
			syncMap.Store(id, id)
		}()
	}
	wg.Wait()
	length := 0
	syncMap.Range(func(key, value interface{}) bool {
		length++
		return true
	})
	fmt.Println(length)
	fmt.Println(idgen.NextId())
	fmt.Println(len(strconv.Itoa(int(idgen.NextId()))))
}
