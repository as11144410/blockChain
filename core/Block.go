package core

import (
        "crypto/sha256"
        "encoding/hex"
        "strconv"
        "time"
)

type Block struct {
        Index         int64  // 区块编号
        Timestamp     int64  // 区块时间戳
        PrevBlockHash string // 上一个区块哈希值
        Hash          string // 当前区块哈下值
        Data          string // 区块数据
}

// hash生成
func calculateHash(b Block) string {
        blockData := strconv.FormatInt(b.Index, 10) + strconv.FormatInt(b.Timestamp, 10) + b.PrevBlockHash + b.Data
        hashInBytes := sha256.Sum256([]byte(blockData))
        return hex.EncodeToString(hashInBytes[:])
}

// GenerateNewBlock 生成新区块
func GenerateNewBlock(preBlock Block, data string) Block {
        newBlock := Block{}
        newBlock.Index = preBlock.Index + 1
        newBlock.PrevBlockHash = preBlock.Hash
        newBlock.Timestamp = time.Now().Unix()
        newBlock.Data = data
        newBlock.Hash = calculateHash(newBlock)
        return newBlock
}

// GenerateGenesisBlock 创始区块
func GenerateGenesisBlock() Block {
        preBlock := Block{}
        preBlock.Index = -1
        preBlock.Hash = ""
        return GenerateNewBlock(preBlock, "Genesis Block")
}