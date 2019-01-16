import sys
import numpy
import matplotlib.pyplot as plt

data = numpy.loadtxt(sys.argv[1], delimiter="\n")

print(data)

plt.hist(data, 100)
plt.show()