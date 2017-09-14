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