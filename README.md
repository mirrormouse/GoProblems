


# プログラミング初心者のためのGo初級問題集

provided by mirrormouse

| 番号 | 内容 | 概要 |
| ---- | ---- | ---- |
| Q00~Q09 | [初歩編](#初歩編) | HelloWorldからFizzBuzzまで |
| Q10~Q19 | [基礎編](#基礎編) | 配列、スライス、再帰関数 |
| Q20~Q29 | 初級文法編(未定) | 型、構造体、インタフェース |
| Q30~Q39 | 初級アルゴリズム編(未定) | ソート、貪欲法、動的計画法、高階関数 |
| Q40~Q49 | 初級実装編(未定) | 並行処理、例外、埋め込み |

## 使い方<br>

0. Goのコンパイラをダウンロード。
1. 画面右上にある緑色の「Code▼」を選択。
2. gitが使える方は、適当なフォルダにclone。そうでない方は、Download zipを選択し、適当なフォルダに展開。
3. cloneまたは展開したフォルダを適当なエディタで開く。メモ帳でも何とかなるが、VSCodeがおすすめ。
4. コマンドプロンプトやターミナルを開き、cdコマンドで"GoProblems"フォルダに移動。
5. questionフォルダ内のファイルを開き、問題に解答。
6. コマンドプロンプトまたはターミナルで"test"フォルダに移動後、"go test -run Test{問題番号}"（例えばQ00なら"go test -run Test00"。）を実行する。こうすることでコードが正しいかテストできる。


### テストについての補足

"go test -run Test0"とするとQ00からQ09まで、"go test -run Test1"とするとQ10からQ19までのテストを一度に実行できる。Test2以降も同様。また"go test -run Test"とすると全てのテストを一度に実行できる。<br>
PASSと表示されれば正解。FAILの場合、エラーが出ているか、出力に誤りがある。
テストの際に与えられる入力は、問題文に入力例として示されているもの以外にも存在し、入力例に示されていないテストのみで失敗している場合は、"Secret Case Failed"と表示される。


## 初歩編

<br>

### Q.00 Hello World!
以下と同じになるようにプログラムをQ00.goを書き換え、「使い方」の説明に従ってテストせよ。「PASS」と表示されれば成功である。以降の問題も同様にテストすること。


```Go
package problems

import "fmt"

func Hello() {
	fmt.Printf("Hello World\n")
}
```


### Q.01 平均の計算１
2つの整数x,yを入力として受け取り、この平均を出力として返す関数AverageをQ01.goに実装せよ。ただし小数点以下は絶対値が小さくなるように切り捨て、整数として出力すること。
以下に入出力の例を示す。テストではここに示した例以外の入力についてもテストが行われることに注意すること。

| 入力例 | 出力例 |
| ---- | ---- |
| 1, 3 | 2 |
| 5, 0 | 2 |
| 5, -8 | -1 |

### Q.02 if文の使用
3つの整数a,b,cを入力として受け取り、この最大値を出力として返す関数MaxをQ02.goに実装せよ。


| 入力例 | 出力例 |
| ---- | ---- |
| 3, 9, 7 | 9 |
| 2, 2, 2 | 2 |
| 0, -5, 3 | 3 |

### Q.03 for文の使用
整数xを受け取りxの2乗を出力として返す関数SquareをQ03.goに実装せよ。また正の整数nを受け取り、1からnまでの値の2乗の和を返す
関数SquareSumを、関数Squareを利用する形で実装せよ。テストはSquareSumに対してのみ行われる。


| 入力例 | 出力例 |
| ---- | ---- |
| 3 | 14 |
| 5 | 55 |
| 7 | 140 |

### Q.04 switch文の使用
正の整数nを受け取り、4で割った余りを英語表記した文字列を出力として返す関数Mod4をQ04.goに実装せよ。表記は小文字でも大文字でも構わない。

| 入力例 | 出力例 |
| ---- | ---- |
| 7 | three |
| 12 | Zero |
| 5 | ONE |

### Q.05 文字列結合演算子「+」
文字列sと文字列tを受け取り、s,tを結合したうえで末尾に"!"を付けた文字列を出力として返す関数StringConcatをQ05.goに実装せよ。

| 入力例 | 出力例 |
| ---- | ---- |
| Hello,World | HelloWorld! |
| Hello,Go | HelloGo! |
| Go,Program | GoProgram! |


### Q.06 累乗関数
整数xと自然数aを受け取り、xのa乗を返す関数mypowをQ06.goに実装せよ。なお、ここで自然数は0を含むものとする。

| 入力例 | 出力例 |
| ---- | ---- |
| 2,3 | 8 |
| 2,0 | 1 |
| -3,3 | -27 |

### Q.07 シンプルなFizzBuzz
正の整数nを受け取り、nが3の倍数でかつ15の倍数ではないなら"Fizz"、nが5の倍数でかつ15の倍数でないなら"Buzz"、nが15の倍数なら"FizzBuzz"、それ以外なら整数nを文字列に変換したものを出力として返す関数FizzBuzzSingleを実装せよ。
<details><summary>ヒント</summary><div>
数値xを文字列に変換したいときは、"strconv"をimportしたうえで以下のように書く。

```Go
s:=strconv.Itoa(i)
```

解答用ファイルの最初に"strconv"をimportする文がコメントアウトされているので、"/*"と"*/"を削除してimportする文を有効化すること。
</div></details>

| 入力例 | 出力例 |
| ---- | ---- |
| 7 | 7 |
| 12 | Fizz |
| 30 | FizzBuzz |

### Q.08 FizzBuzz問題
正の整数nを受け取り、1からnまでの値を順にQ.06で実装した関数FizzBuzzSingleに入力した結果を空白区切りで標準出力に表示する関数FizzBuzzを実装せよ。
<details><summary>ヒント</summary><div>
Q.06のFizzBuzzSingle関数を利用すれば簡単。<br>
標準出力への出力には、Q00.goの関数Helloで使ったfmt.Printf関数を利用する、ただし、Q00.goにならってfmtのimportを行うこと。
</div></details>

| 入力例 | 出力例 |
| ---- | ---- |
| 15 | 1 2 Fizz 4 Buzz Fizz 7 8 Fizz Buzz 11 Fizz 13 14 FizzBuzz |

### Q.09 FizzBuzz応用１
正の整数n,a,bを受け取り、1からnまでの値i ($1\leq i\leq n$) について、iがa×bの倍数ならFizzBuzz,a×bの倍数以外でaの倍数ならFizz、aの倍数以外でbの倍数ならBuzz、それ以外ならiを文字列に変換したものを空白区切りで順に標準出力に表示する関数FizzBuzz2を実装せよ。

| 入力例 | 出力例 |
| ---- | ---- |
| 7,2,3 | 1 Fizz Buzz Fizz 5 FizzBuzz 7 |
| 12,2,6 | 1 Fizz 3 Fizz 5 Fizz 7 Fizz 9 Fizz 11 FizzBuzz |
| 12,4,6 | 1 2 3 Fizz 5 Buzz 7 Fizz 9 10 11 Fizz |

## 基礎編

### Q.10 階乗関数
正の整数nを受け取り、nの階乗(n×(n-1)×(n-2)×...×2×1)の値を求める関数factを実装せよ。

| 入力例 | 出力例 |
| ---- | ---- |
| 3 | 6 |
| 5 | 120 |
| 7 | 5040 |

### Q.11 フィボナッチ数列
フィボナッチ数列$\{ F_{n} \}$は、以下の式で定められる。
$$
  \begin{cases}
	F_{n}=F_{n-1}+F_{n-2}&(n \geq 2)\\
	F_{1}=F_{0}=1
  \end{cases}
$$
自然数nを受け取り、$F_{n}$の値を返す関数fibを実装せよ。

| 入力例 | 出力例 |
| ---- | ---- |
| 2 | 2 |
| 5 | 8 |
| 7 | 21 |

### Q.12 配列の総和
整数の配列sを受け取り、sに格納されている値の総和を返す関数arraysumを実装せよ。

| 入力例 | 出力例 |
| ---- | ---- |
| [3,2,4] | 9 |
| [0,-2,5,8] | 11 |
| [2,2,2,2,2] | 10 |

### Q.13 配列の２乗
整数の配列sを受け取り、sに格納されている各値を２乗した配列を返す関数arraysquareを実装せよ。

| 入力例 | 出力例 |
| ---- | ---- |
| [3,2,4] | [9,4,16] |
| [0,-2,5,8] | [0,4,25,64] |
| [2,2,2,2,2] | [4,4,4,4,4] |

### Q.14 スライスの反転
整数のスライスsを受け取り、sを反転したものをs'とする。このとき、sの末尾にs'を結合したスライスを返す関数reverseconcatを実装せよ。

| 入力例 | 出力例 |
| ---- | ---- |
| [3,2,4] | [3,2,4,4,2,3] |
| [0,-2,5,8] | [0,-2,5,8,8,5,-2,0] |
| [2,2,2,2,2] | [2,2,2,2,2,2,2,2,2,2] |

### Q.15 偶数分割問題
整数のスライスsと整数kを受け取る。また、スライスtには初め要素が入っていないものとする。
以下の操作をiを1からNまで変化させて順に行ったとき、得られるスライスtを求めよ。
操作:<br>
sのi番目の要素$s_{i}$が偶数ならば$\frac{s_{i}}{2}$を2つ、奇数であれば$s_{i}$を1つ、tの末尾に追加するという操作を、

| 入力例 | 出力例 |
| ---- | ---- |
| [3,2,4] | [3,1,1,2,2] |
| [0,-2,5,8] | [0,0,-1,-1,5,4,4] |
| [2,2,2,2,2] | [1,1,1,1,1,1,1,1,1,1] |


### Q.16 盤面上の駒
文字列のスライスのスライスSを受け取る。Sは要素数wのスライスh個のスライスである。
Sのi番目のスライスのj番目の要素を$S_{i,j}$と表すこととする。
いま、縦hマス横wマスの盤面が存在し、盤面のi行目のj列目に自分の駒が存在するとき$S_{i,j}="A"$、相手の駒が存在するとき$S_{i,j}="B"$、駒が存在しないとき$S_{i,j}="O"$となっている。
また、各マスに駒は最大1つしか存在しない。
このとき、盤面上に存在する自分の駒の数から相手の駒の数を引いた値を求めよ。

| 入力例 | 出力例 |
| ---- | ---- |
| [[A,B,O],<br> [O,A,A],<br> [B,O,O]] | 1 |
| [[A,B,B,A],<br> [O,A,A,B],<br> [B,B,O,A],<br> [B,O,O,O]] | -1 |
| [[O,B,B,A,A],<br> [O,O,A,B,O],<br> [B,O,O,A,A],<br> [B,B,O,O,B],<br> [O,A,A,O,O]] | 0 |

### Q.17 出現回数数え上げ


### Q.18 最小公倍数と最大公約数



### Q.19 FizzBuzz応用２

## 初級文法編

### Q.20 平均の計算２
2つの整数x,yを入力として受け取り、この平均をfloat32型の小数として出力する関数Average2をQ10.goに実装せよ。

| 入力例 | 出力例 |
| ---- | ---- |
| 5, 3 | 4 |
| 6, 9 | 7.5 |
| -8, 3 | -2.5 |
