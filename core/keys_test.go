package core

import (
    "crypto/elliptic"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/suite"
    "math/big"
    "testing"
)

type KeysTestSuite struct {
    suite.Suite
}

func (suite *KeysTestSuite) TestCreation() {
    u := User{Name: "test.user", CryptoSalt: "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=", SigningSalt: "AQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQE="}
    k, err := MakeKeys(&u, "password")
    assert.NoError(suite.T(), err, "Key generation failed")
    assert.NotNil(suite.T(), k, "Nil keys")

    cryptoKey := []byte{
        0x92, 0x99, 0xE5, 0xD9, 0x06, 0x87, 0x46, 0xCB,
        0x45, 0x3F, 0x5D, 0x09, 0x16, 0x60, 0xB9, 0xCA,
        0x7A, 0xA1, 0xD9, 0x05, 0xE1, 0x2F, 0x75, 0x3F,
        0x49, 0xD9, 0x08, 0x5A, 0xE6, 0x26, 0xAB, 0xA0,
    }
    assert.Exactly(suite.T(), k.CryptoKey, cryptoKey, "Cryptographic key does not match")

    signingKeyBytes := []byte{
        0x01, 0x68, 0x18, 0x80, 0xEA, 0x2B, 0x6E, 0x5E,
        0x43, 0xA1, 0x2B, 0x99, 0x04, 0x1C, 0x8A, 0xF5,
        0xD9, 0xA0, 0x64, 0xA3, 0x49, 0x50, 0x31, 0x6D,
        0x05, 0x0A, 0xB3, 0xCE, 0xBD, 0x85, 0x84, 0x3B,
        0x3F, 0x69, 0x5F, 0x2C, 0xD9, 0xFF, 0xE1, 0xA3,
        0xBB, 0x5F, 0x16, 0xAB, 0x41, 0x5B, 0x02, 0xB4,
        0x2F, 0x57, 0xA1, 0x08, 0xE2, 0x89, 0xDA, 0x5B,
        0xF9, 0x70, 0x18, 0x44, 0xD2, 0x8D, 0x59, 0xFE,
        0x23, 0x12,
    }
    signingKey := new(big.Int).SetBytes(signingKeyBytes)
    p := k.PublicSigningKey()
    assert.Exactly(suite.T(), p.Curve, elliptic.P521(), "Signing key elliptic curve differs")
    assert.Exactly(suite.T(), k.SigningKey.D, signingKey, "Signing key does not match")
}

func TestKeysTestSuite(t *testing.T) {
    suite.Run(t, new(KeysTestSuite))
}