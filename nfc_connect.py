import nfc
import binascii

class CardReader():
    '''docstring
    ・カードリーダークラス
    ・Felicaを検知したらIDmを返す
    '''
    def __init__(self):
        self.idm = 0
    def on_startup(self,targets):
        for target in targets:
            target.sensef_req = bytearray.fromhex("0000030000")
        return targets
    def on_connect(self,tag):
        print("Detected!!")
        self.idm = binascii.hexlify(tag.idm)

    def read_id(self):
        clf = nfc.ContactlessFrontend('usb')
        print("Waiting Felica...")
        try:
            clf.connect(rdwr = {'targets':['212F'],'on-startup':self.on_startup,'on-connect':self.on_connect})
            print(str(self.idm))
            clf.close()
        except KeyboardInterrupt:
            print("KeyboardInterrupt!!")
