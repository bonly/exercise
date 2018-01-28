package Command

import (
	"encoding/json"
)

type ICmd interface{
	Decode([]byte)error;
};

type TCmd struct{
	 Func string;
	 Data interface{};
};

func (this *TCmd)Decode(pack []byte)(error){
	return json.Unmarshal(pack, this);
}