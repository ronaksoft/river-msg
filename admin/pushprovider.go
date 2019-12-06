package msgAdmin

func IsValidType(t PushProviderType) bool {
	switch t {
	case PushProviderType_Firebase, PushProviderType_Apple:
		return true
	}
	return false
}
