package file

import (
	"crypto/rand"
	"io"
	"os"

	"github.com/rs/zerolog/log"
)

const BUFF_SIZE = 1024 * 2 // some reasonable IO buffer size for reading from /dev/random and writing to file
const SHRED_N_TIMES = 3

// Open the file at `path`, and mangle its contents SHRED_N_TIMES before deleting it
func Shred(path string) error {

	for i := 0; i < SHRED_N_TIMES; i++ {
		if err := Mangle(path); err != nil {
			return err
		}
	}

	if err := os.Remove(path); err != nil {
		return err
	}

	return nil
}

// open the file at the path and randomize its data using an IO buffer
func Mangle(path string) error {
	stats, err := os.Stat(path)
	if err != nil {
		log.Error().Err(err).Str("filename", path).Msg("Unable to get stats on the file. Does the file exist?")

		return err
	}

	fileSize := stats.Size()
	remainingBytes := fileSize

	file, err := os.OpenFile(path, os.O_WRONLY, stats.Mode())

	if err != nil {
		log.Error().Err(err).Str("filename", path).Msg("Unable to open file. Are the permissions valid?")

		return err
	}

	defer file.Close()

	for remainingBytes > BUFF_SIZE {
		// Read random bytes from the crypto/rand source
		if err := writeRandom(remainingBytes, fileSize-remainingBytes, file); err != nil {
			return nil
		}
		remainingBytes -= BUFF_SIZE
	}

	if remainingBytes > 0 {
		log.Info().Int("remaining Bytes", int(remainingBytes)).Send()
		if err := writeRandom(remainingBytes, fileSize-remainingBytes, file); err != nil {
			return err
		}
	}

	if err := file.Sync(); err != nil {
		log.Error().Err(err).Msg("Unable to flush the file to disk.")

		return err
	}

	return nil
}

// Write `writeSize` bytes to `file`
func writeRandom(writeSize int64, writeOffset int64, file *os.File) error {
	randomBytes := make([]byte, writeSize)
	if _, err := io.ReadFull(rand.Reader, randomBytes); err != nil {
		log.Error().Err(err).Msg("Unable to get random data from rand.Reader.")

		return err
	}

	if _, err := file.WriteAt(randomBytes, writeOffset); err != nil {
		log.Error().Err(err).Msg("Unable to write random data to the file.")

		return err
	}

	return nil
}
