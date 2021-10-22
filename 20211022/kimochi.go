package main

import(
	"time"
	"fmt"
)


// ニ○リで体験したサービスをデータ構造で表すとどうなるか試したくなっただけ
/*
体験したこと（雑）
・ベッドを買う
・配送手続きにすすむ
・住所、名前、受取日、受け取り時間を記入して渡す
・支払う
*/

// 規格
type Standard struct {
	color string
	size string // S,M,L
}

// ベッド
type Bed struct {
	name string
	standard Standard
	price float64
	tax_rate float64
}

// 顧客情報
type Customer struct {
	name string
	address string
}

// 在庫情報
type Stock struct {
	bed Bed
	stock int
}

// 注文と商品は1on1とする
type OrderDetail struct {
	cutomer Customer
	bed Bed
	bed_count int
	delivery_date string // 配送日
	purchase_date string // 顧客の購入日
	payment_status int // 1:未払い　2:支払い済み リアル店舗想定だといらないかな・・・。ECが絡むなら必要
	is_deliverd int // 配送したかどうかを示す 1:未配達　2:配達済み
	check_sum float64
	check_sum_status int // 1:仮売上 2:売上計上
}

// メソッドを作成。上から順に流していく。
/*
体験したこと（雑）
・ベッドを買う　→　ベッドの在庫が減少する
・配送手続きにすすむ
・住所、名前、受取日、受け取り時間を記入して渡す　→　受注日、配送業者への手配、配送したか、売り上げ
　・受注日　→　記入表を受け取った時点で取得する
　・配送業者への手配　→　これはよくわからんけど、いつ取りにくるのか？
　・配送したか　→　配送業者から連絡がきたらステータスを変更する　1:未配達　2:配達済み
・支払う
　・売り上げ　→　金額だけ。webではなく店舗だから。
*/


// ・ベッドを買う　→　ベッドの在庫が減少する
func BuyBed() OrderDetail {
	// 顧客データを教えてもらう
	customer := Customer{"zatu", "kakuu kakuu"}
	// ベッドの規格を選んでもらう
	standard := Standard{"blue", "S"}
	// ベッドを選んでもらう 逆だよな・・・
	bed := Bed{"hukahuka", standard, 20000.0, 0.08}
	// 買い上げ合計
	check_sum := bed.price * bed.tax_rate
	// 配送希望日
	delivery_date := time.Date(2020, 5, 20, 12, 0, 0, 0, time.Local).String()
	// 買い物した日
	now := time.Now()
	purchase_date := now.String()
	// 注文詳細データ作成
	order := OrderDetail{customer, bed, 3, delivery_date, purchase_date, 1, 1, check_sum, 1}

	return order
}

// 配送処理
func sendItem(order OrderDetail) {
	// 決済ステータスを支払い済みに変更する
	order.is_deliverd = 2
	// 在庫を減らす
	// ????
	// 売上に計上する
	order.check_sum_status = 2
}

func main(){
	order := BuyBed()
	sendItem(order)
	fmt.Printf("(%%+v) %+v\n", order)
}

/*
(%+v) {
	cutomer:{name:zatu address:kakuu kakuu} 
	bed:{name:hukahuka standard:{color:blue size:S} price:20000 tax_rate:0.08} 
	bed_count:3 
	delivery_date:2020-05-20 12:00:00 +0000 UTC 
	purchase_date:2021-10-22 14:08:23.913156754 +0000 UTC m=+0.000680834 
	payment_status:1 
	is_deliverd:1 
	check_sum:1600 
	check_sum_status:1
}
*/