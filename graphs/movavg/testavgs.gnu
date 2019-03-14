l(x) = y0 + m*x
m(x) = y1 + b*x

set xlabel "Number of Node in Cluster"
set ylabel "Time (ms)"
set title "Test Time vs. Cluster Size"
set xrange [3:51]

fit l(x) '../data/dinghy.dat' via y0, m
fit m(x) '../data/vanilla.dat' via y1, b

plot l(x) lt rgb '#0011ff' title "Using Dinghy", \
     m(x) lt rgb '#FFA700' title "Without Dinghy"