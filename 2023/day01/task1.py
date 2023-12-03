import sys


def task1():
    if len(sys.argv) == 2:
        try:
            inf = open(sys.argv[1])
        except OSError:
            print("Could not open file")
            exit(1)
    else:
        inf = sys.stdin

    sum = 0
    first = 0
    last = 0

    for line in inf:
        line = line.rstrip('\n')
        for char in line:
            if char.isnumeric():
                first = int(char)
                break
        for char in reversed(line):
            if char.isnumeric():
                last = int(char)
                break

        sum += first * 10 + last
        # print(first * 10 + last)

    print(sum)

    if inf is not sys.stdin:
        inf.close()


if __name__ == "__main__":
    task1()
