/*
counterparty compatible mnemonic, different from BIP39
 */
package mnemonic

import (
	"strings"
	"encoding/hex"
	"strconv"
)

type Mnemonic struct {
	seed []byte
}

func (m *Mnemonic) Words() (string, error) {
	entropy := m.seed

	sentenceLength := len(entropy) / 4

	// Slice to hold words in
	words := make([]string, sentenceLength * 3)

	n := uint32(len(WordList))

	for i := 0; i < sentenceLength; i++ {
		b := append([]byte{}, entropy[i * 4], entropy[(i * 4) + 1], entropy[(i * 4) + 2], entropy[(i * 4) + 3])

		x64, err := strconv.ParseInt(hex.EncodeToString(b), 16, 64)
		if err != nil {
			panic(err)
		}
		x := uint32(x64)

		w1 := x % n;
		w2 := (((x / n) >> 0) + w1 ) % n;
		w3 := (((((x / n) >> 0) / n ) >> 0) + w2 ) % n;

		words[i * 3] = WordList[w1]
		words[i * 3 + 1] = WordList[w2]
		words[i * 3 + 2] = WordList[w3]
	}

	return strings.Join(words, " "), nil
}

func (m *Mnemonic) Hex() (string, error) {
	entropy := m.seed

	sentenceLength := len(entropy) / 4

	// Slice to hold words in
	words := make([]string, sentenceLength)

	for i := 0; i < sentenceLength; i++ {
		b := append([]byte{}, entropy[i * 4], entropy[(i * 4) + 1], entropy[(i * 4) + 2], entropy[(i * 4) + 3])

		words = append(words, hex.EncodeToString(b))
	}

	return strings.Join(words, ""), nil
}

func MnemonicFromSeed(seed []byte) *Mnemonic {
	return &Mnemonic{seed}
}

func MnemonicFromSeedHex(seedHex string) *Mnemonic {
	seed, err := hex.DecodeString(seedHex)
	if err != nil {
		panic(err)
	}

	return &Mnemonic{seed}
}
