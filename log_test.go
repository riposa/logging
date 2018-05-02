package logging

import (
	"testing"
	"os"
	"errors"
	"time"
	"log"
)

func TestNewRotateHandler(t *testing.T) {
	h := NewRotateHandler(
		"/Users/hengha/go/LogService",
		"test",
		"log",
		2,
	)
	l := &Logger{}
	l.Init()
	l.AddHandler(os.Stdout, h)

	err := errors.New("atewtaw")
	for {
		t1 := time.Now().UnixNano()
		l.Error(err)
		t2 := time.Now().UnixNano()
		log.Printf("time cost: %d μs", (t2 - t1) / 1000)
		time.Sleep(5 * time.Second)
	}

}
