import numpy as np
from scipy.stats import gaussian_kde
from numpy import linspace
import matplotlib.pyplot as plt


def make_plot(data, label):
    kde = gaussian_kde(data)
    dist_space = linspace(min(data), max(data), 1000)
    plt.plot(dist_space, kde(dist_space), label=label)


p1 = [float(input()) for _ in range(1000)]
p2 = [float(input()) for _ in range(1000)]
p3 = [float(input()) for _ in range(1000)]

print("mean1:", np.mean(p1))
print("mean2:", np.mean(p2))
print("mean3:", np.mean(p3))

make_plot(p1, "inaccurate estimate")
make_plot(p2, "accurate estimate")
make_plot(p3, "with updates")

plt.legend()
plt.show()
