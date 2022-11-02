
class geeks: 
    def __init__(self, name, roll): 
        self.name = name 
        self.roll = roll

class AccessorExceptions:
    def review_exceptions(self):
        list_of_things = None

        try:
            print(f"{list_of_things[0].foo_bar}")
        except Exception as unknown_none_access_exception:
            print(f"Unknown Acess Exception: {type(unknown_none_access_exception)}")

        list_of_things = []

        try:
            print(f"{list_of_things[0].foo_bar}")
        except Exception as unknown_empty_list:
            print(f"Empty List Exception: {type(unknown_empty_list)}")

        list_of_things = [None, None]

        try:
            print(f"{list_of_things[0].foo_bar}")
        except Exception as unknown_list_of_none:
            print(f"List of None Exception: {type(unknown_list_of_none)}")

        list_of_things = []
        list_of_things.append(geeks("ferg", 20))

        try:
            print(f"{list_of_things[0].name}")
        except Exception as unknown_valid_access:
            print(f"Valid Access Exception: {type(unknown_valid_access)}")

        try:
            print(f"{list_of_things[5].foo_bar}")
        except Exception as unknown_out_of_bounds:
            print(f"Out of Bounds Exception: {type(unknown_out_of_bounds)}")
