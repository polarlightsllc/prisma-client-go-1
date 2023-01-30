package unpack

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/polarlightsllc/prisma-client-go-1/binaries"
	"github.com/polarlightsllc/prisma-client-go-1/binaries/platform"
	"github.com/polarlightsllc/prisma-client-go-1/logger"
)

// TODO check checksum after expanding file

// noinspection GoUnusedExportedFunction
func Unpack(data []byte, name string, version string) {
	start := time.Now()

	file := fmt.Sprintf("prisma-query-engine-%s", name)

	// TODO check if dev env/dev binary in ~/.prisma
	// TODO check if engine in local dir OR env var

	tempDir := binaries.GlobalUnpackDir(version)

	dir := platform.CheckForExtension(platform.Name(), path.Join(tempDir, file))

	if err := os.MkdirAll(tempDir, os.ModePerm); err != nil {
		panic(fmt.Errorf("mkdirall failed: %w", err))
	}

	if _, err := os.Stat(dir); err == nil {
		logger.Debug.Printf("query engine exists, not unpacking. %s", time.Since(start))
		return
	}

	if err := os.WriteFile(dir, data, os.ModePerm); err != nil {
		panic(fmt.Errorf("unpack write file: %w", err))
	}
	logger.Debug.Printf("unpacked at %s in %s", dir, time.Since(start))
}
