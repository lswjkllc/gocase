package main

import (
	"encoding/json"
	"fmt"
	fbs "gocase/rpc/flatbuf/fbs/Block"
	"time"

	flatbuffers "github.com/google/flatbuffers/go"
)

type Block struct {
	Id   int64
	Hash string
	Flag bool
	Txs  []Tx
}

type Tx struct {
	Hash  string
	Value float64
}

func main() {
	TestJson()
	TestFlatbuf()
}

func TestJson() {
	st := time.Now()
	buf := EncodeJson()
	fmt.Println("json encode time:", time.Since(st))
	fmt.Println("json encode size:", len(buf))

	st = time.Now()
	djstr := DecodeToJson(buf)
	fmt.Println("json decode time:", time.Since(st))
	fmt.Println("json    deser block:", djstr)
}

func EncodeJson() []byte {
	txone := Tx{Hash: "adfadf", Value: 123}
	txtwo := Tx{Hash: "adfadf", Value: 456}
	block := &Block{Id: 1, Hash: "aadd", Flag: true, Txs: []Tx{txone, txtwo}}
	buf, _ := json.Marshal(block)
	fmt.Println("json ser bytes:", buf)

	return buf
}

func DecodeToJson(buf []byte) Block {
	var block Block
	json.Unmarshal(buf, &block)
	return block
}

func TestFlatbuf() {
	st := time.Now()
	buf := EncodeBlock()
	fmt.Println("flatbuf encode time:", time.Since(st))
	fmt.Println("flatbuf encode size:", len(buf))

	st = time.Now()
	dblock := DecodeToBlock(buf)
	fmt.Println("flatbuf decode time:", time.Since(st))
	fmt.Println("flatbuf deser block:", dblock)
}

func EncodeBlock() []byte {
	txone := Tx{Hash: "adfadf", Value: 123}
	txtwo := Tx{Hash: "adfadf", Value: 456}
	block := Block{Id: 1, Hash: "aadd", Flag: true}
	//初始化buffer，大小为0，会自动扩容
	builder := flatbuffers.NewBuilder(0)
	//第一个交易
	txoneh := builder.CreateString(txone.Hash) //先处理非标量string,得到偏移量
	fbs.TxStart(builder)
	fbs.TxAddHash(builder, txoneh)
	fbs.TxAddValue(builder, txone.Value)
	ntxone := fbs.TxEnd(builder)
	//第二个交易
	txtwoh := builder.CreateString(txtwo.Hash)
	fbs.TxStart(builder)
	fbs.TxAddHash(builder, txtwoh)
	fbs.TxAddValue(builder, txtwo.Value)
	ntxtwo := fbs.TxEnd(builder)
	//block
	//先处理数组，string等非标量
	fbs.BlockStartTxsVector(builder, 2)
	builder.PrependUOffsetT(ntxtwo)
	builder.PrependUOffsetT(ntxone)
	txs := builder.EndVector(2)
	bh := builder.CreateString(block.Hash)
	//再处理标量
	fbs.BlockStart(builder)
	fbs.BlockAddId(builder, block.Id)
	fbs.BlockAddHash(builder, bh)
	fbs.BlockAddFlag(builder, block.Flag)
	fbs.BlockAddTxs(builder, txs)
	nb := fbs.BlockEnd(builder)
	builder.Finish(nb)
	buf := builder.FinishedBytes() //返回[]byte
	fmt.Println("flatbuf ser bytes:", buf)

	return buf
}

func DecodeToBlock(buf []byte) Block {
	var (
		block Block
	)
	// buf, err := ioutil.ReadFile(filename)
	// if err != nil {
	// 	panic(err)
	// }
	//传入二进制数据
	b := fbs.GetRootAsBlock(buf, 0)
	block.Flag = b.Flag()
	block.Hash = string(b.Hash())
	block.Id = b.Id()
	len := b.TxsLength()
	for i := 0; i < len; i++ {
		tx := new(fbs.Tx)
		ntx := new(Tx)
		if b.Txs(tx, i) {
			ntx.Hash = string(tx.Hash())
			ntx.Value = tx.Value()
		}
		block.Txs = append(block.Txs, *ntx)
	}
	return block
}
