package main

import (
	"os"
	"os/exec"
    "syscall"
    "fmt"
)

func main() {
    switch os.Args[1] {
    case "run":
        parent()
    case "child":
        child()
    default:
        panic("invalid command")
    }
}

func parent() {
    cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
    cmd.SysProcAttr = &syscall.SysProcAttr{
        Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
    }
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    must(syscall.Sethostname([]byte("container")))
    must(syscall.Chroot("/home/den/liz/ubuntufs"))
    must(os.Chdir("/"))
    must(syscall.Mount("proc", "proc", "proc", 0, ""))
    must(cmd.Run())

    // if err := cmd.Run(); err != nil {
    //     fmt.Println("ERROR", err)
    //     os.Exit(1)
    // }
}

func child() {
    cmd := exec.Command(os.Args[2], os.Args[3:]...)
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    if err := cmd.Run(); err != nil {
        fmt.Println("ERROR", err)
        os.Exit(1)
    }
}

func must(err error) {
    if err != nil {
        panic(err)
    }
}
