from abc import abstractmethod
from typing import Any
from unicodedata import name

class Abstract_Test_Class:
    @abstractmethod
    def hello_name(self, name:Any):
        pass

# documentation, replacing test to force change, more text, another change.
class Concrete_Test_Class(Abstract_Test_Class):
    def hello_name(self, name:str):
        print(f"Hello, {name}")
