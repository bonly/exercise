/*
auth: bonly
create: 2015.10.27
*/
package proto

type Proto interface{
	Marshal(obj interface{})(ret []byte, err error);
	Unmarshal(pack []byte, vf interface{})(err error);
};

type Lock struct{
	Appid string `json:"appid"`;
	Sign string `json:"sign,omitempty"`;
};

func (this *Lock) Marshal(appid string, key string)(ret []byte, err error){
	this.Appid = appid;

	return ret, nil;
}