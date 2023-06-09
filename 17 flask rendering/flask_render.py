# 100DaysOfCode CryptoBallX
# https://flask.palletsprojects.com/en/2.2.x/quickstart/

from flask import Flask

app = Flask(__name__)

from flask import render_template

# render_template(filename, *args)
# looks for filename in "templates" folder
# and renders it using args

@app.route("/hello/")
@app.route("/hello/<name>")
def hello(name=None):
	return render_template("hello.html", name=name)

