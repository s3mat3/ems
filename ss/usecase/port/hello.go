/**
 * @file hello.go
 *
 * @copyright © 2025 s3mat3
 * This code is licensed under the MIT License, see the LICENSE file for details
 *
 * @author s3mat3
 */
package port

type HelloInput interface {
    ProvidePubkey() (Pubkey, error)
    ProvideSsid(Himitsu) (Ssid, error)
}

type HelloOut interface {
    ResponsePubkey(string) Pubkey
    ResponseSsid(string) Ssid
}
/* //<-- hello.go ends here */
