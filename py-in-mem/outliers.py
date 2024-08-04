"""Detect outliers"""
import numpy as np


def detect(data):
    #print("PY outliers", len(data))
    """Return indices where values more than 2 standard deviations from mean"""
    out = np.where(np.abs(data - data.mean()) > 2 * data.std())
    # np.where returns a tuple for each dimension, we want the 1st element
    #print("PY returning ", len(out[0], out[0]))
    return out[0]
