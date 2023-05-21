# HTTP Methods
# https://flask.palletsprojects.com/en/2.2.x/quickstart/

from flask import Flask

app = Flask(__name__)

from flask import request

login_form = "<form method=\"post\" action=\"/login\"><input><br><input><br><input type=\"submit\"></form>"
login2_form = "<form method=\"post\" action=\"/login2\"><input><br><input><br><input type=\"submit\"></form>"

@app.route("/login", methods=["GET", "POST"])
def login():
	if request.method == "POST":
		return "logging in..."
	else:
		return login_form

@app.get("/login2")
def login2_get():
	return login2_form

@app.post("/login2")
def login2_post():
	return "Logging in..."