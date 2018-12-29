package main

import (
	"github.com/drewrip/dinghy/testing/raft"
	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"log"
	"database/sql"
	"flag"
	_ "github.com/mattn/go-sqlite3"
	"fmt"
	"time"
)

func main(){
	var initial bool
	flag.BoolVar(&initial, "i", false, "Whether to create the necessary rows of the database")
	flag.Parse()
	db, err := sql.Open("sqlite3", "../data/benchmarks.db")
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}
	defer db.Close()
	create, _ := db.Prepare("CREATE TABLE IF NOT EXISTS benchmarks (Nodes INTEGER, Time REAL, Trials INTEGER)")
	create.Exec()
	
	MAXNODES:=51
	TRIALSFOREACH:=1
	//data:=make([]int64,0)
	//data: numberOfNodes -> []consensusTime (ms)
	data:=make([]raft.Point, 0)
	// Running for i nodes


	// Runs the test for n = 1,3,5,...,51
	// Originally set up to take multiple trials per run, but now just does,
	// and automates processes using a bash script and sqlite database
	for i:=1; i<=MAXNODES; i+=2{
		np:=raft.Point{Nodes: i}
		for n:=0; n<TRIALSFOREACH; n++{
			cl:=raft.MakeCluster(i, true, nil)
			cl.Leader()
			np.Time += float64(cl.Test())
			cl.Close()
			time.Sleep(750 * time.Millisecond)
		}
		np.Trials = TRIALSFOREACH
		data = append(data, np)
		s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
		color.Cyan("Taking a quick pause to make sure everything can catch up...")
		s.Start()
		time.Sleep(500 * time.Millisecond)        
		s.Stop()
	}
	/*for i:=0; i<len(data); i++{
		data[i].Time = float64(data[i].Time)/float64(data[i].Trials)
	}*/
	for i:=0; i<len(data); i++{
		fmt.Printf("%d\t%f\t%d\n", data[i].Nodes, data[i].Time, data[i].Trials)
		if initial{
			newNodeRow, err :=db.Prepare("INSERT INTO benchmarks (Nodes, Time, Trials) VALUES (?, ?, ?)")
			if err != nil{
				log.Fatal(err)
			}
			newNodeRow.Exec(data[i].Nodes, 0, 0)
		}
		updatedb, err := db.Prepare("UPDATE benchmarks SET Time = Time + ?, Trials = Trials + ? WHERE Nodes = ?")
		if err != nil{
			log.Fatal(err)
		}
		updatedb.Exec(data[i].Time, data[i].Trials, data[i].Nodes)
	}
}