package utils

import (
	"crypto/md5"
	"encoding/hex"
	"go.uber.org/zap"
	"io"
)

func GetMD5(file io.Reader) string {
	m := md5.New()
	_, err := io.Copy(m, file)
	if err != nil {
		zap.S().Errorf("failed to copy file to hash: %v", err)
		return ""
	}
	sum := m.Sum(nil)
	return hex.EncodeToString(sum)
}
