# What's This?
memo in self learning

# Reference

## Official
* [A Tour of Go](https://tour.golang.org/welcome/1)
* [How to Write Go Code](https://golang.org/doc/code.html)
* [The Go Programming Language Specification](https://golang.org/ref/spec)
* [Effective Go](https://golang.org/doc/effective_go.html)
* [Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
* [FAQ](https://github.com/astaxie/build-web-application-with-golang)

## Other
* [An Introduction to Programming in Go](http://www.golang-book.com)
* [Learn X in Y minutes:Where X=Go](http://learnxinyminutes.com/docs/go/)
* [Go by Example](https://gobyexample.com)
* [build web application with golang](https://github.com/astaxie/build-web-application-with-golang)


# 全体
[An Introduction to Programming in Go](http://www.golang-book.com) は言葉が正確で、channel なり、goroutine なりslice なりの説明が非常に簡潔かつ正確なので、ある程度理解した後に再読しても、新たな発見があるはず。  

# 依存 package のチェック？

* 全部
go list ...

* カレントディレクトリ配下
go list ./...

## package

package で構成される
package 宣言はファイルの最初に来る。
package name はそのファイルの filepath の basename (=ディレクトリ名)
にするのがしきたり。

# Interface
Interface は method セットをグルーピングした型。  
method 群ということは、つまり振る舞いセットを定義した型と言える。  
引数で interface 型を受け取る時、受け取る 変数の'振る舞いへの期待' を表明している。  

Interface-A 型の変数には、Interface-A を満たす(=satisfy, 期待に答える) 変数は全て格納可能。  
Interface 型の変数は、この変数は"こういった目的で使いますよ" と表明していると言える。  
例えば Stringer Interface の定義は以下の通りだが  

```go
type Stringer interface {
    String() string
}
```

` var s Stringer = var1` だと、ver1 は var1.String() が呼ばれる事を覚悟しておかなければならない。 
→ "覚悟する" というか、String() メソッドが定義されていなければ、型チェックで弾かれる。

# Pointer
ポインタはちゃんと理解しないといけない。  
~~function, map, slice は参照タイプなので、ポインタ渡す必要なし(??? need check)~~  ~~そんな事無い。~~
やっぱそんなことある。 map, slice はポインタ渡す必要なし、元々 Pointer

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

```go
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

```go
ch := make(chan string)
ch <- "hello"
fmt.Println(<-ch)
```

以下だと、hello と表示される。  
バッファサイズ 1 があるので、2行目でブロックせず、3行目に進むから。  
```go
ch := make(chan string, 1)
ch <- "hello"
fmt.Println(<-ch)
```

以下の例では、一つ目の無名関数のhello1は表示されずにプログラムが終了する。 2秒sleep なので、fmt.Println(<-ch) は2つ目のhello2を受け取る。  
```go
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

```go
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

```go
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

```go
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
```go
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

```go
members := make([]*Member, len(nodes))
```

`[]*Member` は `*Member` のスライス(=コレクション)  
`[]` が来た時点で瞬時にサイズ指定のないコレクション(=スライス)だ、  
何のスライスだろう？`*` に出会い、pointer か、何への Pointer か、あ、Member がその先に入ってるのね。。。  
という、前→後への自然な直読が出来る様になること、その後、塊として瞬間的に理解できるようになる。  
これは、英語の語学学習と同じ話。  

# new() と make()
new() は Pointer を返す。
make() はmake した型 T 自体を返す。

```go
func new(Type) *Type
```
```go
func make(Type, size IntegerType) Type
```

# Tips
[effective-goではない何か](http://yoppi.hatenablog.com/entry/2014/01/07/084154)


# builtin

# import

```go
import "fmt"
import "math/cmplx"
```

   ↓

```go
import(
  "fmt"
  "math/cmplx"
)
```

```go
import(
  . "fmt" // dot をつけると、Println("hoge") の様にprefix(fmt.) 無しで呼べる
  "math/cmplx"
)
```

```go
const f = "%T(%v)\n" // 定数は `const` keyword をつける。
constant に `:=` による型推論は使えない
```

毎回 var 付けなくても、次のようにまとめられる
```go
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)
```


# ゼロ値初期化

explicit initial value で初期化しなかった変数は、ゼロ的な値で初期化される
  0 for numeric
  false for boolean
  "" for strings

# effective-go メモ

## Redeclaration and reassignment(再宣言と、再代入の(特例??))
[redeclaration](https://golang.org/doc/effective_go.html#redeclaration)
```go
f, err := os.Open(name)
if err != nil {
    return err
}
d, err := f.Stat()
if err != nil {
    f.Close()
    return err
}
codeUsing(f, d)
```
上記のコードは err が同一スコープで `:=` で代入されているが、これは以下の特例によるもので、この特例は err を上記な様なケースで使うことを可能にするために設けられているようだ。
`:=` は宣言と代入を同時に行う演算子なので、同一スコープの同じ変数名(identifier)に対して`:=`を２度使うのはだめなはず。しかし、  
再宣言と、再代入の特例は以下の条件を満たす限り合法である。
* 既存の変数vと同一スコープで、
* 再代入する値は既存変数vにassignable(型)であり、
* 最低、一つ以上の全く新しい変数が存在する時
つまり、最後の条件は multiple asiginment(多重代入)式で、その内の変数の一つでも完全に新規だったら、`:=` つかっても合法だよ。まあこれは err の再代入を許可できたほうが便利だから設けた特例だよ。
という理解をした。  

## For

array の index や、map の key にしか興味がなければ以下でOK.
```go
for key := range m {
    println(key)
}
```

## Type switch

`switch` は interface 型変数の動的な型を特定するためにも使うことが出来る。
このような型switch は、括弧内を `type` にして、type assertion の構文を使う。
switch の行で、変数代入しておけば、対応する型が代入され、switch 内で使うことが出来る。
このようなケースで名前を再利用(tを再利用している)するのは、イディオム。

```go
var t interface{}
t = functionOfSomeType()
switch t := t.(type) {
default:
    fmt.Printf("unexpected type %T", t)       // %T prints whatever type t has
case bool:
    fmt.Printf("boolean %t\n", t)             // t has type bool
case int:
    fmt.Printf("integer %d\n", t)             // t has type int
case *bool:
    fmt.Printf("pointer to boolean %t\n", *t) // t has type *bool
case *int:
    fmt.Printf("pointer to integer %d\n", *t) // t has type *int
}
```

上記で、t を `interface{}` に型定義しているのは、functionOfSomeType() で何が返ってくるか未定だからだ。
で、一旦何でも代入できる `interface{}` 型の変数で受け取った(代入)した後、type switch で type discover している。  
`switch t := t.(type) ` の部分で、使っている左の `t` は別に `val`とか`t`じゃない変数名でもよい。これ、イディオムらしいが、例としては混乱するな。。  
switch の行は完全に別スコープだから、`t` という新規の変数を宣言and代入しているだけで、別に再宣言and再代入ではない。
interface{} は、中身が何もない interface なので、このインターフェイスは「何の期待もしない」インターフェイスである。
メソッド呼び出しに対する期待(or要件)がゼロのインターフェイスなので、全ての型がこの、'期待' に答える事が出来る。
interface{} は存在しているだけでよい、何も出来なくても良い。という究極に寛容な型であるので、どんな型でも代入することが出来きる。  


## Pointer vs Value
理解したいので全部訳してみる。
[Effective Go 日本語訳](http://go.shibu.jp/effective_go.html#vs) を参考にしたが、とんんでもなく間違っている。英語の原語で読まないとエラく損をする、という良い例だな。
技術翻訳は、きれいな日本語でなくても良いから、意味を少しも捨てないように努力して訳す必要がある。
日本語訳はこの部分しか見ていないが、苦手でも英語直で読むほうが結局近道だ。

> Pointers vs. Values  

ポインタ vs 値  

> As we saw with ByteSize, methods can be defined for any named type (except a pointer or an interface); the receiver does not have to be a struct.  

ByteSize で我々が見てきた様に、メソッドはどんな名付けられた型(以下named type)(ポインタとインターフェイスを除く)に対しても定義することが出来る。レシーバは struct である必要はない。  


> In the discussion of slices above, we wrote an Append function. We can define it as a method on slices instead. To do this, we first declare a named type to which we can bind the method, and then make the receiver for the method a value of that type.  

上述のsliceの議論に於いて、我々は Append 関数を書いた。これをslice のメソッドとして定義する事も出来る。そうするには、最初にメソッドを紐付ける型を、named type として宣言し、つぎに、そのメソッドのレシーバをその型の値にする。  

```go
type ByteSlice []byte

func (slice ByteSlice) Append(data []byte) []byte {
    // Body exactly the same as above
}
```

> This still requires the method to return the updated slice. We can eliminate that clumsiness by redefining the method to take a pointer to a ByteSlice as its receiver, so the method can overwrite the caller's slice.  

これではしかし、まだメソッドから更新したスライスを返す必要がある。この煩雑さを解消するには、メソッド再定義して、ByteSlice へのポインタをレシーバとして受け取るようにするこだ。そうすればメソッドは呼び出し側のスライスを更新(overwrite)できる。  

```go
func (p *ByteSlice) Append(data []byte) {
    slice := *p
    // Body as above, without the return.
    *p = slice
}
```

> In fact, we can do even better. If we modify our function so it looks like a standard Write method, like this,  

実際のところ、もっとよく出来る。関数を標準の Write メソッドと同じになるように書き換える、こんな風に。  

```go
func (p *ByteSlice) Write(data []byte) (n int, err error) {
    slice := *p
    // Again as above.
    *p = slice
    return len(data), nil
}
```

> then the type `*ByteSlice` satisfies the standard interface io.Writer, which is handy. For instance, we can print into one.  

こうすると、`*ByteSlice` は io.Write の標準インターフェイスを充足するから、使い勝手がよくなる。例えば print で書き込むことも出来る。  

```go
    var b ByteSlice
    fmt.Fprintf(&b, "This hour has %d days\n", 7)

```

> We pass the address of a ByteSlice because only `*ByteSlice` satisfies io.Writer. The rule about pointers vs. values for receivers is that value methods can be invoked on pointers and values, but pointer methods can only be invoked on pointers.  

我々はここで、ByteSlice のアドレスを渡した。理由は io.Writer (のインターフェイス)を満たしているのは `*ByteSlice` のみだからだ。レシーバの"ポインタ vs 値"についての規則はこうだ。value メソッドはポインタに対しても、値(value)に対しても呼び出せるが、ポインタメソッドはポインタに対してのみ呼び出せる。  

> This rule arises because pointer methods can modify the receiver; invoking them on a value would cause the method to receive a copy of the value, so any modifications would be discarded. The language therefore disallows this mistake. There is a handy exception, though. When the value is addressable, the language takes care of the common case of invoking a pointer method on a value by inserting the address operator automatically. In our example, the variable b is addressable, so we can call its Write method with just b.Write. The compiler will rewrite that to (&b).Write for us.  

ポインタメソッドがレシーバを書き換える事が出来るから、こういう規則がある。つまり、これら(ポインタメソッド)を値に対して呼び出せてしまうと、メソッドは値のコピーを受け取るから、どんな変更も破棄されるだろう。そこで言語レベルで、このミスを許していないのだ。ただ、便利な例外規則がある。値がアドレスを特定できる類のものであれば(値がaddressableであれば), ポインタメソッドを値に対して呼び出す一般的なケースを、言語がケアして、自動でアドレス演算子(&)を挿入する。我々の例でいうと、変数 b は adressable だから、メソッド Write は、単に b.Write でも呼び出せる。コンパイラが我々のために、(&b).Write に書き換えてくれる。  

> By the way, the idea of using Write on a slice of bytes is central to the implementation of bytes.Buffer.  

ところで、byte のslice に対して、Write を使うアイデアは、 bytes.Buffer 実装の根幹だ。  

### 感想
つまり、メソッドを `func (v *T) meth(){}` の様に定義した時、

1. 型Tの値変数vに対して呼ぶのは間違い。
2. (Pointer to T)型に対して呼ぶのが正しい。
3. 便利な例外として v が addressable であれば、勝手にコンパイラが&を挿入して1のミスを修正してくれる。


```go
type T byte[]
v := T{}
v.meth()    // 1であげた間違い。v は値(value)なのに Pointer method を呼んでいる。
(&v).meth() // 2. これが正しい。レシーバもPointerだからPointer method が呼べる。
v.meth()    // しかし、3の"優しい" 例外によって、コンパイラが(&).meth() にして実行してくれる
```
3.の便利なケアが、個人的には好きではない。曖昧な理解でも勝手に修正されて動いてしまうと、ミスを指摘されることで学習するフィードバック型学習のチャンスが奪われてしまう。。しかし、毎回&をつけるのは面倒だから、つけたんだろう。しかし

* 本当は変だけど、便利のためにコンパイラがやってくれていると知っていて使う

のと、

* 経験から、これが正しいと(まちがった)理解をして使う

のとでは大違い。どちらも動く。しかし後者は正しくはないし、コンパイラのケアに対して無知なので間抜けだ。
