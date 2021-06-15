package syscalltimeexample

import (
	"os/exec"
	"syscall"
	"time"
)

func SetTimeWithSyscall(t time.Time) error {
	timeVal := syscall.NsecToTimeval(t.UnixNano())
	err := syscall.Settimeofday(&timeVal)
	return err
}

func SetTimeWithExec(t time.Time) error {
	timeVal := t.Format("2006-01-02 15:04:05")
	cmd := exec.Command("sudo", "date", "-s", timeVal)
	err := cmd.Run()
	return err
}

