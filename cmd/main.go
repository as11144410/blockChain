package main

import "demochain/core"

func main() {
        bc := core.NewBlockChain()
        bc.SendData("send 1 BTC to YST")
        bc.SendData("send 1 EOS to YST")

        bc.Print()
}
