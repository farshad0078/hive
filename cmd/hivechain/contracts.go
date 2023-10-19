package main

import _ "embed"

//go:embed bytecode/gencode.bin
var gencodeCode []byte

//go:embed bytecode/genlogs.bin
var genlogsCode []byte

//go:embed bytecode/genstorage.bin
var genstorageCode []byte

//go:embed bytecode/deployer.bin
var deployerCode []byte

//go:embed bytecode/callme.bin
var callmeCode []byte

//go:embed bytecode/callenv.bin
var callenvCode []byte

// //go:embed bytecode/deposit.bin
// var depositCode []byte
// 
// const depositContractAddr = "0x00000000219ab540356cBB839Cbe05303d7705Fa"
