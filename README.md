
>The proof-of-stake mechanism was originally proposed and applied by PeerCoin (block generation probability = token quantity * coin age). Simply put, whoever has more coins has a greater probability of block generation. But digging deeper, who will calculate the probability of block generation? What should I do if I encounter the problem of no-cost benefit relationship? This consensus algorithm is very simple at first glance, but there are actually many problems to be solved. Let’s see when Ethereum can completely switch to the POS mechanism

<br>

block structure
```go
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
```
Declare two node pools\
mineNodesPool Used to store specified mining nodes\
probabilityNodesPool  The number of tokens used to deposit into the mining node * the probability of obtaining the currency age
```go
// used to store the blockchain
var blockchain []block
//represent mining node
type node struct{
	// number of tokens
	tokens int
	// pledge time
	days  int
	//node address
	address string
}
//mining node
var mineNodesPool []node
//probability node pool

var  probabilityNodesPool []node
```
Initialize the node pool:
```go
func init () {
	// Manually add two nodes
	mineNodesPool = append(mineNodesPool,node{1000,1,"AAAAAAAAAA"})
	mineNodesPool = append(mineNodesPool,node{100,3,"BBBBBBBBBB"})
	//Initialize the random node pool (the mining probability is related to the number of tokens and the age of tokens)
	for _,v:=range mineNodesPool{
		for i:=0;i<=v.tokens * v.days; i ++ {
			randNodesPool = append(randNodesPool,v)
		}
	}
}
```
Each mining will randomly select the address of the node that obtains the block right from the probability node pool
```go
// Randomly get the mining address (the mining probability is related to the number of tokens and the age of the token)
func getMineNodeAddress() string{
	bInt:=big.NewInt(int64(len(randNodesPool)))
	/ / Get a random number, the maximum does not exceed the size of the random node pool
	rInt,err:=rand.Int(rand.Reader,bInt)
	if err != nil {
		log.Panic(err)
	}
	return randNodesPool[int(rInt.Int64())].address
}
```


```go
func main() {
	//create genesis block
	genesisBlock := block{"0000000000000000000000000000000000000000000000000000000000000000","",time.Now().Format("2006-01-02 15:04:05"),"I am the genesis block",1,"0000000000"}
	genesisBlock.getHash()
	//Add the genesis block to the blockchain
	blockchain = append(blockchain,genesisBlock)
	fmt.Println(blockchain[0])
	i:=0
	for  {
		time.Sleep(time.Second)
		newBlock:=generateNewBlock(blockchain[i],"I am block content","00000")
		blockchain = append(blockchain,newBlock)
		fmt.Println(blockchain[i + 1])
		i++
	}
}
```
