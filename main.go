package main

import "errors"

type MerkleTreeHash struct {
}

func NewMerkleTreeHash() {

}

func (mth *MerkleTreeHash) FindMerkleTreeHash(fileHashes []string) error {
	// need to find merkle tree hash of all file hashes

	if fileHashes == nil || len(fileHashes) == 0 {
		return errors.New("no file hashes found")
	}

}

func main() {

}
