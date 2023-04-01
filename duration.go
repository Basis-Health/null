package null

import "time"

type Duration Int

func NewDuration(d time.Duration, valid bool) Duration {
	return Duration(NewInt(d.Nanoseconds(), valid))
}

func DurationFrom(d time.Duration) Duration {
	return Duration(IntFrom(d.Nanoseconds()))
}

func DurationFromPtr(d *time.Duration) Duration {
	if d == nil {
		return Duration(NewInt(0, false))
	}
	return Duration(DurationFrom(*d))
}

func (d Duration) ValueOrZero() int64 {
	return Int(d).ValueOrZero()
}

func (d Duration) MarshalJSON() ([]byte, error) {
	return Int(d).MarshalJSON()
}

func (d *Duration) UnmarshalJSON(data []byte) error {
	return (*Int)(d).UnmarshalJSON(data)
}

func (d Duration) MarshalText() ([]byte, error) {
	return Int(d).MarshalText()
}

func (d *Duration) UnmarshalText(text []byte) error {
	return (*Int)(d).UnmarshalText(text)
}

func (d *Duration) SetValid(n time.Duration) {
	(*Int)(d).SetValid(n.Nanoseconds())
}

func (d Duration) Ptr() *int64 {
	return Int(d).Ptr()
}

func (d Duration) IsZero() bool {
	return Int(d).IsZero()
}

func (d Duration) Equal(other Duration) bool {
	return Int(d).Equal(Int(other))
}
