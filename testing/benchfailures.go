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
	"math/rand"
)

func main(){
	var initial bool
	flag.BoolVar(&initial, "i", false, "Whether to create the necessary rows of the database")
	flag.Parse()
	fmt.Println(initial)
	// This test performs the same basic function as benchmark.go,
	// but also includes the 25% failure rate

	db, err := sql.Open("sqlite3", "../data/benchfailures.db")
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
	// and automates processes using bash script and sqlite database
	rnum := &RandNum{rand.New(rand.NewSource(time.Now().UnixNano()))}
	for i:=3; i<=MAXNODES; i+=2{
		np:=raft.Point{Nodes: i}
		for n:=0; n<TRIALSFOREACH; n++{
			cl:=raft.MakeCluster(i, true, nil)
			cl.Leader()
			model := make([]bool, i)
			for{
				toKill:=0
				for f:=0; f<i; f++{
					choice := rnum.PercentBernoulli(0.25)
					fmt.Println(choice)
					model[f] = choice
					if choice{
						toKill++
					}

				}
				color.Cyan("KILLING RATIO: %d/%d", toKill, i)
				if toKill <= i/2{
					color.Red("Killing %d nodes...", toKill)
					break
				}
			}

			for f:=0; f<i; f++{
				if model[f]{
					color.Red("Killed node%d", f)
					cl.Kill(f)
				}
			}

			time.Sleep(200 * time.Millisecond)
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

type RandNum struct {
	*rand.Rand
}

// p = probability of successful trial
func (r *RandNum)PercentBernoulli(p float64) bool {
	if r.Float64() <= p{return true}else{return false}
}