class HelloWorld:

    name = "stranger"

    def setName(self, inputName):
        if inputName != "":
            self.name = inputName
    
    def helloWorld(self):
        print(f"Hello, {self.name}!")
