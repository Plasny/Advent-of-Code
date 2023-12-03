import sys


def task2():
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
        # print(line)

        for i in range(0, len(line)):
            if line[i].isnumeric():
                first = int(line[i])
                break
            if (tmp := check_for_word(i, line)):
                first = tmp
                break

        for i in range(len(line) - 1, -1, -1):
            if line[i].isnumeric():
                last = int(line[i])
                break
            if (tmp := check_for_word(i, line)):
                last = tmp
                break

        sum += first * 10 + last
        # print(first * 10 + last)

    print(sum)

    if inf is not sys.stdin:
        inf.close()


def check_for_word(i, line):
    d = {
        "one": 1,
        "two": 2,
        "three": 3,
        "four": 4,
        "five": 5,
        "six": 6,
        "seven": 7,
        "eight": 8,
        "nine": 9,
    }

    for num in d:
        if line[i:i + len(num)] == num:
            return d[num]

    return None


if __name__ == "__main__":
    task2()
