# Probability of reaching consensus for given cluster size and constant failure rate
library(ggplot2)
#library(extrafont)
cprob <- read.csv("/home/drew/go/src/github.com/drewrip/dinghy/stats/cprob.csv")
print(cprob)
ggplot() +
  geom_line(data=cprob, aes(y=prob, x=size), color="#0072B2", size=1)+
  ggtitle("Probability of Achieving Consensus Given a 25% Node Fail Rate")+
  xlab("Number of Nodes in Cluster") + 
  ylab("Frequency")+
  #geom_smooth(method="lm", se=FALSE, size=.75)+
  #annotate(geom="text", label=paste("R=", sprintf("%.3f", cor(bench$size, bench$time))))+
  #annotate(geom="text", label=paste("R^2=", sprintf("%.3f", cor(bench$size, bench$time)^2)))+
  theme_bw(base_family = "Times New Roman") +
  theme(plot.title = element_text(hjust = 0.5))

