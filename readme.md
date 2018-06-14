# VRChat OSC Bridge (Extremely Beta)
VRChat の WebPanel が OSC信号を扱うためのサーバー・クライアントプラグイン

## 内容
 - ワールド側組み込みプラグイン (`vrc-osc-bridge.js`)
 - ワールド側組み込みデモページ (`index.html`) 
 - クライアント側プラグイン (`bin/main.exe`)

## 説明
VRChat の WebPanel に組みこまれた `vrc-osc-bridge.js` が、クライアント側プラグインとやり取りをすることで、`OSCメッセージ` を間接的に WebPanel から VRChat に送信します。

## 注意
このソフトウェアの利用に際し、一切の責任を負いません。
また、開発中であり実用には適しません。

## 関連
 - [YouTube Live のスーパーチャット発生イベントをメッセージ送信元とするブックマークレット
 (VRChat OSC Bridge 専用ツール)](https://gist.github.com/yukimochi/c0f9b756188fefa6d01a89865446a740)