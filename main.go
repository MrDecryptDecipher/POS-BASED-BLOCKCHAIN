package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"strconv"
	"time"
)

type block struct {
	// hash of the previous block
	prehash string
	//this block hash
	hash string
	//time stamp
	timestamp string
	// block content
	data string
	// block height
	height int
	//Dig out the address of this block
	address string
}

// used to store the blockchain
var blockchain []block

//represent mining node
type node struct {
	// number of tokens
	tokens int
	// pledge time
	days int
	//node address
	address string
}

//mining node
var mineNodesPool []node

//probability node pool
var probabilityNodesPool []node

//initialization
func init() {
	// Manually add two nodes

	mineNodesPool = append(mineNodesPool, node{1000, 1, "AAAAAAAAAA"})
	mineNodesPool = append(mineNodesPool, node{100, 3, "BBBBBBBBBB"})
	//Initialize the random node pool (the mining probability is related to the number of tokens and the age of tokens)
	for _, v := range mineNodesPool {
		for i := 0; i <= v.tokens*v.days; i++ {
			probabilityNodesPool = append(probabilityNodesPool, v)
		}
	}
}

//generate a new block
func generateNewBlock(oldBlock block, data string, address string) block {
	newBlock := block{}
	newBlock.prehash = oldBlock.hash
	newBlock.data = data
	newBlock.timestamp = time.Now().Format("2006-01-02 15:04:05")
	newBlock.height = oldBlock.height + 1
	newBlock.address = getMineNodeAddress()
	newBlock.getHash()
	return newBlock
}

// hash itself
func (b *block) getHash() {
	sumString := b.prehash + b.timestamp + b.data + b.address + strconv.Itoa(b.height)
	hash := sha256.Sum256([]byte(sumString))
	b.hash = hex.EncodeToString(hash[:])
}

// Randomly get the mining address (the mining probability is related to the number of tokens and the age of the token)
func getMineNodeAddress() string {
	bInt := big.NewInt(int64(len(probabilityNodesPool)))
	// Get a random number, the maximum does not exceed the size of the random node pool
	rInt, err := rand.Int(rand.Reader, bInt)
	if err != nil {
		log.Panic(err)
	}
	return probabilityNodesPool[int(rInt.Int64())].address
}

func main() {
	//create genesis block
	genesisBlock := block{"0000000000000000000000000000000000000000000000000000000000000000", "", time.Now().Format("2006-01-02 15:04:05"), "Created By Sandeep Kumar Sahoo", 1, "0000000000"}
	genesisBlock.getHash()
	//Add the genesis block to the blockchain
	blockchain = append(blockchain, genesisBlock)
	fmt.Println(blockchain[0])
	i := 0
	for {
		time.Sleep(time.Second)
		newBlock := generateNewBlock(blockchain[i], "Sandeep's POS Blockchain", "00000")
		blockchain = append(blockchain, newBlock)
		fmt.Println(blockchain[i+1])
		i++
	}
}
