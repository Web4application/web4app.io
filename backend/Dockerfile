FROM python:3.11
WORKDIR /code
COPY ./app /code/app
RUN pip install fastapi uvicorn python-multipart
CMD ["uvicorn", "app.main:app", "--host", "0.0.0.0", "--port", "8000"]
