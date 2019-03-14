// elections.go calculates the probability of a split election given some latency l, and a number of nodes n
// in this instance we assume that only two nodes have split the vote, and no nodes have been downed

package main

import(
	"fmt"
	"github.com/atgjack/prob"
	"math"
)

// Function defined pg. 125 https://ramcloud.stanford.edu/~ongaro/thesis.pdf
func splitprob(n, l float64) float64 {
	var total float64
	k:=float64(2)
	for ; k<=n; k++{
		total+= prob.BinomialCoefficient(n, k) * math.Pow(l, k) * math.Pow((1-l),(n-k))
	}
	return total
}

func main(){
	fmt.Println("size,prob")
	for i:=3; i<53; i+=2{
		fmt.Printf("%d,%f\n", i, splitprob(float64(i), .1))
	}
}