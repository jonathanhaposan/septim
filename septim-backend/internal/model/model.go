package model

type Transaction struct {
	StockCode string `json:"stock_code" bson:"stock_code"`
	Amount    int64  `json:"amount" bson:"amount"`
	Price     int64  `json:"price" bson:"price"`
	Type      string `json:"type" bson:"type"`
}

type Summary struct {
	TotalSellStock int64 `json:"total_sell_stock"`
	TotalBuyStock  int64 `json:"total_buy_stock"`
}
