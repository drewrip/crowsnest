set style line 1 \
    linecolor rgb "#FFA700" \
    linetype 1 linewidth 1 \
    pointtype 2 pointsize 1

set xrange [1:51]
set xlabel "Number of Node in Cluster"
set ylabel "Probability of Reaching Consensus"
set title "Probability of Cluster Reaching Consensus vs. Cluster Size"

plot '../data/cprob.dat' with linespoints linestyle 1 title "", \