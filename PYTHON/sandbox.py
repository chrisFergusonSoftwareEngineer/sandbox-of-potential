import sys

from builtin.helloWorldClass import HelloWorld

trivialClass = HelloWorld()

if len(sys.argv) > 1:
    trivialClass.setName(sys.argv[1])

trivialClass.helloWorld()
