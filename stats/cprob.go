package main

import (
	"fmt"
	"github.com/atgjack/prob"
	"flag"
	"log"
	"os"
)

const (
	LIVERATE = 0.6
)
func main(){
	var uppernodes int
	flag.IntVar(&uppernodes, "u", 50, "Sets the upper limit of nodes to compute")
	flag.Parse()
	f, err := os.Create("cprob.dat")
	if err != nil{
		log.Fatal("Problem creating data file")
	}
    
	for i := 1; i <= uppernodes; i++{
		binom, err := prob.NewBinomial(float64(i), LIVERATE)
		if err != nil{
			log.Fatal("Bad binomial")
		}
		consp := binom.Cdf(float64(int((i/2)+1)))
		consp = 1 - consp
		fmt.Printf("%d\t%f\n", i, consp)
		f.WriteString(fmt.Sprintf("%d\t%f\n", i, consp))
	}
	f.Sync()
	f.Close()
}