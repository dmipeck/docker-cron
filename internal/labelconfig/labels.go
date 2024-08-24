package labelconfig

// Namespace is prefix added to all confguration labels
type Namespace string

// Label takes a configuration property key "key" and returns a namespaced label string for that
// property key
func (ns Namespace) LabelKey(key string) string {
	return string(ns) + "." + key
}

// Label takes a configuration property "property" and returns a namespaced label string for that
// property
func (ns Namespace) LabelKeyValue(key string, value string) string {
	return string(ns) + "." + key + "=" + value
}
