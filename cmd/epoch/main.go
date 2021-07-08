package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"

	"github.com/snowfork/ethashproof"
)

func getHomeDir() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr.HomeDir
}

func main() {

	dataDir := filepath.Join(getHomeDir(), ".ethash")
	err := os.MkdirAll(dataDir, 0755)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	cacheDir := filepath.Join(getHomeDir(), ".ethashproof")
	err = os.MkdirAll(cacheDir, 0755)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}


	for i := 0; i < 512; i++ {
		os.RemoveAll(dataDir)
		root, err := ethashproof.CalculateDatasetMerkleRoot(uint64(i), false, dataDir, cacheDir)
		if err != nil {
			fmt.Printf("Calculating dataset merkle root failed: %s\n", err)
			return
		}
		err = ioutil.WriteFile(
			fmt.Sprintf("%d.txt", i),
			[]byte(root.Hex()),
			0644,
		)
		if err != nil {
			fmt.Printf("Write merkle root to file: %s\n", err)
			return
		}
	}
}
