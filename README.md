# gin_tutorial

## install

### goのinstall

```
brew install go
```

## projectの作成

```
mkdir && cd [project_name]
```

### projectの初期化
```
go mod init [module_name]
```

* `go.mod`が作成される

### ginのinstall

```
go get -u github.com/gin-gonic/gin
```

### `main.go`で読み込む

```
import "github.com/gin-gonic/gin"
```

## HotReloadの有効化

### `air`のinstall<br />
https://github.com/air-verse/air

```
go install github.com/air-verse/air@latest
```

### `.air.toml`の作成

* ビルド設定などを記述する。
* このプロジェクトでは、デフォルトの記述のまま利用している
  * https://github.com/Yuta-Matsumoto999/go_gin_tutorial/blob/main/.air.toml

### 起動

```
air
```

## 静的解析ツールの導入

### `golangci-lint`のinstall

```
brew install golangci-lint@latest
```

* `go install` では、なぜか`v1`がinstallされてしまうため、`brew`経由でinstallを行った。

### 静的解析設定ファイルの追加

* `.golangci.yml`を作成
  * https://github.com/Yuta-Matsumoto999/go_gin_tutorial/blob/main/.golangci.yml
  * 内容については、下記ドキュメント参照
    * https://golangci-lint.run/usage/configuration/

### 静的解析実行

```
golangci-lint run
```
* `linters`で指定されたlinterが実行される

```
golangci-lint run --fix
```
* formatエラーの修正を実行する

### workflowの追加

* PRの作成で、`golangci-lint run`を実行するworkflowを追加
  * https://github.com/Yuta-Matsumoto999/go_gin_tutorial/blob/main/.github/workflows/analysis.yml
    * `reviewdog/action-golangci-lint@v2`actionを利用して、静的解析の実行とエラーをPRコメントに自動で追加するよう設定している。

## APIについて

* ３層アーキテクチャを採用している
  * `controller`
  * `service`
  * `repository`

※ 今後、DDDを意識したアーキテクチャに変更する可能性あり

