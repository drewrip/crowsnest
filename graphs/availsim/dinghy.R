# Compares Dinghy and Vanilla Raft in time required to recover from downed leader
library(ggplot2)
library(extrafont)
loadfonts()

datize <- function(x){
  datamding <- data.matrix(x)
  trials <- nrow(datamding)
  sdsding <- colSds(datamding)
  meansding <- colMeans(datamding)
  cis <- qt(.99,trials-1)*sdsding/sqrt(trials)
  return(data.frame("size"=seq(3, 51, 2), "mean"=meansding, "sd"=sdsding, "ci"=cis, "trials"=trials))
}

dinghy <- read.csv("/home/drew/go/src/github.com/drewrip/dinghy/data/availsim/dinghytest.csv")
raft <- read.csv("/home/drew/go/src/github.com/drewrip/dinghy/data/availsim/rafttest.csv")

dData <- datize(dinghy)
rData <- datize(raft)
dData
rData
ggplot() +
  geom_line(data=dData, aes(y=dData$mean, x=dData$size, color="Dinghy n=659"))+
  geom_errorbar(width=.25, aes(x=rData$size, ymin=dData$mean-dData$ci, ymax=dData$mean+dData$ci),
                position=position_dodge(.9)) +
  geom_line(data=rData, aes(y=rData$mean, x=rData$size, color="Vanilla Raft n=715"))+
  geom_errorbar(width=.25, aes(x=rData$size, ymin=rData$mean-rData$ci, ymax=rData$mean+rData$ci),
                position=position_dodge(.9)) +
  scale_color_manual(name="Implementation", values = c('Dinghy n=659' = '#0072B2', 'Vanilla Raft n=715' = '#D55E00'))+
  ggtitle("Time For Cluster To Recover From Downed Leader with 99% Confidence")+
  xlab("Number of Nodes in Cluster") + 
  ylab("Time To Recover (ms)")+
  #geom_smooth(method="lm", se=FALSE, size=.75)+
  #annotate(geom="text", label=paste("R=", sprintf("%.3f", cor(bench$size, bench$time))))+
  #annotate(geom="text", label=paste("R^2=", sprintf("%.3f", cor(bench$size, bench$time)^2)))+
  theme_bw(base_family = "Times New Roman") +
  theme(plot.title = element_text(hjust = 0.5))

