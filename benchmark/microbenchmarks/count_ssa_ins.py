import sys


def read_in(path):
    funcs = []
    currentfunc = []
    with open(path) as f:
        for line in f:
            if line.strip() == "":
                funcs.append(currentfunc)
                currentfunc = []
            else:
                currentfunc.append(line)
    if len(currentfunc) != 0:
        funcs.append(currentfunc)
    return funcs


def is_assertion(func):
    x = func[0].split('.')
    x = x[-1].strip()
    if (x.startswith('spec_')) or x == "assertEq" or x == "tryCast":
        return True
    else:
        return False


def is_code(line):
    return not (line.startswith('#') or line.strip() == "")

def printfunc(func):
    for line in func:
        print(line, end="")
    print()


def print_normal_funcs(funcs):
    for x in funcs:
        if not is_assertion(x):
            printfunc(x)

def print_code(funcs):
    for x in funcs:
        for line in x:
            if is_code(line):
                print(line, end="")

def main(str):
    x = read_in(str)
    print_code(x)


if __name__ == "__main__":
    fname = sys.argv[1]
    main(fname)