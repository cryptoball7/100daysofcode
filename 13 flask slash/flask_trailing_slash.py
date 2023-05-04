# https://flask.palletsprojects.com/en/2.2.x/quickstart/
# Unique URLS / Redirection Behavior

from flask import Flask

app = Flask(__name__)

# This one acts like directory

@app.route("/home/")
def home():
	return "The home page"

# This behaves like a file

@app.route("/index.html")
def index():
	return "Index file"