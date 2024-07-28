import sys
import matplotlib.pyplot as plt
import numpy as np

tmpFile = "/dev/shm/SchmottkyTmp"

x_range = float(sys.argv[1])
y_range = float(sys.argv[2])

xs = []
ys = []

def process(chunk):
    global xs, ys
    floats = np.frombuffer(chunk, dtype=np.float64)
    xs.append(floats[0])
    ys.append(floats[1])

# load data from file
with open(tmpFile, 'rb') as f:
    while True:
        chunk = f.read(16)  
        if not chunk:
            break
        process(chunk)

# plot
figure, axes = plt.subplots()
axes.set_aspect(1)
axes.get_xaxis().set_ticks([])
axes.get_yaxis().set_ticks([])
plt.plot(xs, ys, color='black', linewidth=0.5)

axes.set_xlim([-x_range, x_range])
axes.set_ylim([-y_range, y_range])
axes.set_aspect(1) 
plt.show()