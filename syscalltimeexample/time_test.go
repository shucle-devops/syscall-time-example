package syscalltimeexample

import (
	"testing"
	"time"
)

func TestSetTimeWithSyscall(t *testing.T) {
	for i := 0; i < 1000; i++ {
		err := SetTimeWithSyscall(time.Now().Add(time.Minute * 1))
		if err != nil {
			panic(err)
		}
	}
}
func TestSetTimeWithExec(t *testing.T) {
	for i := 0; i < 1000; i++ {
		err2 := SetTimeWithExec(time.Now().Add(time.Minute * -1))
		if err2 != nil {
			panic(err2)
		}
	}
}

