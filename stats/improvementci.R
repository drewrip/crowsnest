library(matrixStats)
library(xtable)
dinghy <- read.csv("/home/drew/go/src/github.com/drewrip/dinghy/data/availsim/dinghytest.csv")
raft <- read.csv("/home/drew/go/src/github.com/drewrip/dinghy/data/availsim/rafttest.csv")

datize <- function(x){
  datamding <- data.matrix(x)
  trials <- nrow(datamding)
  sdsding <- colSds(datamding)
  meansding <- colMeans(datamding)
  cis <- qt(.99,trials-1)*sdsding/sqrt(trials)
  return(data.frame("size"=seq(3, 51, 2), "mean"=meansding, "sd"=sdsding, "ci"=cis, "trials"=trials))
}

dData <- datize(dinghy)
rData <- datize(raft)

rownames(dData) <- NULL
rownames(rData) <- NULL

dData
rData

diffmeans <- rData$mean - dData$mean
diffstds <- sqrt((rData$sd^2/(rData$trials)) + (dData$sd^2/(dData$trials)))
cisdif = qt(.99,(rData$trials-1)+(dData$trials))*diffstds
cisdif

finaldata <- data.frame("Cluster Size"=dData$size, "Vanilla Raft"=rData$mean, "Dinghy Raft"=dData$mean, "Improvement with Dinghy"=sprintf("%.2fpm%.2f",diffmeans,cisdif)) 

finaldata

xtable(finaldata)
