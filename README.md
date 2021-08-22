# go-sesame3
セサミ３のAPIをたたくクライアントアプリ
開発環境はwsl2上で
goのバージョンは`go version go1.16.6 linux/amd64`
pythonのバージョンは`Python 3.8.10`
使用する外部モジュールは
- nfcpy
- python-dotenv
- pycryptodome

ビルド後のディレクトリは
```bash
.
├── LICENSE
├── README.md
├── export
│   ├── export.go
│   ├── export.h
│   ├── export.so
│   └── go.mod
├── main.py
├── nfc_connect.py
└── notify.wav
```
