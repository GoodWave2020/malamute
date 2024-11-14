# malamute
malamuteは言語環境に依存せずにgit hooksを設定できるcliツールです。

## 使用方法
1. git管理されているプロジェクトの直下にシングルバイナリを配置する。

    ※windows環境にて`go build -o malamute.exe`を実行したとして進めます。

2. バイナリを実行する。

    下記のように出力されたら成功
    ```
    PS C:~~~> ./malamute.exe init
    Set Git hooks path to .malamute
    Malamute setup completed successfully.
    ```

3. .malamute配下でgit hooksの変更、修正が可能となります。

    ※.malamuteが既にgit管理されている場合もinitでhooksPathが変更されます。

## 機能

### init

initすると`.malamute`ディレクトリが作成され、core.hooksPathの指定先を`.malamute`に変更します。

以降はhooksは.malamuteディレクトリを参照します。

初回実行ではpre-commitとpre-pushのbashファイルが作成されます。

既に.malamuteディレクトリが存在する場合は、core.hooksPathの変更のみ行います。

他hooksを追加したい場合には.malamuteディレクトリに追加してください。

### reset

malamute initした際の設定を破棄します。

## ビルド方法

### linux
`go build -o malamute`

### windows
`go build -o malamute.exe`

