package enums

import (
	bytes "bytes"
	database_sql_driver "database/sql/driver"
	errors "errors"

	github_com_go_courier_enumeration "github.com/go-courier/enumeration"
)

var InvalidResType = errors.New("invalid ResType type")

func ParseResTypeFromLabelString(s string) (ResType, error) {
	switch s {
	case "":
		return RES_TYPE_UNKNOWN, nil
	case ".doc":
		return RES_TYPE__DOC, nil
	case ".pdf":
		return RES_TYPE__PDF, nil
	case ".jpeg,.jpg,.png...":
		return RES_TYPE__IMAGE, nil
	case ".mp4":
		return RES_TYPE__VIDEO, nil
	case ".apk":
		return RES_TYPE__APK, nil
	}
	return RES_TYPE_UNKNOWN, InvalidResType
}

func (v ResType) String() string {
	switch v {
	case RES_TYPE_UNKNOWN:
		return ""
	case RES_TYPE__DOC:
		return "DOC"
	case RES_TYPE__PDF:
		return "PDF"
	case RES_TYPE__IMAGE:
		return "IMAGE"
	case RES_TYPE__VIDEO:
		return "VIDEO"
	case RES_TYPE__APK:
		return "APK"
	}
	return "UNKNOWN"
}

func ParseResTypeFromString(s string) (ResType, error) {
	switch s {
	case "":
		return RES_TYPE_UNKNOWN, nil
	case "DOC":
		return RES_TYPE__DOC, nil
	case "PDF":
		return RES_TYPE__PDF, nil
	case "IMAGE":
		return RES_TYPE__IMAGE, nil
	case "VIDEO":
		return RES_TYPE__VIDEO, nil
	case "APK":
		return RES_TYPE__APK, nil
	}
	return RES_TYPE_UNKNOWN, InvalidResType
}

func (v ResType) Label() string {
	switch v {
	case RES_TYPE_UNKNOWN:
		return ""
	case RES_TYPE__DOC:
		return ".doc"
	case RES_TYPE__PDF:
		return ".pdf"
	case RES_TYPE__IMAGE:
		return ".jpeg,.jpg,.png..."
	case RES_TYPE__VIDEO:
		return ".mp4"
	case RES_TYPE__APK:
		return ".apk"
	}
	return "UNKNOWN"
}

func (v ResType) Int() int {
	return int(v)
}

func (ResType) TypeName() string {
	return "github.com/saitofun/items/pkg/consts.ResType"
}

func (ResType) ConstValues() []github_com_go_courier_enumeration.IntStringerEnum {
	return []github_com_go_courier_enumeration.IntStringerEnum{RES_TYPE__DOC, RES_TYPE__PDF, RES_TYPE__IMAGE, RES_TYPE__VIDEO, RES_TYPE__APK}
}

func (v ResType) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidResType
	}
	return []byte(str), nil
}

func (v *ResType) UnmarshalText(data []byte) (err error) {
	*v, err = ParseResTypeFromString(string(bytes.ToUpper(data)))
	return
}

func (v ResType) Value() (database_sql_driver.Value, error) {
	offset := 0
	if o, ok := (interface{})(v).(github_com_go_courier_enumeration.DriverValueOffset); ok {
		offset = o.Offset()
	}
	return int64(v) + int64(offset), nil
}

func (v *ResType) Scan(src interface{}) error {
	offset := 0
	if o, ok := (interface{})(v).(github_com_go_courier_enumeration.DriverValueOffset); ok {
		offset = o.Offset()
	}

	i, err := github_com_go_courier_enumeration.ScanIntEnumStringer(src, offset)
	if err != nil {
		return err
	}
	*v = ResType(i)
	return nil
}
