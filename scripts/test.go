package main

import (
	"os/exec"
	"log"
	"strings"
	"time"
	"fmt"
)

const (
	DINGHY_BIN = "dinghy1.1"
)
func main(){
	fmt.Println(startCluster(5))
}
func getNodes()([]string){
	out, err := exec.Command("pgrep", DINGHY_BIN).Output()
	if err != nil {
        log.Fatal(err)
    }
    pout := strings.TrimSuffix(string(out), "\n")
    nodepids := strings.Split(pout, "\n")
    return nodepids
}

func down(pid string){
	exec.Command("kill", pid).Output()
}

func startCluster(size int)([]int){
	pids := make([]int,0)
	initcmd := exec.Command("./"+DINGHY_BIN, "-r", "7000", "-h", "8000", "--bootstrap")
	initcmd.Start()
	pids = append(pids, initcmd.Process.Pid)
	time.Sleep(1 * time.Second)
	for i:=1; i<size-1; i++{
		cmd := exec.Command("./"+DINGHY_BIN, "-r", string(7000+i), "-h", string(8000+i), "--join=127.0.0.1:8000")
		cmd.Start()
		pids = append(pids, cmd.Process.Pid)
	}
	lastcmd := exec.Command("./"+DINGHY_BIN, "-r", string(7000+(size-1)), "-h", string(8000+(size-1)), "--join=127.0.0.1:8000")
	lastcmd.Start()
	pids = append(pids, lastcmd.Process.Pid)
	return pids
}