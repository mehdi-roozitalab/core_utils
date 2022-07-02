package core_utils

type Bool3 int8

const (
	B3True  Bool3 = 1
	B3False Bool3 = 0
	B3Null  Bool3 = -1
)

func (b Bool3) HaveValue() bool     { return b != B3Null }
func (b Bool3) IsTrue() bool        { return b == B3True }
func (b Bool3) IsTrueOrNull() bool  { return b != B3False }
func (b Bool3) IsFalse() bool       { return b == B3False }
func (b Bool3) IsFalseOrNull() bool { return b != B3True }
