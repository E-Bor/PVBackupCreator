package internal

import (
	"awesomeProject/internal/utils"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
)

type HashChecker struct {
	hashes       []string
	hashFilePath string
}

func NewHashChecker(hashRootDir string, hashFileName string) HashChecker {
	if hashRootDir == "" {
		hashFileName = "hashes.json"
	}
	hashFilePath := fmt.Sprintf("%s/%s", hashRootDir, hashFileName)
	hashes := loadHashes(hashFilePath)
	return HashChecker{hashes: hashes, hashFilePath: hashFilePath}
}

func (h HashChecker) CheckOrCreateHash(filePath string) bool {
	newFileHash := h.getFileHash(filePath)

	if utils.ContainsElement(h.hashes, newFileHash) {
		return false
	}
	h.hashes = append(h.hashes, newFileHash)
	return true
}

func (h HashChecker) SaveHashes() {
	hashesToSave, err := json.Marshal(h.hashes)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(h.hashFilePath, hashesToSave, os.ModePerm)
}

func loadHashes(hashFilePath string) []string {
	data, err := os.ReadFile(hashFilePath)
	if err != nil {
		return []string{}
	}
	var hashes []string
	err = json.Unmarshal(data, &hashes)
	return hashes
}

func (h HashChecker) getFileHash(filePath string) string {
	data, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:])
}
