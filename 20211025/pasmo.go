package main

/*
紛失時すぐ止められるPASMOを作る

クレジットカードは止められる
PASMOは止められない
→　なぜか？

クレジットカード：紛失後、電話し、名前と住所、カード番号などを合わせると停止可能
PASMO：紛失後、止められない

違和感：どうしてクレジットカードは電話なのか？ウェブからではダメなのか？

理想の仕様：
・電話すれば止められる
・アカウントを作成することで紐付けが発生し、ウェブからでも止められる

上記を新規導入する場合に起きるめんどくさいこと
・既存のアカウントなしユーザーの扱いはどうするのか？
　→　アカウント作成後、既存ユーザーにカード番号を入れてもらう？
　→　それはだめ。誰かのカードを拾って悪用可能だから。
　→　駅の窓口経由で止めてもらうのはどうだろう？

・既存ユーザーは買い直す場合、初回はタダでカードを発行し直す
　→　カード内にアクティブか否かのデバイスを埋め込む
　→　既存のマシンに埋め込みできないとか・・・？

仮定：
・ユーザーが途中でカードを止められるシステムを構築する
・既存マシンとの親和性は考えないものとする

仕様：
・電話すれば止められる
・アカウントを作成することで紐付けが発生し、ウェブからでも止められる

利用フロー：
・アプリかウェブ経由で個人情報を登録してもらう
・個人情報にカードIDを紐付ける
・その時点で状態がアクティブになる
・止めたい場合、窓口、電話、券売機、ウェブシステムから止められる
・決済手段の紐付けはウェブサイトから実行できる

開発内容：
・GO言語を用いた各種データの登録用APIを作成する
・WEBUIはReactで、MobileUIはFlutterで作成する
・改札側で認識するための疑似コードもGOで作成する
*/

/*
機能：

カード系：
・個人情報を登録/更新/削除　する
・カードを登録/更新/削除　する
・複数可能
・決済手段を登録/更新/削除　する
・各種カードの利用状態を登録/更新　する
　1：利用可能
　2：利用不可

改札系：
・カードがアクティブ状態か否か判定する
・これをウェブ経由で問い合わせる（これしんどくない？）　→　これを仮実装対象とする
　→　利用可能か否かを弾くのではなく、カードを完全停止する

高信頼性ネットワークの構築が重要だ・・・。

GOでやる必要があるサービスか少し怪しくてなってきた。
goldmanのフレームワークをGO用にリライトできないだろうか？
*/
