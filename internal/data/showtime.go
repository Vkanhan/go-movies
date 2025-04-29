package data

import (
	"fmt"
	"strconv"
)

//Customize the Showtime to output in string format
type Showtime int32

func (s Showtime) MarshalJSON() ([]byte, error) {
	json := fmt.Sprintf("%d mins", s)
	quotedJSON := strconv.Quote(json)
	return []byte(quotedJSON), nil
}


