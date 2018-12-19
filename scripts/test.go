package main

import (
	"os/exec"
	"os"
	"log"
	"fmt"
	"time"
)

const (
	DINGHY_BIN = "dinghy1.1"
)

type Node struct {
	Process *os.Process
	HTTPPort int
	RaftPort int
}

func main(){
	cluster:=StartCluster(5)
	time.Sleep(2 * time.Second)
	cluster[0].Kill()
	time.Sleep(10 * time.Second)
}

func Start(args ...string) (p *os.Process, err error) {
	if args[0], err = exec.LookPath(args[0]); err == nil {
		var procAttr os.ProcAttr
		procAttr.Files = []*os.File{os.Stdin,
			os.Stdout, os.Stderr}
		p, err := os.StartProcess(args[0], args, &procAttr)
		if err == nil {
			return p, nil
		}
	}
	return nil, err
}

// Returns references to node processes
func StartCluster(n int)([]Node){
	pcs := make([]Node, 0)
	pi, err := Start("./"+DINGHY_BIN, "-h8000", "-r7000", "--bootstrap")
	if err != nil{
			log.Fatal(err)
	}
	go pi.Wait()
	pcs = append(pcs, Node{Process: pi, HTTPPort: 8000, RaftPort: 70002})
	time.Sleep(2 * time.Second)
	for i:=1; i<n; i++{
		harg:=fmt.Sprintf("-h%d", 8000+i)
		parg:=fmt.Sprintf("-r%d", 7000+i)
		p, err := Start("./"+DINGHY_BIN, harg, parg, "--join=127.0.0.1:8000")
		if err != nil{
			log.Fatal(err)
		}
		go p.Wait()
		pcs = append(pcs, Node{Process: p, HTTPPort: 8000+i, RaftPort: 7000+i})
		time.Sleep(500 * time.Millisecond)
	}
	return pcs
}

func (n *Node)Kill()(error){
	err := n.Process.Kill()
	return err
}