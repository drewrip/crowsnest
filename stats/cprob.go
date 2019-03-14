package main

import (
	"fmt"
	"github.com/atgjack/prob"
	"flag"
	"log"
	"os"
)

const (
	LIVERATE = 0.75
)
func main(){
	var uppernodes int
	flag.IntVar(&uppernodes, "u", 51, "Sets the upper limit of nodes to compute, it is recommended you chose an odd number. MIN=3")
	flag.Parse()
	f, err := os.Create("cprob.csv")
	if err != nil{
		log.Fatal("Problem creating data file")
	}
    fmt.Println("size,prob")
	for i := 3; i <= uppernodes; i+=2{
		binom, err := prob.NewBinomial(float64(i), LIVERATE)
		if err != nil{
			log.Fatal("Bad binomial")
		}
		consp := binom.Cdf(float64(int((i/2)+1)))
		consp = 1 - consp
		fmt.Printf("%d,%f\n", i, consp)
		f.WriteString(fmt.Sprintf("%d,%f\n", i, consp))
	}
	f.Sync()
	f.Close()
}
