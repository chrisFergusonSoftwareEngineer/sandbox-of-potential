"""Simple module for hello world"""


class HelloWorld:
    """Simple class for Hello World"""

    name = "stranger"

    def set_name(self, inputName):
        if inputName != "":
            self.name = inputName

    def hello_world(self):
        print(f"Hello, {self.name}!")
