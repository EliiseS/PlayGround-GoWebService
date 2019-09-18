from flask import Flask, render_template, redirect, url_for, request
from models.student import Student
app = Flask(__name__)
students = []


@app.route("/", methods=["GET", "POST"])
def students_page():
    if request.method == "POST":
        new_student_id = request.form.get("student-id", "")
        new_student_name = request.form.get("name", "")
        new_student_last_name = request.form.get("last-name", "")

        new_student = Student(new_student_name, new_student_id)
        students.append(new_student)
        save_file(new_student_name)
        return redirect(url_for("students_page"))
    read_file()
    return render_template("index.html", students=students)


def save_file(student_name: str):
    try:
        f = open("students.txt", "a")
        f.write(student_name + "\n")
        f.close()
    except:
        print("Could not save file")


def read_file():
    try:
        f = open("students.txt", "r")
        for student_name in f.readlines():
            students.append(Student(student_name))

        f.close()
    except:
        print("Could not read file")


if __name__ == "__main__":
    app.run()
