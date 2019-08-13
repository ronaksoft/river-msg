package msg

const (
	LangCodeEn = "en"
	LangCodeFa = "fa"
)

// LangCodeStr converts lang code from proto enum
// to its corresponding string value
func LangCodeStr(code LangCode) string {
	switch code {
	case LangCode_LangCodeEn:
		return LangCodeEn
	case LangCode_LangCodeFa:
		return LangCodeFa
	default:
		return LangCodeEn
	}
}
