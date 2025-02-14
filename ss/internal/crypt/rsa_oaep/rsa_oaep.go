/**
 * @file rsa_oaep.go
 *
 * @copyright © 2025 s3mat3
 * This code is licensed under the MIT License, see the LICENSE file for details
 *
 * @author s3mat3
 */
package rsa_oaep

import (
    "fmt"
    "os"
    "crypto/rand"
    "crypto/rsa"
    "crypto/sha256"
    "crypto/x509"
    "encoding/pem"
    "errors"
)
// read RSA public key from file for distribute client
func ReadPublicKey(fileName string) (string, error) {
    contents, err := os.ReadFile(fileName)
    if err != nil {
        return "", err
    }
    return string(contents), nil
}

func ReadPrivateKey(pemFile string) (*rsa.PrivateKey, error) {
    bytes, err := os.ReadFile(pemFile)
    if err != nil {
        return nil, err
    }

    block, _ := pem.Decode(bytes)
    if block == nil {
        return nil, errors.New("invalid private key data")
    }

    var key *rsa.PrivateKey
    if block.Type == "RSA PRIVATE KEY" {
        key, err = x509.ParsePKCS1PrivateKey(block.Bytes)
        if err != nil {
            return nil, err
        }
    } else if block.Type == "PRIVATE KEY" {
        keyInterface, err := x509.ParsePKCS8PrivateKey(block.Bytes)
        if err != nil {
            return nil, err
        }
        var ok bool
        key, ok = keyInterface.(*rsa.PrivateKey)
        if !ok {
            return nil, errors.New("not RSA private key")
        }
    } else {
        return nil, errors.New (fmt.Sprintf("invalid private key type : %s", block.Type))
    }

    key.Precompute()

    if err := key.Validate(); err != nil {
        return nil, err
    }

    return key, nil
}

func Decrypt(ciphered []byte, priv *rsa.PrivateKey ,label []byte) ([]byte, error){
    plainBytes, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, priv, ciphered, label)
    if err != nil {
        return nil, errors.New("Fail decrypt")
    }
    return plainBytes, nil
}
/* //<-- rsa_oaep.go ends here */
