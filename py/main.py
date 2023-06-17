import speech_recognition as sr
import json
from flask import Flask

app = Flask(__name__)

def perintah():
    mendengar = sr.Recognizer()
    with sr.Microphone() as source:
        print('Mendengarkan....')
        suara = mendengar.listen(source, phrase_time_limit=60)
        try:
            print('Diterima...')
            dengar = mendengar.recognize_google(suara, language='id-ID')
            print(dengar)
        except:
            dengar = 'Suara tidak dikenali'
        return dengar

def run_voice():
    Layanan = perintah()
    print(Layanan)
    data = {
        'message': Layanan,
        'status': 'success'
    }
    return json.dumps(data)

@app.route('/')
def run():
    return run_voice()

if __name__ == '__main__':
    app.run(port=9090)
