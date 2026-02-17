package cluster

import "github.com/DumbNoxx/testing-version-go/internal/processor/sanitizer"

func Cluster(log []byte, idLog string) []byte {
	text := sanitizer.Sanitizer(log, idLog)
	normalizeLog := NormalizeLog(text)

	return normalizeLog
}
