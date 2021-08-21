import nfc
import binascii

class CardReader():
    '''docstring
    ・カードリーダークラス
    ・Felicaを検知したらIDmを返す
    '''
    def on_connect(self,tag):
        print("Detected!!")
        self.idm = binascii.hexlify(tag.idm)
        return str(self.idm)
    def read_id(self):
        clf = nfc.ContactlessFrontend('usb')
        print("Waiting Felica...")
        try:
            clf.connect(rdwr = {'on-connect':self.on_connect})
        finally:
            clf.close()