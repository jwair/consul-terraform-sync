package config

// Bool returns a pointer to the given bool.
func Bool(b bool) *bool {
	return &b
}

// BoolVal returns the value of the boolean at the pointer, or false if the
// pointer is nil.
func BoolVal(b *bool) bool {
	if b == nil {
		return false
	}
	return *b
}

// BoolCopy returns a copy of the boolean pointer
func BoolCopy(b *bool) *bool {
	if b == nil {
		return nil
	}

	return Bool(*b)
}

// String returns a pointer to the given string.
func String(s string) *string {
	return &s
}

// StringVal returns the value of the string at the pointer, or "" if the
// pointer is nil.
func StringVal(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

// StringCopy returns a copy of the string pointer
func StringCopy(s *string) *string {
	if s == nil {
		return nil
	}

	return String(*s)
}

// StringPresent returns a boolean indicating if the pointer is nil, or if the
// pointer is pointing to the zero value.
func StringPresent(s *string) bool {
	if s == nil {
		return false
	}
	return *s != ""
}
