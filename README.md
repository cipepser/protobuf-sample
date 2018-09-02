# protobuf-sample

[Protocol Buffers](https://github.com/google/protobuf)でいろいろ検証したメモ

## `int32` -> `bool`の変換

* [protobufのboolはどこまでcompatibleなのか - 逆さまにした](http://cipepser.hatenablog.com/entry/protobuf-bool)

※`user/`や`UserWrite.go`、`UserRead.go`は上記記事参照。

## enum

TODO: ブログにまとめる

※`max/`や`EnumWrite.go`、`EnumRead.go`は上記記事参照。


## 64-bit -> 32-bitの変換

TODO： ブログにまとめる

※フォルダやファイルは`max`でつけているもの

## エンコーディング

```proto
// zero.proto
syntax = "proto3";

package zero;

message Person {
  Name name = 1;
  Age age = 2;
}

message Name {
  string value = 1;
}

message Age {
  int32 value = 1;
}
```

```go
// ZeroWrite.go
package main

import (
	"fmt"
	"io/ioutil"
	"log"

	pb "github.com/cipepser/protobuf-sample/zero"
	"github.com/golang/protobuf/proto"
)

func main() {
	p := &pb.Person{
		Name: &pb.Name{Value: "Alice"},
		Age:  &pb.Age{Value: 20},
	}

	out, err := proto.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out)
	if err := ioutil.WriteFile("./zero/person.bin", out, 0644); err != nil {
		log.Fatalln("Failed to write:", err)
	}
}
```

```sh
❯ go run zero.go
[10 7 10 5 65 108 105 99 101 18 2 8 20]

❯ hexdump zero/person.bin
0000000 0a 07 0a 05 41 6c 69 63 65 12 02 08 14
000000d
```

`out`を`[]byte`のまま標準出力したときと、ファイル出力してhexdumpしたときで違う？  
→表記が違うだけ

```bin
10  7 10  5  65 108 105 99 101 18 2 8 20
              A  l   i   c   e
10  7 10  5  66 111 98 18 2 8 20

0a 07 0a 05 41 6c 69 63 65 12 02 08 14

(65)10 = (0100 0001)b = (41)hex
(108)10 = (0110 1100)b = (6c)hex
(105)10 = (0110 1001)b = (69)hex
(99)10 = (0110 0011)b = (63)hex
(101)10 = (0110 0101)b = (65)hex
(18)10 = (0001 0010)b = (12)hex
```

`fmt.Printf("% x", out)`で以下のようにできる。

```go
❯ go run zero.go
0a 07 0a 05 41 6c 69 63 65 12 02 08 14
```

### バイナリの読み方

まずprotobufではどの構造体にパースするのかを与えるので、以下で定義した`Person`にパースしていく。

```proto
message Person {
  Name name = 1;
  Age age = 2;
}
```

[Protocol Buffers のエンコーディング仕様の解説](https://qiita.com/aiueo4u/items/38195248a29e9ff719c7)にあるように以下が基本となる。

> key = タグナンバー * 8 + タイプ値

タイプ値は、[公式ドキュメント](https://developers.google.com/protocol-buffers/docs/encoding)でwire typesとして以下のように定義されている。

|  Type | Meaning | Used For |
|  ------ | ------ | ------ |
|  0 | Varint | int32, int64, uint32, uint64, sint32, sint64, bool, enum |
|  1 | 64-bit | fixed64, sfixed64, double |
|  2 | Length-delimited | string, bytes, embedded messages, packed repeated fields |
|  3 | Start group | groups (deprecated) |
|  4 | End group | groups (deprecated) |
|  5 | 32-bit | fixed32, sfixed32, float |


改めて以下をパースしていく。

```
0a 07 0a 05 41 6c 69 63 65 12 02 08 14
```

`()`内の数字は何進数で表記しているかを表す。

まず初めの`0a(16)`は、  
`10(10)` = タグ`name(1)` * 8 + `Length-delimited(type 2)`  
となる。  
※`Name`は自身で定義したmessageであり、表中の`embedded message`が該当し、`Length-delimited`となる。

続く`07(16)`は、`name`のlengthなので、`0a 05 41 6c 69 63 65`を`Name`として読んでいく。

なので、`Name`最初の`0a(16)`は、  
`10(10)` = タグ`value(1)` * 8 + `Length-delimited(type 2)`  
となる。

続く`05(16)`は、`value`のlengthなので、`41 6c 69 63 65`を`string`として読んでいく。
utf8(この文字範囲ならASCIIと同じだけど)として読むと`41 6c 69 63 65`は`Alice`となる。

このまま残りの`12 02 08 14`も読んでいく。

`0a(12)`は、  
`18(10)` = タグ`age(2)` * 8 + `Length-delimited(type 2)`  
となる。

続く`08(16)`は、`value(1)`のlengthなので、`14`を`int32`として読んで`14(16)` = `20(10)`

![protobuf](https://github.com/cipepser/protobuf-sample/blob/master/img/protobuf.png)

#### vimでhex表記するときのコマンド

```sh
❯ vim b <file>
:%!xxd
```

編集して保存

```sh
:%!xxd -r
:wq
```

## Rustでprotobufを試す

以下にブログでまとめた。

[Golangで出力したprotobufバイナリをRustで読み込む - 逆さまにした](https://cipepser.hatenablog.com/entry/protobuf-read-in-rust)

## mapのバイナリを読む

`map.proto`を以下のように定義。

```proto
syntax = "proto3";
package map;

message User {
  map<string, int32> Name2Age = 1;
}
```

以下をバイナリで書き込む。

```go
m := map[string]int32{}
m["Alice"] = 20
user := &pb.User{
  Name2Age: m,
}
```

結果

0a09 0a05 416c 6963 6510 14


0x0a: 0d10 = 1 * 8 + 2
> key = タグナンバー * 8 + タイプ値

type2: Length-delimited
string, bytes, embedded messages, packed repeated fields

これが`09`バイト分 = `0a05 416c 6963 6510 14`。

また`0a`なので、Length-delimitedで`05`バイト

`41 6c 69 63 65`は`Alice`

残るは`10 14`

0x10: 0d16 = 2 * 8 + 0

なので、タグ番号2の`Varint`として`0x14` = `0b20`を読む

`key`がタグ番号1、`value`がタグ番号2であることについては、[Protocol Buffers - Language Guide (proto3)](https://developers.google.com/protocol-buffers/docs/proto3)の`Backwards compatibility`にmapが以下と同一であることが書かれている。

```proto
message MapFieldEntry {
  key_type key = 1;
  value_type value = 2;
}

repeated MapFieldEntry map_field = N;
```

### mapに複数入っている場合は？

Bobを追加する。

```go
m := map[string]int32{}
m["Alice"] = 20
m["Bob"] = 25
user := &pb.User{
  Name2Age: m,
}
```

結果

0a09 0a05 416c 6963 6510 140a 070a 0342 6f62 1019

先程の結果と並べてみる。

0a09 0a05 416c 6963 6510 14

`0a 070a 0342 6f62 1019`がうしろに増えている。

先程と同じように読むと

`0a`(タグ番号1のLength-delimited)で`07`バイト

`0a`(タグ番号1のLength-delimited)で`03`バイト = `42 6f62`(Bob)

`10`(タグ番号2のVarint)で`0x19` = `0d25`

## repeatedのバイナリを読む

`repeated.proto`を以下のように定義する。

```proto
syntax = "proto3";
package repeated;

message User {
  repeated string Name = 1;
}
```

以下を読み込む。

```go
user := &pb.User{
  Name: []string{"Alice"},
}
```

※ここで`string`をsliceにしないと以下のようにコンパイルが通らない。

```sh
./RepeatedWrite.go:13:9: cannot use "Alice" (type string) as type []string in field value
```

さて、出力のバイナリは以下のようになった。

0a05 416c 6963 65

`0a`(タグ番号1のLength-delimited)で`05`バイト

`41 6c 69 63 65`は`Alice`


### `Bob`を追加

```go
user := &pb.User{
  Name: []string{"Alice", "Bob"},
}
```

読み込み結果。

0a05 416c 6963 650a 0342 6f62

`0a`(タグ番号1のLength-delimited)で`05`バイト

`41 6c 69 63 65`は`Alice`

`0a`(タグ番号1のLength-delimited)で`03`バイト

`42 6f 62`は`Bob`

### mapと比較

#### `Alice`のみ

| synta | binary |
| ------ | ------ |
|  map | 0a09 0a05 416c 6963 6510 14 |
|  repeated | 0a05 416c 6963 65 |


#### `Alice`と`Bob`

| synta | binary |
| ------ | ------ |
|  map | 0a09 0a05 416c 6963 6510 140a 070a 0342 6f62 1019 |
|  repeated | 0a05 416c 6963 650a 0342 6f62 |

## oneofのバイナリを読む

```proto
syntax = "proto3";
package oneof;

message User {
  oneof Result {
    string Ok = 1;
    string Err = 2;
  }
}
```

以下(`Ok`)で書き出す。

```go
user := &pb.User{
  Result: &pb.User_Ok{Ok: "Alice"},
}
```

結果。

0a05 416c 6963 65

`0a`(タグ番号1のLength-delimited)で`05`バイト

`41 6c 69 63 65`は`Alice`

`Err`でも見てみる。

```go
user := &pb.User{
  Result: &pb.User_Err{Err: "Alice"},
}
```

1205 416c 6963 65

`0x12` = `0d18(=2*8+2)`(タグ番号2のLength-delimited)で`05`バイト

`41 6c 69 63 65`は`Alice`


最初(`0a`->`12`)が違うだけ。

## enumのreseavedについて

まずはenumではなく、単純に`message`でreservedしたフィールドを使おうとすると怒られることを確認する。

```proto
syntax = "proto3";
package reserved;

message User {
  reserved 1;
  enum Type {
    NORMAL = 0;
    PREMIUM = 1;
  }
  bool type = 1;
}
```

上記でタグ`1`を予約した上で、`bool type = 1;`で使おうとした状況で`protoc`してみる。

```sh
❯ protoc -I=reserved/ --go_out=reserved/ reserved.proto
reserved.proto: Field "type" uses reserved number 1.
```

確かに怒られる。


## References
* [Proto3 Language Guide（和訳）](https://qiita.com/CyLomw/items/9aa4551bd6bb9c0818b6)
* [Protocol Buffers - Encoding](https://developers.google.com/protocol-buffers/docs/encoding)
* [Protocol Buffers のエンコーディング仕様の解説](https://qiita.com/aiueo4u/items/38195248a29e9ff719c7)
* [The Go Programming Language Specification](https://golang.org/ref/spec)
* [Protocol Buffers - Language Guide (proto3)](https://developers.google.com/protocol-buffers/docs/proto3)
* [Protocol Buffers - Google's data interchange format](https://github.com/google/protobuf)
* [protobufのboolはどこまでcompatibleなのか - 逆さまにした](http://cipepser.hatenablog.com/entry/protobuf-bool)
* [RustでProtocol Buffers - ザネリは列車を見送った ブログという名の備忘録](https://www.zaneli.com/blog/20161217)