package fixedfloat

import (
	"bytes"
	"strconv"
)

type Bool bool

func (bit *Bool) UnmarshalJSON(b []byte) error {
	txt := string(bytes.Trim(b, `"`))
	*bit = Bool(txt == "1" || txt == "true")
	return nil
}

func (bit *Bool) Toint() int {
	if *bit {
		return 1
	}
	return 0
}

type Integer int

func (i *Integer) UnmarshalJSON(b []byte) error {
	txt := string(bytes.Trim(b, `"`))
	integer, err := strconv.ParseInt(txt, 10, 64)
	if err != nil {
		return err
	}
	*i = Integer(integer)
	return nil
}
