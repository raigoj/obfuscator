package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/binary"
	"io"
	"log"
	"os"

	"github.com/amenzhinsky/go-memexec"
)

const (
	META = 42
)

var (
	nonce []byte = make([]byte, 12)
)

func main() {
	if obf, meta := checkMeta(); obf {
		uncrypt(meta)
		os.Exit(0)
	}
	encrypt()
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func fileInfo(f string) ([]byte, int64) {
	bytes, err := os.ReadFile(f)
	check(err)
	size := int64(len(bytes))
	return bytes, size
}

func execute(b []byte) {
	bin, err := memexec.New(b)
	check(err)
	cmd := bin.Command()
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	check(cmd.Run())
}

func getMeta(off uint64, size uint64) []byte {
	meta := make([]byte, META)
	binary.BigEndian.PutUint64(meta[10:20], off)
	binary.BigEndian.PutUint64(meta[20:30], size)
	copy(meta[0:10], []byte("OBFUSCATED"))
	copy(meta[30:42], nonce)
	return meta
}

func checkMeta() (bool, []byte) {
	ex, err := os.Executable()
	check(err)
	self, err := os.OpenFile(ex, os.O_RDONLY, 0600)
	check(err)
	bytes := make([]byte, META)
	self.Seek(-META, 2)
	self.Read(bytes)
	if string(bytes[0:10]) == "OBFUSCATED" {
		return true, bytes
	}
	return false, nil
}

func crypt(bytes []byte, isEncrypt bool) []byte {
	aesCipher, err := aes.NewCipher([]byte{0x58, 0x8e, 0x79, 0x6b, 0x6a, 0x21, 0xc3, 0x9d, 0xa0, 0xfc, 0x3b, 0x40, 0xed, 0x51, 0x40, 0xb0, 0x49, 0xe2, 0x68, 0x5f, 0x22, 0xba, 0x2b, 0x2a, 0xb4, 0xcb, 0xc7, 0x10, 0x2b, 0xa9, 0xc1, 0xd3})
	check(err)
	aesGCM, err := cipher.NewGCM(aesCipher)
	check(err)
	if isEncrypt {
		_, err = io.ReadFull(rand.Reader, nonce)
		check(err)
		return aesGCM.Seal(nil, nonce, bytes, nil)
	}
	dec, err := aesGCM.Open(nil, nonce, bytes, nil)
	check(err)
	return dec
}

func encrypt() {
	ex, err := os.Executable()
	check(err)
	bytes, size := fileInfo(ex)
	target, _ := fileInfo("bin")
	enc := crypt(target, true)
	meta := getMeta(uint64(size), uint64(len(enc)))
	out := append(bytes, enc...)
	out = append(out, meta...)
	err = os.WriteFile("newbin", out, 0700)
	check(err)
}

func uncrypt(meta []byte) {
	off := binary.BigEndian.Uint64(meta[10:20])
	size := binary.BigEndian.Uint64(meta[20:30])
	ex, err := os.Executable()
	check(err)
	self, err := os.OpenFile(ex, os.O_RDONLY, 0600)
	check(err)
	enc := make([]byte, size)
	self.Seek(int64(off), 0)
	self.Read(enc)
	nonce = meta[30:42]
	dec := crypt(enc, false)
	bytes, _ := os.ReadFile(ex)
	enc = crypt(dec, true)
	copy(bytes[off:off+uint64(len(enc))], enc)
	copy(bytes[len(bytes)-12:], nonce)
	os.Remove(ex)
	os.WriteFile(ex, bytes, 0700)
	execute(dec)
}
