package main

import (
	"crypto/sha1"
	"encoding/hex"
)

type Hash [20]byte

func (h Hash) String() string {
	return hex.EncodeToString(h[:])
}

func ComputeHash(data []byte) Hash {
	return sha1.Sum(data)
}

type Hashable interface {
	hash() Hash
}

type Block string

func (b Block) hash() Hash {
	return ComputeHash([]byte(b))
}

type EmptyBlock struct{}

func (EmptyBlock) hash() Hash {
	return [20]byte{}
}
