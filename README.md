# clubhook-backend

https://clubhook-backend.herokuapp.com/

## Requirement

* Docker

* docker-compose

## Usage

ローカルでは、docker-composeでコンテナを立ち上げることでサーバが起動します。

```
$ docker-compose build
$ docker-compose up
```

#### 注意

初回ビルド時などに、appコンテナがpostgresコンテナよりも先に起動し、
データベースとのコネクションエラーが起きることがあります。

その場合は一度コンテナを`Ctrl+C`などで停止して、
もう一度`docker-compose up`でコンテナを起動してください。

## Contribution

VSCode Remote Containers向けの設定を用意していますので、
Visual Studio Codeで開発することをおすすめします。

### Visual Studio Codeで開発用コンテナを起動する方法

* Visual Studio Codeを起動後、ウインドウ左下端の緑の部分をクリック。

* メニューから「Remote-Containers:Open Folder in Container...」をクリック。

#### 注意

`docker-compose build`を行った後に開発用コンテナを起動しようとするとエラーが起きます。
その場合は`Rebuild Container`で開発用コンテナを再ビルドしてください。

このエラーを回避するため、開発中の動作確認はdocker-composeを使用せず、
Visual Studio Codeのターミナルから`go run main.go`でサーバを起動することをおすすめします。
