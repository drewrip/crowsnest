set style line 1 \
    linecolor rgb '#348700' \
    linetype 1 linewidth 1 \
    pointtype 2 pointsize 1
set style line 2 \
    linecolor rgb '#800098' \
    linetype 1 linewidth 1 \
    pointtype 4 pointsize 1

set xlabel "Number of Node in Cluster"
set ylabel "Time (ms)"
set title "Test Time vs. Cluster Size"

plot '../data/dinghy.dat' with linespoints linestyle 1 title "Using Dinghy", \
     '../data/vanilla.dat' with linespoints linestyle 2 title "Without Dinghy"