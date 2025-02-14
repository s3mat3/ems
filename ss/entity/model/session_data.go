/**
 * @file session_data.go
 *
 * @copyright © 2025 s3mat3
 * This code is licensed under the MIT License, see the LICENSE file for details
 *
 * @author s3mat3
 */
package model

import (
    "crypto/sha256"

    "ems/ss/pkg/b64"
    "ems/ss/pkg/gen"
)

type SessionData struct {
    secretKey []byte
    tokenKey []byte
    permission uint
}

func NewSessionData() *SessionData {
    return &SessionData {}
}

func (s *SessionData) GetSecretKey() []byte {
    return s.secretKey
}

func (s *SessionData) GetTokenKey() []byte {
    return s.tokenKey
}

func (s *SessionData) SetSecretKey(nv []byte) {
    s.secretKey = nv
}

func (s *SessionData) SetTokenKey(nv []byte) {
    s.tokenKey = nv
}

func (s *SessionData) GenarateID() string {
    seed := sha256.Sum256(gen.RandomLetters(32))
    ssid := b64.EncodeURLB(seed[:])
    return string(ssid)
}

/* //<-- session_data.go ends here */
