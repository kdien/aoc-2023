#!/usr/bin/env python3


def find_diff(l: list[int], n: list[int]):
    s = set(l)
    if len(s) == 1 and sum(s) == 0:
        return

    diffs = []
    for i in range(1, len(l)):
        diffs.append(l[i] - l[i - 1])

    n.append(diffs[-1])
    find_diff(diffs, n)


if __name__ == "__main__":
    with open("./input", "r") as f:
        input = [list(map(int, x.split())) for x in f.readlines()]

    next_val_lists = []
    for l in input:
        next_val_list = []
        find_diff(l, next_val_list)
        next_val_lists.append(next_val_list)

    totalPartOne = 0
    for i in range(0, len(input)):
        totalPartOne += input[i][-1] + sum(next_val_lists[i])

    print("Part 1:", totalPartOne)
