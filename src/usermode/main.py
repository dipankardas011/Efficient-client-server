from flask import Flask, request as req
import requests, os, socket


app = Flask(__name__)


class User:
  def __init__(self, request = '', response = ''):
    self.request = request
    self.response = response


def ts(s, str):
   s.send(str.encode())
   data = ''
   data = s.recv(1024).decode()
   print (data)

@app.route('/request', methods=['POST'])
def sendMessage():
  usr = User(req.form['message'])

  s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
  host ="127.0.0.1"
  port =8000 # for the client's server socket listener
  s.connect((host,port))

  r = input('enter')
  ts(s, r)


  s.close ()

  url = 'http://localhost:8080'
  params = {'message': usr.request}
  resp = requests.get(url, params=params)
  return {
    "Response": resp,
    "Request": usr.request,
    "Status": 200
  }


if __name__ == "__main__":
  port = int(os.environ.get('PORT', 5000))

  usr = User()

  app.run(debug=True, host='0.0.0.0', port=port)