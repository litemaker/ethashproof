package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strconv"

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
	if len(os.Args) < 2 {
		fmt.Printf("Epoch number param is missing. Please run ./cache <epoch_number> instead.\n")
		return
	}
	if len(os.Args) > 2 {
		fmt.Printf("Please pass only 1 param as a epoch number. Please run ./cache <epoch_number> instead.\n")
		return
	}
	number, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("Please pass a number as epoch number. Please run ./cache <integer> instead.\n")
		fmt.Printf("Error: %s\n", err)
		return
	}

	dataDir := filepath.Join(getHomeDir(), ".ethash")
	err = os.MkdirAll(dataDir, 0755)
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

	root, err := ethashproof.CalculateDatasetMerkleRoot(uint64(number), true, dataDir, cacheDir)
	if err != nil {
		fmt.Printf("Calculating dataset merkle root failed: %s\n", err)
		return
	}

	fmt.Printf("Root: %s\n", root.Hex())
}
