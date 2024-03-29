"""Subscription looper class."""

import time
from threading import Thread


class SubscribeLooper(Thread):
    """Subscription loop runner."""

    def __init__(self, scrabble: list, callback):
        self.scrabble = scrabble
        self.callback = callback

        Thread.__init__(self)

    def run(self):
        while True:
            for item in self.scrabble:
                self.callback(item)
                time.sleep(0.25)
            time.sleep(2)
