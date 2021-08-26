from ctypes import *
import ctypes
import datetime
from Crypto.Hash import CMAC
from Crypto.Cipher import AES
import struct
from time import sleep
import os
from dotenv import load_dotenv
import subprocess
from nfcreader import CardReader

<<<<<<< HEAD
# 環境変数の読み込みにはpython-dotenvを使用している
# 事故が起きないように.envファイルはhomeディレクトリもしくは相対パスで指定する
#HOME = os.environ['HOME']
# load_dotenv(HOME+'.env')
load_dotenv('../.env')
=======
class go_string(Structure):
    _fields_ = [("p", c_char_p),("n", c_int)]
#GO言語用string変換

def GoString(s):
    u=s.encode('utf-8')
    return go_string(c_char_p(u), len(u))


print(os.environ['HOME'])
HOME = os.environ['HOME']
load_dotenv(HOME+'/.env')
>>>>>>> 87391ae (update)
SECRET_KEY = os.environ["SECRET_KEY"]
API_TOKEN = os.environ["API_TOKEN"]
UUID = os.environ["UUID"]
ANDROIDO = os.environ["ANDROIDO"].encode()
SUICA = os.environ["SUICA"].encode()


class MySesame3:
    '''docstring
    ・セサミ3のクラス、libはgoでビルドしたバイナリファイルを読み込んでいる
    ・API_TOKENはこちら(https://dash.candyhouse.co/login)で取得する
    ・UUIDはアプリに記載されている対象のセサミ3のUUID
    ・SECRET_KEYはアプリを立ち上げて「鍵のシェア（オーナー）」で生成されるQRコードを読み込んで得られる文字列
    ・施錠/解錠の際はSECRET_KEYとタイムスタンプをAES-CMACによって暗号化する必要がある(encryptmyKey)
    ・作成した暗号キーをsignとしてAPI_TOKEN,UUIDとともにPOSTすることで施錠解錠ができる(lockOrunlock)
    ・現在の鍵の状態(施錠中/解除中)はgolang側で判断する
    ・施錠中であれば解錠コマンドを、解錠中であれば施錠コマンドを打ち込む
    '''
    def __init__(self):
        self.lib = cdll.LoadLibrary("./export/export.so")
        self.key = SECRET_KEY
        self.api = API_TOKEN
        self.uuid = UUID
        self.sign = ""

    def encyptmyKey(self):
        timestamp = int(datetime.datetime.now().timestamp()).to_bytes(4, 'little', signed=False)[1:4]
        cmac = CMAC.new(bytes.fromhex(self.key), ciphermod=AES)
        cmac.update(timestamp)
        self.sign =  cmac.hexdigest()

    def lockOrunlock(self):
        self.lib.executeSesame3.restype=c_char_p
        self.lib.executeSesame3(self.sign.encode('utf-8'),self.api.encode('utf-8'),self.uuid.encode('utf-8'))

def ismyID(id):
    return bool(id==ANDROIDO or id ==SUICA)

if __name__ == '__main__':

     mySesame3 = MySesame3()
     try:
        while True:
            # nfcpyによるNFC入力待機
            myreader = CardReader()
            myreader.read_id()
            detectedID = myreader.idm
            # NFCの入力を検知したらスピーカーから通知音を出す
            subprocess.call("aplay notify.wav" ,shell=True)
            if ismyID(detectedID):
                # セサミ3インスタンスの作成
                # secret_keyを暗号化
                mySesame3.encyptmyKey()
                # 施錠と解錠の実行
                mySesame3.lockOrunlock()
                detectedID = 0
            sleep(2)

     except KeyboardInterrupt:
        print("KeyboardInterrupt!!")
       
