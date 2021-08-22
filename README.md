# go-sesame3
セサミ３のAPIをたたくクライアントアプリ
cgoを使ってpythonからgolangを呼び出しています。

## 開発環境
wsl2上で
goのバージョンは`go version go1.16.6 linux/amd64`
pythonのバージョンは`Python 3.8.10`

使用する外部モジュールは
- nfcpy
- python-dotenv
- pycryptodome

使用機器は
- Raspberry Pi 3 Model B+
- NFCリーダライタ SONY PaSoRi RC-S380
- [スピーカー　８Ω８Ｗ: パーツ一般 秋月電子通商\-電子部品・ネット通販](https://akizukidenshi.com/catalog/g/gP-03285/)
- [３．５ｍｍステレオミニプラグ⇔スクリュー端子台: パーツ一般 秋月電子通商\-電子部品・ネット通販](https://akizukidenshi.com/catalog/g/gC-08853/)
- [ＰＡＭ８０１２使用２ワットＤ級アンプモジュール: 組立キット\(モジュール\) 秋月電子通商\-電子部品・ネット通販](https://akizukidenshi.com/catalog/g/gK-08217/)

## ビルド後ディレクトリ
```bash
.
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

使用方法
上記環境を整えたうえで`export`配下で
```bash
$ go build -buildmode=c-shared -o export.so
```
を実行してください。
export配下にファイルが出力されます。
デフォルトでは相対パスで一つ上の階層の.envファイルから環境変数を参照しています。
一つ上の階層に.envファイルを作成してSECRET＿KEY、API＿TOKEN、UUID、を登録します。
ANDOROID、SUICAはそれぞれFelicaのIDmなので、適当な数値を入れておきます。
`go-sesame3`直下で
```bash
sudo python3 main.py
```
でプログラムが起動します。
この状態でカードリーダーに登録したいカードをかざすとコンソール上にバイト列を返します。
数値部分のみを.envのANDOROIDもしくはSUICAに格納して保存すると、次回起動時に利用することができます。
