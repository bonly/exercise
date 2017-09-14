package main 

import(
// "crypto"
"crypto/rsa"
"crypto/sha256"
"crypto/rand"
// "crypto/cipher"
"crypto/x509"
"encoding/pem"
"fmt"
"os"
"log"
)

func main(){
	pem_data := []byte(
`-----BEGIN RSA PRIVATE KEY-----
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
//`ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCrbF1yRWn54DoNLCHv6+XcC76/onLFUvn/GQLWXj+lsC2BQ+vTmQHjNbJ5Hvg+mJihGOBzyZnua/hXt8Pfeyut+eXL9ReY7VW9uiheiUSndjF89dqC0qbFe+aAC+ME+TITSU2TMFA7yIUDZxTXZ9RFS7CQhw42Gp2A6DyXsrPDVnIZn9q5lb9w5Qky8uKfjebKQKIgBWxN7PtViWLpynLeiuC3xKiLcI5Ot6iEQxvlwghl85z6Qgmw92BIqRlfcGyW2W3KDM/1bPZffqTFecO3XduesRDeIUGzhG1z3WsiTs2K60jvgS/Fwx0l38fiLie73G+7jJAYBQQ9Z7oUCZ2P bonly@163.com`);

	block, _ := pem.Decode(pem_data); 
	if block == nil || block.Type != "RSA PRIVATE KEY" {
        log.Fatal("No valid PEM data found")
    }

	private_key, err := x509.ParsePKCS1PrivateKey(block.Bytes); 
	if err != nil {
        log.Fatalf("Private key can't be decoded: %s", err);
    }

    public_key := &private_key.PublicKey;

	message := []byte("this is bonly's test");

	encode, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, public_key, message, []byte("label"));
	if err != nil{
		fmt.Println(err);
		os.Exit(1);
	}

	fmt.Printf("cd: %s\n", string(encode));
}