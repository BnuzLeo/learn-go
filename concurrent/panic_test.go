package concurrent

import (
	"fmt"
	"testing"
	"time"
)

func TestPanic(t *testing.T) {
	for {
		go func() {
			defer func() {
				if err := recover(); err != nil {
					fmt.Println("finish", err)
				}
			}()
			time.Sleep(time.Second * 3)
			panic("hi")
		}()
		time.Sleep(time.Second * 3)
	}
}
