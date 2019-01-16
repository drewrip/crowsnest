package main

import (
	"github.com/drewrip/dinghy/testing/raft"
	"github.com/fatih/color"
	"fmt"
	"os"
	"log"
	"time"
)

func main(){

	NN:=3
	//data:=make([]int64,0)
	//data: numberOfNodes -> []consensusTime (ms)
	// Running for i nodes

	data:=make([]int64,0)

	// Runs the test for n = 1,3,5,...,51
	// Originally set up to take multiple trials per run, but now just does,
	// and automates processes using a bash script and sqlite database

	// n int, bootstrap bool, conf *Config, sp int, kl bool
	cl:=raft.MakeCluster(NN, true, nil, 5, false)

	color.Cyan("STARTING LOOP\n")
	//cl.StartRetarget()
	lastCh:=cl.GetNode(1).ContactCh()
	lasttime:=time.Now()

	for i:=0; i<5000; i++{
		select{
			case y:=<-lastCh:
				el:=time.Since(lasttime)
				lasttime=time.Now()
				data = append(data, int64(el/time.Millisecond))
				color.Green("node%d: Contact %d = %v\n", 1, i, y)
		}
	}
	
	//time.Sleep(2 * time.Second)
	//cl.StartRetarget()
	//fmt.Println(float64(cl.Test(100)))
	f, err := os.Create("cdistrib.dat")
	if err != nil{
		log.Fatal("Problem creating data file")
	}
	fmt.Println(data)
	for _,n := range data{
		f.WriteString(fmt.Sprintf("%d\n", n))
	}
	f.Sync()
	f.Close()

}