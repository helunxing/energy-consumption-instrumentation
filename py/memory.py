from memory_profiler import profile

# copy and paste
# https://towardsdatascience.com/profile-memory-consumption-of-python-functions-in-a-single-line-of-code-6403101db419


@profile
def my_func():
    a = [1] * (10 ** 6)
    b = [2] * (2 * 10 ** 7)
    c = [3] * (2 * 10 ** 8)
    handle()
    del b
    return a

def handle():
    a='degnb'
    print("sege")

if __name__ == '__main__':
    my_func()
