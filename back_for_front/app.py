from flask import Flask, request
import psycopg2


app = Flask(__name__)

try:    
    conn = psycopg2.connect(dbname='mars', user='postgres', password='111111', host='localhost', port = '5432')
    print('Connection established')

except psycopg2.Error as e:
    print(e)
    print('Can`t establish connection to database')

@app.route('/post', methods=['POST'])
def post_func():    

    req = request
    print('-------')
    print(req)
    for file_name in req.files.keys():
        file = req.files[file_name]
        #print(file_path)


    with open(file_name, 'wb') as f:
        f.write(file.read())
    print(req.form['text'])
    print(file_name)
    print(req.form['name'])
    # Берем данные из запроса (отчет сначала в переменную, потом в бд, а файл сразу в папку) 
    # Берем переменную отчета и пишем в бд: conn.exec("INSERT запрос")
    try:
        cur = conn.cursor()
        res = cur.execute("INSERT INTO public.report(report_text, path, name) VALUES (%s, %s, %s);", 
        (req.form['text'], file_name, req.form['name']))
        print(res)
        cur.close()
    except:
        print('Can`t insert into database')

    conn.close()
    return 'OK'




if __name__ == '__main__':
    app.run(host='127.0.0.1', port=5000)