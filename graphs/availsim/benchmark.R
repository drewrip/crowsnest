# Setup
library(ggplot2)
library(extrafont)
loadfonts()
bench <- read.csv("/home/drew/go/src/github.com/drewrip/dinghy/data/availsim/benchmark.csv")
print(bench)
ggplot(bench, aes(y=time, x=size)) +
  ggtitle("Number of Nodes in a Cluster vs. Avg. Time to Complete Test")+
  geom_line(size=.75) +
  xlab("Number of Nodes in Cluster") + 
  ylab("Avg. Time to Complete Test")+
  geom_smooth(method="lm", se=FALSE, size=.75)+
  annotate(geom="text", label=paste("R=", sprintf("%.3f", cor(bench$size, bench$time))), x=49, y=1955)+
  annotate(geom="text", label=paste("R^2=", sprintf("%.3f", cor(bench$size, bench$time)^2)), x=49, y=1960)+
  theme_bw(base_family = "Times New Roman") +
  theme(plot.title = element_text(hjust = 0.5))
