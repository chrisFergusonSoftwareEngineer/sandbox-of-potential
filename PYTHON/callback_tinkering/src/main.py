"""Tinkering with threadding in Python."""

# CAUTION FOR MANUAL RUNS: Each thread is independent of the main process, it does NOT SAFELY EXIT.  Nor can they be interrupted; they can only be stopped by killing the terminal instance.

# demonstrates multiple threads inteleaving with the main thread.

import time
from subscribe_looper import SubscribeLooper


def main_routine():
    print("starting")
    symbol = SubscribeLooper(["!", "@", "#", "$", "%"], printScrabbleTile)
    number = SubscribeLooper(
        ["1", "2", "3", "4", "5", "6", "7", "8", "9", "0"], printScrabbleTile
    )
    alpha = SubscribeLooper(
        [
            "a",
            "b",
            "c",
            "d",
            "e",
            "f",
            "g",
            "h",
            "i",
            "j",
            "k",
            "l",
            "m",
            "n",
            "o",
            "p",
            "q",
            "r",
            "s",
            "t",
            "u",
            "v",
            "w",
            "x",
            "y",
            "z",
        ],
        printScrabbleTile,
    )

    symbol.start()

    number.start()

    alpha.start()

    while True:
        print("======")
        if symbol.is_alive():
            print("Symbol healthy.")
        if number.is_alive():
            print("Number healthy.")
        if alpha.is_alive():
            print("Alpha healthy.")
        print("======")
        time.sleep(2)


def printScrabbleTile(tile):
    print(" [ %s ] " % tile)


if __name__ == "__main__":  # pragma: no cover
    main_routine()
