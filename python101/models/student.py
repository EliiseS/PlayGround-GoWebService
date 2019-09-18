class Student():
    school_name = "Sprinfield Elem"

    def __init__(self, name: str, id: str = "1"):
        self.name: str = name
        self.student_id: str = id

    def __str__(self):
        return "Student " + self.name

    def get_name_capitalized(self):
        return self.name.capitalize()
