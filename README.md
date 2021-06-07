# clubhook-backend

https://clubhook-backend.herokuapp.com/

## Getting Started

Docker(docker-compose)で開発環境を用意する場合はgit cloneでソースを取得。

#### Clone Repository

```
$ git clone https://github.com/kougakusai/clubhook-backend.git
$ cd clubhook-backend
```

#### Docker(docker-compose)で実行する場合

```
$ docker-compose build
$ docker-compose up -d

$ docker-compose exec app go run main.go
# Access to "localhost:8000"

$ docker-compose exec app /bin/bash
$ go run main.go
# Access to "localhost:8000"
$ exit

$ docker-compose down
```

#### VSCodeでDevContainerを起動し開発する場合
VSCodeを起動後、ウインドウ左下端の緑の部分をクリック。
メニューから「Remote-Containers:Open Folder in Container...」をクリック。
git cloneしたディレクトリを選択。
