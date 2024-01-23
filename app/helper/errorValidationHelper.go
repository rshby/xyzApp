package helper

func GetErrorByTag(tag string) string {
	switch tag {
	case "min":
		return "harap isi 16 digit"
	default:
		return ""
	}
}
