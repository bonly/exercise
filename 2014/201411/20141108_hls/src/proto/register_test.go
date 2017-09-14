package proto

import (
"testing"
)

func Test_register(ts *testing.T){
	var srv SRV;
	srv.pack = []byte{};
	srv.srv(ts);
}

