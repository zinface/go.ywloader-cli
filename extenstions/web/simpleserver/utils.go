package simpleserver

import (
	"log"
	"strconv"
	"strings"
)

// (k, m, g) or (K, M, G)
func SpeedStringToInt(speed string) int {
	// 检查是否有任何支持的后缀
	hasSuffix := strings.HasSuffix(speed, "m") || strings.HasSuffix(speed, "M") ||
		strings.HasSuffix(speed, "k") || strings.HasSuffix(speed, "K") ||
		strings.HasSuffix(speed, "g") || strings.HasSuffix(speed, "G")

	if !hasSuffix {
		log.Fatalf("speed must have a suffix (m, k, g), got: %s", speed)
	}

	// 处理后缀
	multiplier := 1
	speedValue := speed

	switch {
	case strings.HasSuffix(speed, "k") || strings.HasSuffix(speed, "K"):
		multiplier = 1024 // 1KB = 1024 bytes
		speedValue = strings.TrimSuffix(speed, "k")
		speedValue = strings.TrimSuffix(speedValue, "K")
	case strings.HasSuffix(speed, "m") || strings.HasSuffix(speed, "M"):
		multiplier = 1024 * 1024 // 1MB = 1024 * 1024 bytes
		speedValue = strings.TrimSuffix(speed, "m")
		speedValue = strings.TrimSuffix(speedValue, "M")
	case strings.HasSuffix(speed, "g") || strings.HasSuffix(speed, "G"):
		multiplier = 1024 * 1024 * 1024 // 1GB = 1024^3 bytes
		speedValue = strings.TrimSuffix(speed, "g")
		speedValue = strings.TrimSuffix(speedValue, "G")
	}

	// 转换为整数
	baseValue, err := strconv.Atoi(speedValue)
	if err != nil {
		log.Fatalf("invalid speed format: %s, error: %v", speed, err)
	}

	return baseValue * multiplier
}

func SpeedStringToUint64(speed string) uint64 {
	return uint64(SpeedStringToInt(speed))
}
