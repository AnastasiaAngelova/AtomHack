from flask import Flask, request, jsonify
import psycopg2
import requests


app = Flask(__name__)

try:    
    conn = psycopg2.connect(dbname='mars', user='postgres', password='111111', host='localhost', port = '5432')
    conn.autocommit = True
    print('Connection established')

except psycopg2.Error as e:
    print(e)
    print('Can`t establish connection to database')

@app.route('/report', methods=['POST'])
def new_report():    

    req = request
    for file_name in req.files.keys():
        file = req.files[file_name]

    file_name += '_' + str(hash(req.form['text'])) 
    with open(file_name, 'wb') as f:
        f.write(file.read())
    # Берем данные из запроса (отчет сначала в переменную, потом в бд, а файл сразу в папку) 
    # Берем переменную отчета и пишем в бд: conn.exec("INSERT запрос")
    try:
        cur = conn.cursor()
        res = cur.execute("INSERT INTO public.report(report_text, path, name) VALUES (%s, %s, %s);", 
        (req.form['text'], file_name, req.form['name']))
        cur.close()
    except:
        print('Can`t insert into database')

    # отправляем другому серверу информационное сообщение
    response = requests.post("http://127.0.0.1:4000/inform", data={'event':'new_report'})

    return 'OK'


@app.route('/reports', methods=['GET'])
def get_reports():    

    try:
        cur = conn.cursor()
        cur.execute("select * from public.report ;")
        res = cur.fetchall()
        
        cur.close()
    except:
        print('Can`t select from database')


    return jsonify({'result': res})

@app.route('/report', methods=['GET'])
def get_report_by_id():   
    print(request.args) 

    try:
        cur = conn.cursor()
        cur.execute("select * from public.report where id = %s ;", (request.args['id'],))
        res = cur.fetchall()
        print(id)
        print( res ) 
        cur.close()
    except:
        print('Can`t select from database')


    return jsonify({'result': res})



if __name__ == '__main__':
    app.run(host='127.0.0.1', port=5000)