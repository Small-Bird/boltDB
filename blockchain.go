package main

type Blockchain struct {
	blocks []*Block
}

func (bc *Blockchain) AddBlock(data string) {
	preBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, preBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewFirstBlock()}}
}
