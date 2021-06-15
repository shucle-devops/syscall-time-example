package main
import (
	"time"
	"github.com/yeongjukang/syscall-time-example/syscalltimeexample"
)

func main() {
        err := syscalltimeexample.SetTimeWithSyscall(time.Now().Add(time.Minute * 1))
        if err != nil {
                panic(err)
        }
        err2 := syscalltimeexample.SetTimeWithSyscall(time.Now().Add(time.Minute * -1))
        if err2 != nil {
                panic(err2)
        }      

}
