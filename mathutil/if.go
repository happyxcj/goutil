package mathutil

// IfInt 当cond是真的时候返回okVal，否则返回elseVal
func IfInt(cond bool, okVal, elseVal int) int {
	if cond {
		return okVal
	}
	return elseVal
}

// IfInt64 当cond是真的时候返回okVal，否则返回elseVal
func IfInt64(cond bool, okVal, elseVal int64) int64 {
	if cond {
		return okVal
	}
	return elseVal
}

// IfInt32 当cond是真的时候返回okVal，否则返回elseVal
func IfInt32(cond bool, okVal, elseVal int32) int32 {
	if cond {
		return okVal
	}
	return elseVal
}

// IfInt16 当cond是真的时候返回okVal，否则返回elseVal
func IfInt16(cond bool, okVal, elseVal int16) int16 {
	if cond {
		return okVal
	}
	return elseVal
}

// IfInt8 当cond是真的时候返回okVal，否则返回elseVal
func IfInt8(cond bool, okVal, elseVal int8) int8 {
	if cond {
		return okVal
	}
	return elseVal
}

// IfUint 当cond是真的时候返回okVal，否则返回elseVal
func IfUint(cond bool, okVal, elseVal uint) uint {
	if cond {
		return okVal
	}
	return elseVal
}

// IfUint64 当cond是真的时候返回okVal，否则返回elseVal
func IfUint64(cond bool, okVal, elseVal uint64) uint64 {
	if cond {
		return okVal
	}
	return elseVal
}

// IfUint32 当cond是真的时候返回okVal，否则返回elseVal
func IfUint32(cond bool, okVal, elseVal uint32) uint32 {
	if cond {
		return okVal
	}
	return elseVal
}

// IfUint16 当cond是真的时候返回okVal，否则返回elseVal
func IfUint16(cond bool, okVal, elseVal uint16) uint16 {
	if cond {
		return okVal
	}
	return elseVal
}

// IfUint8 当cond是真的时候返回okVal，否则返回elseVal
func IfUint8(cond bool, okVal, elseVal uint8) uint8 {
	if cond {
		return okVal
	}
	return elseVal
}

// IfFloat64 当cond是真的时候返回okVal，否则返回elseVal
func IfFloat64(cond bool, okVal, elseVal float64) float64 {
	if cond {
		return okVal
	}
	return elseVal
}

// IfFloat32 当cond是真的时候返回okVal，否则返回elseVal
func IfFloat32(cond bool, okVal, elseVal float32) float32 {
	if cond {
		return okVal
	}
	return elseVal
}

// IfStr 当cond是真的时候返回okVal，否则返回elseVal
func IfStr(cond bool, okVal, elseVal string) string {
	if cond {
		return okVal
	}
	return elseVal
}

// IfObj 当cond是真的时候返回okVal，否则返回elseVal
func IfObj(cond bool, okVal, elseVal interface{}) interface{} {
	if cond {
		return okVal
	}
	return elseVal
}
