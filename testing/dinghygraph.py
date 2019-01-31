import sys
import numpy
import matplotlib.pyplot as plt

d1 = numpy.loadtxt(sys.argv[1], delimiter="\n")
d2 = numpy.loadtxt(sys.argv[2], delimiter="\n")
print(d1)
print(d2)

#plt.plot(d1, 100)
#plt.show()