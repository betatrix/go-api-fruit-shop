package enums

type Role string

const (
	ADMIN Role = "admin"
	USER  Role = "user"
)

func ParseRole(inputString string) Role {
	switch inputString {
	case "admin":
		return ADMIN
	case "user":
		return USER
	default:
		return ""
	}
}

func ParseString(inputRole Role) string {
	switch inputRole {
	case ADMIN:
		return "admin"
	case USER:
		return "user"
	default:
		return ""
	}
}
