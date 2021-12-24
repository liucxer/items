package enums

//go:generate tools gen enum ResType

// ResType 资源类型
type ResType uint8

const (
	RES_TYPE_UNKNOWN ResType = iota
	RES_TYPE__DOC            // .doc
	RES_TYPE__PDF            // .pdf
	RES_TYPE__IMAGE          // .jpeg,.jpg,.png...
	RES_TYPE__VIDEO          // .mp4
	RES_TYPE__APK            // .apk
)
