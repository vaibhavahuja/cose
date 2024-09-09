package main

import (
	"github.com/ldclabs/cose/key/hmac"

	"github.com/ldclabs/cose/iana"
	"github.com/ldclabs/cose/key"
	"log"
)

func main() {
	ExampleMac0Message()
}

func ExampleMac0Message() {
	// load key
	k := key.Key{
		iana.KeyParameterKty: iana.KeyTypeSymmetric,
		iana.KeyParameterAlg: iana.AlgorithmHMAC_256_64,
	}
	//macer, err := hmac.New(k)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//tag, err := macer.MACCreate([]byte("hello world"))
	//log.Println(tag)
	k[iana.SymmetricKeyParameterK] = key.GetRandomBytes(32)
	macer, err := hmac.New(k)
	if err != nil {
		log.Fatal(err)
	}

	tag, err := macer.MACCreate([]byte("hello world"))
	log.Println(tag)
}
