package fixedfloat

import "bytes"

type Bool bool

func (bit *Bool) UnmarshalJSON(b []byte) error {
	txt := string(bytes.Trim(b, `"`))
	*bit = Bool(txt == "1" || txt == "true")
	return nil
}
