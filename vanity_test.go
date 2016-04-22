package main

import (
	"testing"
	"github.com/btcsuite/btcutil/hdkeychain"
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"github.com/btcsuite/btcd/chaincfg"
)

func TestDerive(t *testing.T) {
	masterSeedHex := "beae474043592e130f4f4bc6de4ec5ac"
	masterSeed, _ := hex.DecodeString(masterSeedHex)
	network := &chaincfg.MainNetParams

	masterKey, err := hdkeychain.NewMaster(masterSeed, network)
	assert.Nil(t, err)
	assert.IsType(t, (*hdkeychain.ExtendedKey)(nil), masterKey)
	assert.Equal(t, "xprv9s21ZrQH143K2QbUBYV9aoDSYgQq9yTkHA4tyX7H5HqhRp9x3EJohgjQzbpeiAbsi8NtGhusaKrnb2iFYFNddrfFTQxV7juj2wWbYem2MVb", masterKey.String())

	childKey, err := derive(masterKey, []Child{Child{0, true}})
	assert.Nil(t, err)
	assert.IsType(t, (*hdkeychain.ExtendedKey)(nil), childKey)
	assert.Equal(t, "xprv9uW2EdoBd2QryEpaQtWkda2YMfsjdjxRD8yBShngyc32TF3VFWGxDs5Uoej4hkp2Fim1zH4yyVrXgT9zbvzqoWKXNYK4ix2UJPubgbxq91Q", childKey.String())

	childKey, err = derive(masterKey, []Child{Child{0, false}})
	assert.Nil(t, err)
	assert.IsType(t, (*hdkeychain.ExtendedKey)(nil), childKey)
	assert.Equal(t, "xprv9uW2Edo3HMsto9odkLxswve5gkKJGgNF1D34eWFzW9FZ52rxeLh9V88JAiu3FyjWGV6ESUAosGdgk5tjdWedBSybCTxrcR1My93B8CR2txm", childKey.String())

	childKey, err = derive(masterKey, []Child{Child{0, true}, Child{0, false}, Child{0, false}})
	assert.Nil(t, err)
	assert.IsType(t, (*hdkeychain.ExtendedKey)(nil), childKey)
	assert.Equal(t, "xprv9yvnmYdUWezWkszKrKvkCGfvxuTugMsjqyLx6HxSSMnGAgqxywjtjYNy2stDKAFMbP5VNSUkDitL9oY4GC3jeSJ3GdG7Z4TB3724G5AQMqX", childKey.String())
	address, err := childKey.Address(network)
	assert.Nil(t, err)
	assert.Equal(t, "1EdwtsD7dgcoEG8UgQ2h5UhZV14p4arKaE", address.String())

	addressStr, err := addressFromMasterSeed(masterSeed, network)
	assert.Nil(t, err)
	assert.Equal(t, "1EdwtsD7dgcoEG8UgQ2h5UhZV14p4arKaE", addressStr)



}
