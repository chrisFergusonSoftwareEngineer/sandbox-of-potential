"""Quick test of using an Abstract Method in Python."""

from abc import abstractmethod
from typing import Any


class AbstractTestClass:  # pylint: disable=too-few-public-methods
    """Simple abstract class for testing this process."""

    @abstractmethod
    def hello_name(self, name: Any):
        """Simple abstract method to test implementation."""


# documentation, forced changes.
class ConcreteTestClass(AbstractTestClass):  # pylint: disable=too-few-public-methods
    """Simple concrete class testing abstract extenstion."""

    def hello_name(self, name: str):
        """Simple implementation to show how abstracts work."""
        print(f"Hello, {name}")
