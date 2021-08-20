#-*= coding: utf-8 -*-
from ctypes import *
import ctypes
import datetime
from Crypto.Hash import CMAC
from Crypto.Cipher import AES
import struct
import os
from dotenv import load_dotenv
#import nfc_connect
#import subprocess

class go_string(Structure):
    _fields_ = [("p", c_char_p),("n", c_int)]
#GO言語用string変換

def GoString(s):
    u=s.encode('utf-8')
    return go_string(c_char_p(u), len(u))


load_dotenv('.env')
SECRET_KEY = os.environ["SECRET_KEY"]
API_TOKEN = os.environ["API_TOKEN"]
UUID = os.environ["UUID"]


# if __name__ == '__host__':
#     while True:
#         ID = nfc_connect.detect()
#         subprocess.call("aplay ic.wav" ,shell=True)
#         if ID == b'01010112be1aff08' or b'01407fc7d137b660':
message = int(datetime.datetime.now().timestamp()).to_bytes(4, 'little', signed=False)[1:4]
cmac = CMAC.new(bytes.fromhex(SECRET_KEY), ciphermod=AES)
cmac.update(message)
SIGN = cmac.hexdigest()
#パラメータの準備
lib = cdll.LoadLibrary("./export/export.so")
lib.operation.restype=c_char_p
lib.operation(GoString(SIGN),GoString(API_TOKEN),GoString(UUID),)