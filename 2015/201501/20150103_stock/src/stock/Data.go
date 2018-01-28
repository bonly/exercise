package stock

type Day struct{
	PK int;
	Code string `db:"Code"`;
	Dt   int `db:"Dt"`;
	Open float64 `db:"Open"`;
	High float64 `db:"High"`;
	Low  float64 `db:"Low"`;
	Close float64 `db:"Close"`;
	Volume float64 `db:"Volume"`;	
	Adj_Close float64 `db:"Adj_Close"`;
};

type Symbol struct{
	Code string;
	Update_at int;
	Ask_price float64;
	Ask_size  int;
	Bid_price float64;
	Bid_size  int;
	Last_trade_price float64;
	Previous_close float64;
	Trading_halted bool;
};