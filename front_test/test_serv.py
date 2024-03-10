import requests

url = 'http://127.0.0.1:5000/report'
data = {'name': 'John', 'text': 'text report space mars bla bla bla'}
files = {'test.txt': open('test.txt', 'rb')}

#response = requests.post(url, data=data, files=files)

#print(response)

response = requests.get('http://127.0.0.1:5000/report?id=15')

print(response)
