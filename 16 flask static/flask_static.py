from flask import Flask

app = Flask(__name__)

from flask import url_for

with app.test_request_context():
	print(url_for('static', filename='static_test.html'))

