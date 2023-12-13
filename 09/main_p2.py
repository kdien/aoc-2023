#!/usr/bin/env python3


def find_diff(l: list[int], p: list[int]):
    s = set(l)
    if len(s) == 1 and sum(s) == 0:
        return

    diffs = []
    for i in range(1, len(l)):
        diffs.append(l[i] - l[i - 1])

    p.append(diffs[0])
    find_diff(diffs, p)


if __name__ == "__main__":
    with open("./input", "r") as f:
        input = [list(map(int, x.split())) for x in f.readlines()]

    prev_val_lists = []
    for l in input:
        prev_val_list = []
        find_diff(l, prev_val_list)
        prev_val_lists.append(prev_val_list)

    total_part_2 = 0
    for i in range(0, len(input)):
        beginning_nums = [0]
        prev_val_list = prev_val_lists[i]

        for j in range(len(prev_val_list) - 2, -1, -1):
            beginning_nums.append(prev_val_list[j] - beginning_nums[-1])

        total_part_2 += input[i][0] - beginning_nums[-1]

    print("Part 2:", total_part_2)
