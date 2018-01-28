package view

import(
	log "glog"
)

type Point struct{
	Hight int;
	Low   int;
};

type Line struct{
	Trend     bool;
	Points [2]Point;
};


