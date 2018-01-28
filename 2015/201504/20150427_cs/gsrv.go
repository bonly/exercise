package main

import (
	quic "github.com/lucas-clemente/quic-go"
	"fmt"
	"log"
	// "crypto/rand"
	// "crypto/rsa"
	// "crypto/tls"
	// "crypto/x509"
	// "encoding/pem"
	// "math/big"
)

const addr = "127.0.0.1:4242";

func main(){
	log.Fatal(srv());
}

func srv() error{
	// var cfg quic.Config;
	// cfg.KeepAlive = true;

	// lst, err := quic.ListenAddr(addr, &tls.Config{InsecureSkipVerify: true}, &quic.Config{});
	lst, err := quic.ListenAddr(addr, &quic.Config{});
	if err != nil{
		return err;
	}

	sess, err := lst.Accept();
	if err != nil{
		return err;
	}
	fmt.Printf("sess[%s]: %#v\n", sess.RemoteAddr(), sess);

	stream, err := sess.AcceptStream();
	if err != nil{
		panic(err);
	}
	fmt.Printf("stream: %#v\n", stream);

	stream.Write([]byte("abc"));
	// _, err = io.Copy(loggingWriter{stream}, stream)
	return err;
}

// type loggingWriter struct{ io.Writer }

// func (w loggingWriter) Write(b []byte) (int, error) {
// 	fmt.Printf("Server: Got '%s'\n", string(b))
// 	return w.Writer.Write(b)
// }

// func generateTLSConfig() *tls.Config {
// 	key, err := rsa.GenerateKey(rand.Reader, 1024)
// 	if err != nil {
// 		panic(err)
// 	}
// 	template := x509.Certificate{SerialNumber: big.NewInt(1)}
// 	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &key.PublicKey, key)
// 	if err != nil {
// 		panic(err)
// 	}
// 	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)})
// 	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER})

// 	tlsCert, err := tls.X509KeyPair(certPEM, keyPEM)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return &tls.Config{Certificates: []tls.Certificate{tlsCert}}
// }