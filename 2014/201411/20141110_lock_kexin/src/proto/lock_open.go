/*
auth: bonly
create: 2015.10.27
*/
package proto

import(
"encoding/json"
)

///third/lock/open
type Lock_Open struct{
	Lock;
	Lock_id int `json:"lock_id"`;
};

func (this *Lock_Open) Marshal(appid string, key string)(ret []byte, err error){
	this.Lock.Marshal(appid, key);
	this.Sign, err = Gen_Sign(*this, key);

	return  json.MarshalIndent(*this, " ", " ");
}