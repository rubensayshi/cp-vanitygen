/*
counterparty compatible mnemonic, different from BIP39
 */
package mnemonic_test

import (
	"testing"
	"github.com/rubensayshi/cp-vanitygen/mnemonic"
	"github.com/stretchr/testify/assert"
	"encoding/hex"
)

func TestMnemonicFromHex(t *testing.T) {
	srcHexStr := "00fefe00fefe00fe"
	m := mnemonic.MnemonicFromSeedHex(srcHexStr)

	hexStr, err := m.Hex()
	assert.Nil(t, err)
	assert.Equal(t, srcHexStr, hexStr, "")

	words, err := m.Words()
	assert.Nil(t, err)
	assert.Equal(t, "sway bump guilt hatred scratch lace", words)
}

func TestMnemonicFromSeed(t *testing.T) {
	srcHexStr := "00fefe00fefe00fe"
	srcSeed, err := hex.DecodeString(srcHexStr)

	m := mnemonic.MnemonicFromSeed(srcSeed)

	hexStr, err := m.Hex()
	assert.Nil(t, err)
	assert.Equal(t, srcHexStr, hexStr, "")

	words, err := m.Words()
	assert.Nil(t, err)
	assert.Equal(t, "sway bump guilt hatred scratch lace", words)
}
