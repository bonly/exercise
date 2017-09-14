// https://blogs.oracle.com/janp/entry/how_the_scp_protocol_works
package main

import (
	"fmt"

	"golang.org/x/crypto/ssh"
)

const privateKey = `
-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEAq2xdckVp+eA6DSwh7+vl3Au+v6JyxVL5/xkC1l4/pbAtgUPr
05kB4zWyeR74PpiYoRjgc8mZ7mv4V7fD33srrfnly/UXmO1VvbooXolEp3YxfPXa
gtKmxXvmgAvjBPkyE0lNkzBQO8iFA2cU12fURUuwkIcONhqdgOg8l7Kzw1ZyGZ/a
uZW/cOUJMvLin43mykCiIAVsTez7VYli6cpy3orgt8Soi3COTreohEMb5cIIZfOc
+kIJsPdgSKkZX3BsltltygzP9Wz2X36kxXnDt13bnrEQ3iFBs4Rtc91rIk7NiutI
74EvxcMdJd/H4i4nu9xvu4yQGAUEPWe6FAmdjwIDAQABAoIBAQCYLLCfumMjgRvH
ib/uvhjlSD/i2pDL/qcS/9728d/KdCVRNTxqMM/uJUL0Jrv0xX3otK672uGdN5fU
NrvY7fyOiuKmNkLmbPnKXTFtCqo5hsDTz2RU8i897IlobiTvy0/asx38Zc0z5QO/
S4jC19cmVNS+2EWTXfkn8AGqIOdXpTa+ZAPvldDEeRABALqzASBAFaf9ra9wV58Z
LV63+AYWufrkh8cb7BL1mzRh6l6jTcblq3LN7Qm5BCZ7byiDPjTYlCtwANFkslcv
UUkCSi3F8AWM1bklGNd5G3BVIi9BNqVyagcXxMegaLI9+exwhk9Ja4bwoJM2fPF5
m3ohAydBAoGBAOKLRFr/D3kn9WUhar90fgeEQVM8xwx5f1p37HnvwgvBGfs6Zqba
qGh56V0LUBBh+DXqwOllLbATP+YdO8zRlMKS8qRVDt6Oo4hoGNI1aZ7EzOWdmCpI
0cpeh4y5MspwuRKjso7ghl9/lByxTYBVPu7k/dTKcR+I7digKOzzjGu/AoGBAMG2
Wkq09rELHyBwHf6zJWbSyhK4lrNopJTm9QNa3MBtw3lUECrF4WJK6uxTEYKYyyQy
cTE7HuULG8MTUKCCh09o6zstY6vnzfHctlrrt0Iac406ycrbmcaU1z2AmRhiWo1W
hBjb9p5Mm0o+yBNaMiI22FZv/Lw0O4+YXZh9DoIxAoGALW6Og8049iwsS167QLAF
Ak7kpuO+a2MGRkdclkud16ufIUHiCj45ndItGarGILL1CyMMEwJmejQyEVz1fCrf
sXG01X31YG6snxN5PtbVWrDPApFrTbeS0wnIF1GgvPaENxe7HGosqIQ2WNa52y29
VD16Ji1/KDrWsCXWb0acd5MCgYEAjyyvqd39I7gbpocXtGRGtawAzTt/r1UdqCzY
VUV4OgYm32tBk8u0HUlQP6HaJFO4eaWKwh93Y0UFnPaOtkQjrI7YpmGS9MEdF7HR
Wnw0c/hHBdC4y2XqT7s9J2kAPuSbFzIl6pXRvwjSEls62ArnWSB+X8zf2V8l0qmE
LCvqaFECgYAY6ORS9vqFkD59Mz8EV4vdn8fbnc+0pzJqLJcngax3I4/D+RN0Ki+M
eYRpskmjud4YWUCiqmvikf/m/+6z82lyoP4oAaieDdbI4KvjXqs+S8bbCjTSyymz
XJMjSGu8KMsBovXMQmKv0VnW5YJ8PP6OaKBN4DafCYUwVMYbCZiLmw==
-----END RSA PRIVATE KEY-----`;

func main() {
	signer, _ := ssh.ParsePrivateKey([]byte(privateKey));
	clientConfig := &ssh.ClientConfig{
		User: "apps",
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
	};

	client, err := ssh.Dial("tcp", "120.25.106.243:22", clientConfig);
	if err != nil {
		panic("Failed to dial: " + err.Error());
	}

	session, err := client.NewSession();
	if err != nil {
		panic("Failed to create session: " + err.Error());
	}
	defer session.Close();

	go func() {
		w, _ := session.StdinPipe();
		defer w.Close()
		content := "123456789\n"
		fmt.Fprintln(w, "D0755", 0, "testdir") // mkdir
		fmt.Fprintln(w, "C0644", len(content), "testfile1")
		fmt.Fprint(w, content)
		fmt.Fprint(w, "\x00") // transfer end with \x00
		fmt.Fprintln(w, "C0644", len(content), "testfile2")
		fmt.Fprint(w, content)
		fmt.Fprint(w, "\x00")
	}()

	if err := session.Run("/usr/bin/scp -tr ./"); err != nil { //开终端接收上面的信息处理
		panic("Failed to run: " + err.Error())
	}
}