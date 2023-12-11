#!/usr/bin/env python3


def part1(nodes: dict, directions: list) -> int:
    count = 0
    current_node = "AAA"

    while current_node != "ZZZ":
        for dir in directions:
            current_node = nodes[current_node][dir]
            count += 1
            if current_node == "ZZZ":
                break

    return count


def all_keys_end_with_z(l: list) -> bool:
    for n in l:
        for k, _ in n.items():
            if not k.endswith("Z"):
                return False

    return True


def part2(nodes: dict, directions: list) -> int:
    count = 0
    current_nodes = [{k: nodes[k]} for k in nodes if k.endswith("A")]

    # Part 2
    while not all_keys_end_with_z(current_nodes):
        for dir in directions:
            for i in range(0, len(current_nodes)):
                node = current_nodes[i]
                for k in node:
                    dest = nodes[k][dir]
                    current_nodes[i] = {dest: nodes[dest]}

            count += 1

            print(count, current_nodes)

            if all_keys_end_with_z(current_nodes):
                break

    return count


if __name__ == "__main__":
    with open("./input", "r") as f:
        input = f.read().splitlines()

    nodes = {}

    for i in range(2, len(input)):
        name = input[i].split(" = ")[0]
        left = input[i].split(" = ")[1].strip("(").strip(")").split(",")[0].strip()
        right = input[i].split(" = ")[1].strip("(").strip(")").split(",")[1].strip()
        nodes[name] = {"L": left, "R": right}

    directions = [x for x in input[0]]

    # print("Part 1:", part1(nodes, directions))
    print("Part 2:", part2(nodes, directions))
