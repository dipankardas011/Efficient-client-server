FROM python:3.10.7-bullseye as prod

WORKDIR /app

COPY . .

RUN pip install -r requirements.txt

CMD ["python", "main.py"]

EXPOSE 5000