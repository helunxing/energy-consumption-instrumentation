import os


def pathGen(*args):
    paths, _ = os.path.split(os.path.abspath(__file__))
    return os.path.join(paths, *args)
