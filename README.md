# Dinghy: Horizontally Scaling Raft Clusters

#### _Abstract_

With the introduction of Raft, distributed consensus has become more widely available in the use of cluster design. Different applications of distributed systems often require differing configurations, however, as these clusters begin to increase in the number of participating nodes, they become increasingly less efficient at coming to consensus. Current methods of scaling distributed systems generally implement some variation of data sharding, batch processing, or message coalescing. We investigate however, utilizing dynamically set timeouts and heartbeat intervals to increase network throughput, and propose methods to more effectively handle a growing number of nodes in a cluster, while analyzing their ability to increase the horizontal scalability of a Raft cluster. 

 
 

[See the paper here](paper/main.pdf)
