from flask import Flask
import os

app = Flask(__name__)
environment = os.getenv('ENVIRONMENT', 'unknown')
color = os.getenv('COLOR', 'black')

@app.route("/")
def hello_world():
    return f"<p style='color: {color}; text-weight: 2000'>Environment: {environment}</p>"
