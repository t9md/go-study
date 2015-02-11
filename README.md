# What's This?
memo in self learning

# Reference

## By Google
* [A Tour of Go](https://tour.golang.org/welcome/1)
* [How to Write Go Code](https://golang.org/doc/code.html)
* [The Go Programming Language Specification](https://golang.org/ref/spec)
* [Effective Go](https://golang.org/doc/effective_go.html)
* [Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)

## Other
* [An Introduction to Programming in Go](http://www.golang-book.com)
* [Go by Example](https://gobyexample.com)


# 全体
[An Introduction to Programming in Go](http://www.golang-book.com) は言葉が正確で、channel なり、goroutine なりslice なりの説明が非常に簡潔かつ正確なので、ある程度理解した後に再読しても、新たな発見があるはず。  

# 依存 package のチェック？

* 全部
go list ...

* カレントディレクトリ配下
go list ./...

# Interface
Interface は method セットをグルーピングした型。  
method 群ということは、つまり振る舞いセットを定義した型。  
引数で interface 型を受け取る時、受け取る 変数の'振る舞いへの期待' を表明している。  

Interface-A 型の変数には、Interface-A を満たす(=期待に答える) 変数は全て格納可能。  
Interface 型の変数は、この変数は"こういった目的で使いますよ" と表明していると言える。  
例えば Stringer Interface の定義は以下の通りだが  

```Go
type Stringer interface {
    String() string
}
```

` var s Stringer = var1` だと、ver1 は var1.String() が呼ばれる事を覚悟して置かなければならない。 

# Pointer
ポインタはちゃんと理解しないといけない。  
~~function, map, slice は参照タイプなので、ポインタ渡す必要なし(??? need check)~~  ~~そんな事無い。~~
やっぱそんなことある。 map, slice はポインタ渡す必要なし、勝手に Pointer

custom type の method 定義時、Ponter を受け取るべきか？そもそも関数の引数に Pointer を受け取るべきか？
大体常に Pointer で受け取る様にしとけばいいんじゃないか？スタックへのコピーも発生しないのでパフォーマンス的にも
良いし、function(or method) が引数のフィールドを書き換えたい時も、コピーだと意味ないし、、、
という考え方がある。大体これでよいが、[CodeReviewComments#pass-values](https://github.com/golang/go/wiki/CodeReviewComments#pass-values) ではより突っ込んで検討していて、Value 渡しの方が、Pointer 渡しより適切なケースを上げている。  
でも、やはりこれは結構ややこしいから状況によるし、、結局最後に `Finally, when in doubt, use a pointer receiver.`(最後に、確信がなければPointer をレシーバにしろ) って言ってる。  

# Channel
関数の引数でチャネルを受け取るとき、宣言方法がいろいろあるけどどうちがうのか？  

* [Different ways to pass channels as arguments in function in go (golang)](http://stackoverflow.com/questions/24868859/different-ways-to-pass-channels-as-arguments-in-function-in-go-golang
)

関数宣言時、引数としてのチャネル、direction を明示できるならしたほうが良い。  
別の方向でチャネルを使おうとした場合にコンパイラが error 出す(本当?→本当だった。)  

```Go
func serve(ch <-chan interface{}){ //do stuff } // drection: read only

func serve(ch chan<- interface{}){ //do stuff } // dir: write purpose

func serve(ch chan interface{}){ //do stuff } //dir: bi-directional

func server(ch *chan interface{}){ //do stuff} // dir: depends on
```

## Channel のバッファサイズについて
Channel はバッファサイズを指定しなければ、バッファサイズ0 で作られる。  
この場合、チャネルへの書き込みはブロックする。  
"準備が出来た時" に書き込めるが、"準備が出来る" とは、"読み込み側が読み込もうとした時"、であることに注意  
なので、以下の様なコードは dead lock になる。  
Compile できるが、runntime error になる。(なぜcompile 時に検出できないのか？？)  

```Go
ch := make(chan string)
ch <- "hello"
fmt.Println(<-ch)
```

以下だと、hello と表示される。  
バッファサイズ 1 があるので、2行目でブロックせず、3行目に進むから。  
```Go
ch := make(chan string, 1)
ch <- "hello"
fmt.Println(<-ch)
```

以下の例では、一つ目の無名関数のhello1は表示されずにプログラムが終了する。 2秒sleep なので、fmt.Println(<-ch) は2つ目のhello2を受け取る。  
```Go
ch := make(chan string, 1)
go func() {
  time.Sleep(time.Second * 2)
  ch <- "hello1"
}()
go func() {
  time.Sleep(time.Second * 1)
  ch <- "hello2"
}()
fmt.Println(<-ch)
```

バッファサイズが無いときは、チャネルがつながるまでは書き込めない。  
需要と供給が直接一致した時に書き込み、読み込みが同時に起こる。  
例を示す。  

```Go
func main() {
	ch := make(chan string) // バッファなし
	go func() {
		ch <- "hello1" // 読み込みがと繋がってないので block する。
		fmt.Println("  SENT!!!")
	}()
	time.Sleep(time.Second * 2)
	fmt.Println(<-ch) // この時点で初めて hello1 を書き込む groutine のブロックが解かれる
	time.Sleep(time.Second * 1)
}
// 結果
// hello1
//   SENT!!!
```

```Go
func main() {
	ch := make(chan string, 1) // バッファあり
	go func() {
		ch <- "hello1" // すぐ書き込めるので
		fmt.Println("  SENT!!!") // Channelの読み込みより先にこちら
	}()
	time.Sleep(time.Second * 2)
	fmt.Println(<-ch) // 読み込みの時点では書き込み終わってる(すでにhello1が溜まっている)
	time.Sleep(time.Second * 1)
}
// 結果
//   SENT!!!
// hello1
```

# Array and Slice
Array と Slice は別物。別物というのは、別の型だということ。  
Slice は Array が"実体" ではあるが、データ型としては別物である。  
見分け方は、宣言時、length(=size) があれば、Array, なければ Slice  

```Go
var array_1 [10]int
var slice_1 []int
fmt.Println(array_1) // => [0 0 0 0 0 0 0 0 0 0]
fmt.Println(slice_1) // => []

// 初期化
var array_2 = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
var slice_2 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
fmt.Println(array_2) // => [1 2 3 4 5 6 7 8 9 10]
fmt.Println(slice_2) // => [1 2 3 4 5 6 7 8 9 10]
```

# 変数のスコープ
小文字は package 外から見えない。(private to package)  
大文字は package 外から見える(Public)  
各ファイルの所属 package はファイルの先頭で宣言する。  
複数のファイルが同じ package であれば、同一 package なので、小文字の変数は見える。  
ファイルスコープではなく、package スコープであることに注意。  

`{``}` curly braces によるスコープ。
自分が内包されている curly braces の変数は見える。
```Go
{
  outer := 1
  {
    // ここから outer は見える。
    inner := 2
  }
  // ここで inner は見えない。
}
```

# 宣言の読解の慣れ
C言語でもポインタの宣言を理解するには慣れが必要だった。  
以下のような宣言をパット読めるようにならないと  

```Go
members := make([]*Member, len(nodes))
```

`[]*Member` は `*Member` のスライス(=コレクション)  
`[]` が来た時点で瞬時にサイズ指定のないコレクション(=スライス)だ、  
何のスライスだろう？`*` に出会い、pointer か、何への Pointer か、あ、Member がその先に入ってるのね。。。  
という、前→後への自然な直読が出来る様になること、その後、塊として瞬間的に理解できるようになる。  
これは、英語の語学学習と同じ話。  

# new() と make()
new() は Pointer を返す。
make() はmake した型 T のインスタンスそのものを返す。

# builtin

# import

# aa
