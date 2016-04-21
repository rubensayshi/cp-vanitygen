package mnemonic

import (
	"strings"
	"fmt"
	"encoding/hex"
	"strconv"
)

var DEBUG_ENABLED = false

func DEBUG(v interface{}) {
	if DEBUG_ENABLED {
		fmt.Println(v)
	}
}

type Mnemonic struct {
	seed []byte
}

func (m *Mnemonic) Words() (string, error) {
	entropy := m.seed

	sentenceLength := len(entropy) / 4

	DEBUG(entropy)
	DEBUG(len(entropy))

	// Slice to hold words in
	words := make([]string, sentenceLength * 3)

	n := uint32(len(WordList))

	for i := 0; i < sentenceLength; i++ {
		DEBUG("--------------")

		b := append([]byte{}, entropy[i * 4], entropy[(i * 4) + 1], entropy[(i * 4) + 2], entropy[(i * 4) + 3])
		DEBUG(b)

		DEBUG(hex.EncodeToString(b))
		x64, err := strconv.ParseInt(hex.EncodeToString(b), 16, 64)
		if err != nil {
			panic(err)
		}
		// x := binary.LittleEndian.Uint32(b)
		DEBUG(x64)

		x := uint32(x64)

		w1 := x % n;
		w2 := (((x / n) >> 0) + w1 ) % n;
		w3 := (((((x / n) >> 0) / n ) >> 0) + w2 ) % n;

		DEBUG(w1)
		DEBUG(WordList[w1])
		DEBUG(w2)
		DEBUG(WordList[w2])
		DEBUG(w3)
		DEBUG(WordList[w3])

		words[i * 3] = WordList[w1]
		words[i * 3 + 1] = WordList[w2]
		words[i * 3 + 2] = WordList[w3]
	}

	return strings.Join(words, " "), nil
}

func (m *Mnemonic) Hex() (string, error) {
	entropy := m.seed

	sentenceLength := len(entropy) / 4

	DEBUG(entropy)
	DEBUG(len(entropy))

	// Slice to hold words in
	words := make([]string, sentenceLength)

	for i := 0; i < sentenceLength; i++ {
		DEBUG("--------------")
		DEBUG(i)

		b := append([]byte{}, entropy[i * 4], entropy[(i * 4) + 1], entropy[(i * 4) + 2], entropy[(i * 4) + 3])
		DEBUG(b)

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

func padByteSlice(slice []byte, length int) []byte {
	newSlice := make([]byte, length-len(slice))
	return append(newSlice, slice...)
}
