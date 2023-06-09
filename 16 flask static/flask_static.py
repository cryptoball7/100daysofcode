# 100DaysOfCode CryptoBallX
# https://flask.palletsprojects.com/en/2.2.x/quickstart/

from flask import Flask

app = Flask(__name__)

from flask import url_for

with app.test_request_context():
	print(url_for('static', filename='static_test.html'))

