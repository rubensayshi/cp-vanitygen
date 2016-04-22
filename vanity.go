package main

import (
	"fmt"
	"strings"
	"encoding/binary"
	"sync"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/rubensayshi/cp-vanitygen/mnemonic"
	"time"
	"flag"
)

var MAX_UINT16 = uint16(65535)

type Child struct {
	i uint32;
	hardened bool;
}

type findProgress struct {
	threadNum int
	loops int64
}

type findResult struct {
	threadNum int
	masterSeed []byte
}

func derive(extKey *hdkeychain.ExtendedKey, children []Child) (*hdkeychain.ExtendedKey, error) {
	var err error;
	for _, child := range children {
		i := child.i;
		if child.hardened {
			i += hdkeychain.HardenedKeyStart;
		}

		extKey, err = extKey.Child(i)
		if err != nil {
			return nil, err
		}
	}

	return extKey, nil;
}

func addressFromMasterSeed(masterSeed []byte, network  *chaincfg.Params) (string, error) {
	masterKey, err := hdkeychain.NewMaster(masterSeed, network)
	if err != nil {
		return "", err
	}

	// derive m/0'/0/0
	extKey, err := derive(masterKey, []Child{Child{0, true}, Child{0, false}, Child{0, false}})
	if err != nil {
		return "", err
	}

	address, err := extKey.Address(network);
	if err != nil {
		return "", err
	}

	return address.String(), nil
}

func findPattern(patterns []string, network *chaincfg.Params, wg *sync.WaitGroup, resultChan chan<- interface{}, quitChan <-chan bool, threadNum int) {
	seedLen := 16
	uintLen := 4
	baseSeedLen := seedLen - uintLen

	var c int64 = 0

	for {
		// generate random base seed
		baseSeed, err := hdkeychain.GenerateSeed(uint8(seedLen))
		if err != nil {
			panic(err)
		}

		// slice base seed for space for uint16
		baseSeed = baseSeed[0:baseSeedLen]

		// loop until we reach max unit16
		var i uint16 = 1;
		for ; i <= MAX_UINT16; i++ {

			// check quitChain
			select {
			case <-quitChan:
				wg.Done()
				return
			default:
				// pass
			}

			// turn i into a byte array
			b := make([]byte, uintLen)
			binary.LittleEndian.PutUint16(b, i)

			// append base seed with i byte array
			var masterSeed []byte
			masterSeed = append(baseSeed, b...);

			// derive and create address
			addressstr, err := addressFromMasterSeed(masterSeed, network)
			if err != nil {
				panic(err)
			}

			// check patterns
			for _, pattern := range patterns {
				if (strings.HasPrefix(strings.ToLower(addressstr), pattern)) {
					// fmt.Printf("[%d] %s %x \n", i, addressstr, masterSeed)
					// copy seed to avoid overwriting
					resultSeed := make([]byte, len(masterSeed))
					copy(resultSeed, masterSeed)
					resultChan <- findResult{threadNum: threadNum, masterSeed: resultSeed}
				}
			}

			// increment total loops counter
			c++

			// report progress every 1000 loops
			if i % 1000 == 0 {
				resultChan <- findProgress{threadNum: threadNum, loops: c}
			}
		}
	}
}


func main() {
	threads := flag.Int("threads", 2, "threads to use")

	flag.Parse()

	prefix := strings.ToLower(flag.Arg(0))

	if prefix == "" {
		panic("no prefix")
	}

	network := &chaincfg.MainNetParams
	patterns := []string{prefix}

	start := time.Now()
	wg := sync.WaitGroup{}
	resultsChan := make(chan interface{}, *threads * 3)
	quitChan := make(chan bool, *threads)
	progress := make(map[int]int64)
	for i := 0; i < *threads; i++ {
		wg.Add(1)
		go findPattern(patterns, network, &wg, resultsChan, quitChan, i)
	}

out:
	for {
		select {
		case rawMsg := <-resultsChan:
			switch msg := rawMsg.(type) {

			// record progress
			case findProgress:
				progress[msg.threadNum] = msg.loops

				var totalLoops int64 = 0
				for _, loops := range progress {
					totalLoops += loops
				}

				if totalLoops % 10000 == 0 {
					elapsed := time.Since(start)

					fmt.Printf("%d loops in %.2f sec: %.2f loops/sec \n", totalLoops, elapsed.Seconds(), float64(totalLoops) / elapsed.Seconds())
				}

			// found a match
			case findResult:
				fmt.Println("------------------------")
				fmt.Printf("thread: %d \n", msg.threadNum)
				fmt.Printf("seed: %x \n", msg.masterSeed)

				address, err := addressFromMasterSeed(msg.masterSeed, network)
				if err != nil {
					panic(err)
				}
				fmt.Printf("address: %s \n", address)

				m := mnemonic.MnemonicFromSeed(msg.masterSeed)

				words, err := m.Words()
				if err != nil {
					panic(err)
				}

				h, err := m.Hex()
				if err != nil {
					panic(err)
				}

				fmt.Printf("mnemonic: %s \n", words)
				fmt.Printf("mnemonic HEX: %s \n", h)

				for t := 0; t < *threads; t++ {
					quitChan <- true
				}

				break out

			default:
				panic(fmt.Sprintf("Invalid message type in processHandler: %T", msg))
			}
		}
	}

	wg.Wait()
	fmt.Println("DONE!")
}
