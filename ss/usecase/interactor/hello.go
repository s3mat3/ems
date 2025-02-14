/**
 * @file hello.go
 *
 * @copyright © 2025 s3mat3
 * This code is licensed under the MIT License, see the LICENSE file for details
 *
 * @author s3mat3
 */
package interactor

import(
    "ems/ss/pkg/b64"
    "ems/ss/internal/crypt/rsa_oaep"

    "ems/ss/usecase/port"
    "ems/ss/entity/model"
    "ems/ss/entity/repository"
)

type Hello struct {
    oport port.HelloOut
    repos repository.SessionData
}

func NewHelloUsecase(
    oport port.HelloOut,
    repos repository.SessionData) port.HelloInput{
    return  &Hello {
        oport: oport,
        repos: repos,
    }
}

func (s *Hello) ProvidePubkey() (port.Pubkey, error) {
    key, err := rsa_oaep.ReadPublicKey("public.pem")
    if err != nil {
        return port.Pubkey{}, err
    }
    return s.oport.ResponsePubkey(key), nil
}

func (s *Hello)ProvideSsid(sk port.Himitsu) (port.Ssid, error) {
    dec, err := b64.Decode(sk.Himitsu)
    if err != nil {
        return port.Ssid{}, err
    }
    pk, err := rsa_oaep.ReadPrivateKey("private.pem")
    if err != nil {
        return port.Ssid{}, err
    }
    aesKey, err := rsa_oaep.Decrypt(dec, pk, nil)
    if err != nil {
        return port.Ssid{}, err
    }
    sd := model.NewSessionData()
    sd.SetSecretKey(aesKey)
    ssid := sd.GenarateID()
    s.repos.Update(ssid, *sd)
    return s.oport.ResponseSsid(ssid), nil
}

/* //<-- hello.go ends here */
