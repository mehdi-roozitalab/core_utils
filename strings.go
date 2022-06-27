// Basic types and utilities to be used generally in any version of the application
package core_utils

// StringArrayContains check if a slice of strings contains specific string or not
func StringArrayContains(items []string, item string) bool {
	for _, s := range items {
		if s == item {
			return true
		}
	}
	return false
}

func CloneStringAnyMap(m map[string]interface{}) map[string]interface{} {
	result := map[string]interface{}{}
	for k, v := range m {
		result[k] = v
	}
	return result
}

func CloneStringStringMap(m map[string]string) map[string]string {
	result := map[string]string{}
	for k, v := range m {
		result[k] = v
	}
	return result
}
