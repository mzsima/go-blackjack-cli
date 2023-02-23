# ブラックジャック

<font size="20">🃏</font>

（このreadmeはAIが作成しています。）

## 目次
1. 概要
2. プチ解説
3. インストール
4. 使い方

## 概要

アリス: 「こんにちは、今日は私たちが作ったプロジェクトの概要を説明します。このプロジェクトは、ブラックジャックのゲームをコマンドライン上で遊べるようにするためのものです。」

ボブ: 「ブラックジャックとは、カジノで人気のあるトランプゲームです。」

アリス: 「そうです。プレイヤーとディーラーがそれぞれカードを引き、手札の合計値が21に近い方が勝ちというルールです。ぜひたくさんの方にプレイしていただきたいと思います。」

## プチ解説

アリス「ここのシャッフルアルゴリズムって使ったことある？」

ボブ「シャッフルアルゴリズム？ 何それ？」

アリスはシャッフルアルゴリズムについて教えることにした。

アリス「カードをシャッフルするアルゴリズムはいろいろあるけど、ここではFisher–Yatesアルゴリズムを使ってるわ。」

ボブ「Fisher-Yatesアルゴリズムって何？」

アリスはパソコンを起動して、コードを開いた。

アリス「Fisher-Yatesアルゴリズムは、リストをランダムにシャッフルするアルゴリズムの一つよ。コードはこんな感じ。」

```
func shuffleDeck() {
    for i := range deck {
        j := rand.Intn(i + 1)
        deck[i], deck[j] = deck[j], deck[i]
    }
}
```

アリスはコードを説明しながら進めた。

アリス「i はデッキのインデックスよ。0から順番に見ていくわけ。で、j は 0 から i までの範囲でランダムに選ぶのよ。」

ボブ「ランダムっていうのは、どうやってやるの？」

アリス「こうやって、最初にシード値をセットするの。」

```
rand.Seed(time.Now().UnixNano())
```


アリス「Seed() は疑似乱数ジェネレータにシード値をセットするの。シード値を固定してしまうと、毎回同じシャッフル結果になっちゃうから、現在時刻をシード値として使って、ランダム性を出してるのよ。」

ボブ「なるほど、時間をシード値に使うのはよくやるんだね。」

アリス「そうよ。他にも、ユーザーがマウスを動かした回数や、環境変数を使った方法もあるわ。ただ、本来の目的がランダム性を出すことだから、それに合ったものを使うのが重要よ。」

ボブはアリスに感心しながら、シャッフルアルゴリズムについての理解を深めていった。


## インストール
このプロジェクトを実行するには、Go言語の1.9以上が必要です。以下のコマンドを実行して、プロジェクトをクローンしてください。

```sh
git clone https://github.com/mzsima/go-blackjack-cli.git
```

クローンしたディレクトリに移動し、以下のコマンドを実行して、プログラムを実行してください。

```sh
go run main.go
```

以上で、ゲームをプレイすることができます。

また、ゲームのルールやロジックを改善するために、コードを編集することができます。編集したコードを実行する場合は、以下のコマンドを実行して、実行可能ファイルを生成してください。

```sh
go build main.go
```

生成されたバイナリファイルを実行することで、プログラムを実行することができます。

## 使い方
まずは、ターミナルでソースコードがあるディレクトリに移動してください。次に、以下のコマンドを実行してください。

```go
go run main.go
```

これで、ブラックジャックのゲームが始まります。ゲームの流れは、以下の通りです。

- ディーラーとプレイヤーに2枚のカードが配られます。
- ディーラーの1枚目のカードがオープンします。
- プレイヤーは、自分の手札を確認してから、追加のカードを引くか、スタンドするかを選択します。
- ディーラーは、自分の手札が17以上になるまで、カードを引き続けます。
- ディーラーとプレイヤーの手札を比較して、勝者を決定します。

