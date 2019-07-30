# 課題1 画像変換コマンドを作ろう

- 次の仕様を満たすコマンドを作って下さい
    - ディレクトリを指定する
    - 指定したディレクトリ以下のJPGファイルをPNGに変換（デフォルト）
    - ディレクトリ以下は再帰的に処理する
    - 変換前と変換後の画像形式を指定できる（オプション）
- 以下を満たすように開発してください
    - mainパッケージと分離する
    - 自作パッケージと標準パッケージと準標準パッケージのみ使う
    - 準標準パッケージ：golang.org/x以下のパッケージ
    - ユーザ定義型を作ってみる
    - GoDocを生成してみる

# imgconv

## Usage

```
Usage:
  imgconv [-in_ext] [-out_ext] [-leave] DIR
Arguments:
  -in_ext  input extension (jpg, png, gif)  [default: jpg]
  -out_ext output extension (jpg, png, gif) [default: png]
  -leave   whether to leave input           [default: fasle]
```

## Development

### Build

```
make build
```

### Test

```
make test
```

### Clean

```
make clean
```