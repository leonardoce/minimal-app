FROM python:alpine

WORKDIR /app
RUN pip install flask
COPY app.py ./app.py
CMD python -m flask run --host=0.0.0.0
