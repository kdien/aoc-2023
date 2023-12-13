#!/usr/bin/env python3

from functools import cmp_to_key


def get_hand_strength(hand: str) -> int:
    """
    High card -> 0
    One pair -> 1
    Two pairs -> 2
    Three of a kind -> 3
    Full house -> 4
    Four of a kind -> 5
    Five of a kind -> 6
    """
    cards = [x for x in hand]
    match len(set(cards)):
        case 1:
            return 6
        case 2:
            cards.sort()
            if (
                cards[0] == cards[1] and cards[1] == cards[2] and cards[2] == cards[3]
            ) or (
                cards[1] == cards[2] and cards[2] == cards[3] and cards[3] == cards[4]
            ):
                # four of a kind
                return 5
            else:
                # full house
                return 4
        case 3:
            cards.sort()
            if (cards[0] == cards[1] and cards[2] == cards[3]) or (
                cards[1] == cards[2]
                and cards[3] == cards[4]
                or (cards[0] == cards[1] and cards[3] == cards[4])
            ):
                # two pairs
                return 2
            else:
                # three of a kind
                return 3
        case 4:
            return 1
        case _:
            return 0


card_labels = ["2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"]


def hand_cmp(a: dict, b: dict) -> int:
    if a["strength"] == b["strength"]:
        for i in range(0, len(a["hand"])):
            if a["hand"][i] == b["hand"][i]:
                continue
            if card_labels.index(a["hand"][i]) < card_labels.index(b["hand"][i]):
                return -1
            else:
                return 1

    return a["strength"] - b["strength"]


if __name__ == "__main__":
    with open("./input", "r") as file:
        input = file.read().splitlines()

    hands = []
    for line in input:
        hand = line.split()[0]
        bid = line.split()[1]
        hands.append(
            {"hand": hand, "bid": int(bid), "strength": get_hand_strength(hand)}
        )

    hands = sorted(hands, key=cmp_to_key(hand_cmp))

    total_part_1 = 0
    for i in range(0, len(hands)):
        total_part_1 += hands[i]["bid"] * (i + 1)

    for i, v in enumerate(hands):
        print(f"{v} - {v['bid']} * {i+1}")

    print("Part 1:", total_part_1)
