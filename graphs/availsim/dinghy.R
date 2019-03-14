# Compares Dinghy and Vanilla Raft in time required to recover from downed leader
library(ggplot2)
library(extrafont)
loadfonts()
dinghy <- read.csv("/home/drew/go/src/github.com/drewrip/dinghy/data/availsim/dinghytimes.csv")
raft <- read.csv("/home/drew/go/src/github.com/drewrip/dinghy/data/availsim/rafttimes.csv")
print(dinghy)
print(raft)
ggplot() +
  geom_line(data=dinghy, aes(y=time, x=size, color="Dinghy"))+
  geom_line(data=raft, aes(y=time, x=size, color="Vanilla Raft"))+
  scale_color_manual(name="Implementation", values = c('Dinghy' = '#0072B2', 'Vanilla Raft' = '#D55E00'))+
  ggtitle("Number of Nodes in a Cluster vs. Time To Recover")+
  xlab("Number of Nodes in Cluster") + 
  ylab("Time To Recover (ms)")+
  #geom_smooth(method="lm", se=FALSE, size=.75)+
  #annotate(geom="text", label=paste("R=", sprintf("%.3f", cor(bench$size, bench$time))))+
  #annotate(geom="text", label=paste("R^2=", sprintf("%.3f", cor(bench$size, bench$time)^2)))+
  theme_bw(base_family = "Times New Roman") +
  theme(plot.title = element_text(hjust = 0.5))

