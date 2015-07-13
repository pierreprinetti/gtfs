// gtfs - Google Transit Feed Specification Reader and Writer for Go
package gtfs

import "log"

type Line []byte
type CsvFile []Line

func checkError(err error) {
	if err != nil {
		log.Panicln(err)
	}
}
