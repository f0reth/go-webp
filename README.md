# go-webp

## 実行に必要なファイル

このライブラリは [purego](https://github.com/ebitengine/purego) を使って libwebp の共有ライブラリを動的に読み込みます。事前に以下のファイルをインストールしてください。

| OS | 必要なファイル |
|---|---|
| Windows | `libwebp.dll`、`libwebpdemux.dll` |
| Linux | `libwebp.so`、`libwebpdemux.so` |
| macOS | `libwebp.dylib`、`libwebpdemux.dylib` |

## 使い方

### インポート

```go
import "github.com/f0reth/go-webp"
```

### デコード（WebP → image.Image）

```go
f, _ := os.Open("input.webp")
defer f.Close()

img, err := webp.Decode(f)
```

### アニメーション WebP のデコード

```go
f, _ := os.Open("anim.webp")
defer f.Close()

ret, err := webp.DecodeAll(f)
// ret.Image: 各フレームの画像スライス
// ret.Delay: 各フレームの表示時間（ミリ秒）
```

### エンコード（image.Image → WebP）

```go
f, _ := os.Create("output.webp")
defer f.Close()

err := webp.Encode(f, img)
```

### エンコードオプション

```go
err := webp.Encode(f, img, webp.Options{
    Quality:  90,      // 画質 0〜100（デフォルト: 75）
    Lossless: false,   // true にすると可逆圧縮（Quality は無視）
    Method:   4,       // エンコード速度 0（速い）〜6（高品質、デフォルト: 4）
    Exact:    false,   // true にすると透明部分の RGB 値を保持
})
```

### ライブラリの読み込み確認

```go
if err := webp.Dynamic(); err != nil {
    log.Fatal("共有ライブラリが見つかりません:", err)
}
```
