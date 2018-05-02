package logging

import (
	"errors"
	"log"
	"os"
	"testing"
	"time"
)

func TestNewRotateHandler(t *testing.T) {
	curDir, _ := os.Getwd()
	h := NewRotateHandler(
		curDir,
		"test",
		"log",
		2,
	)
	l := &Logger{}
	l.Init()
	l.AddHandler(os.Stdout, h)

	err := errors.New("test1234567890")
	for {
		t1 := time.Now().UnixNano()
		l.Error(err)
		t2 := time.Now().UnixNano()
		log.Printf("time cost: %d μs", (t2-t1)/1000)
		time.Sleep(5 * time.Second)
	}

}
