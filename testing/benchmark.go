package main

import (
	"github.com/drewrip/dinghy/testing/raft"
	//"github.com/fatih/color"
	"log"
	"database/sql"
	"flag"
	_ "github.com/mattn/go-sqlite3"
	"fmt"
	"time"
)

func main(){
	TESTNUM:=1000
	var initial bool
	var nodenum int
	var retarget bool
	var dbname string
	flag.BoolVar(&initial, "i", false, "Whether to create the necessary rows of the database")
	flag.BoolVar(&retarget, "r", false, "Whether or not to retarget timeouts and interval")
	flag.IntVar(&nodenum, "n", -1, "Number of nodes to run test on")
	flag.StringVar(&dbname, "d", "test", "Name of the database to pipe data to")
	flag.Parse()
	db, err := sql.Open("sqlite3", fmt.Sprintf("../data/%s.db", dbname))
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}
	defer db.Close()
	create, _ := db.Prepare("CREATE TABLE IF NOT EXISTS benchmarks (Nodes INTEGER, Time INTEGER, Trials INTEGER)")
	create.Exec()

	if initial{
		fmt.Printf("Creating Database\n")
		for i:=3; i<=51; i+=2{
			newNodeRow, err :=db.Prepare("INSERT INTO benchmarks (Nodes, Time, Trials) VALUES (?, ?, ?)")
			if err != nil{
				log.Fatal(err)
			}
			newNodeRow.Exec(i, 0, 0)	
		}
		fmt.Printf("Initialized Database\n")
	}
	if nodenum != -1{
		cl:=raft.MakeCluster(nodenum, true, nil, 3, false, TESTNUM/4)
		if retarget{
			//color.Blue("\n\nInitialized Retarget: %v\n", 500 * time.Millisecond)
			//cl.StartRetarget(500 * time.Millisecond)
		}
		testTime := cl.Test(TESTNUM)/1e6
		fmt.Printf("%d\t%d\t%d\n", nodenum, testTime, 1)
		time.Sleep(500 * time.Millisecond)
		updatedb, err := db.Prepare("UPDATE benchmarks SET Time = Time + ?, Trials = Trials + ? WHERE Nodes = ?")
		if err != nil{
			log.Fatal(err)
		}
		updatedb.Exec(testTime, 1, nodenum)	
	}

}