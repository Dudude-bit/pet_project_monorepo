package utils

func GetAddressOfValue[vType any](v vType) *vType {
	return &v
}
