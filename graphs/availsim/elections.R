# Probability of reaching consensus for given cluster size and constant failure rate
library(ggplot2)
#library(extrafont)
ele <- read.csv("/home/drew/go/src/github.com/drewrip/dinghy/stats/elections.csv")
print(ele)
ggplot() +
  geom_line(data=ele, aes(y=prob, x=size), color="#D55E00", size=1)+
  ggtitle("Frequency of an Election Split with One Downed Leader and 10ms Latency")+
  xlab("Number of Nodes in Cluster") + 
  ylab("Frequency")+
  #geom_smooth(method="lm", se=FALSE, size=.75)+
  #annotate(geom="text", label=paste("R=", sprintf("%.3f", cor(bench$size, bench$time))))+
  #annotate(geom="text", label=paste("R^2=", sprintf("%.3f", cor(bench$size, bench$time)^2)))+
  theme_bw(base_family = "Times New Roman") +
  theme(plot.title = element_text(hjust = 0.5))
