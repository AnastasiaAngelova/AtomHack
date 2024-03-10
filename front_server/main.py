from flask import Flask, request
import psycopg2

app = Flask(__name__)
try:
    conn = psycopg2.connect(dbname='test', user='postgres', password='secret', host='host')
except:
    print('Can`t establish connection to database')

@app.route('/post', methods=['POST'])
def post_func():
    # Берем данные из запроса (отчет сначала в переменную, потом в бд, а файл сразу в папку)
    # Берем переменную отчета и пишем в бд: conn.exec("INSERT запрос")

@app.route('/chernovik', methods=['GET'])
def post_func():
    # Вернем из бд черновик

if __name__ == '__main__':

    app.run(host='127.0.0.1', port=5000)
