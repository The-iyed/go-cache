package store

import (
	"regexp"
	"strings"
)

func (store *KeyValueStore) Keys(pattern string) []string {
	store.mu.RLock()
	defer store.mu.RUnlock()

	var keys []string

	useRegex := strings.Contains(pattern, "*")
	var re *regexp.Regexp
	if useRegex {
		re = regexp.MustCompile("^" + strings.ReplaceAll(regexp.QuoteMeta(pattern), "\\*", ".*") + "$")
	}

	for key := range store.data {
		if !useRegex && pattern == "*" || useRegex && re.MatchString(key) {
			keys = append(keys, key)
		}
	}
	return keys
}
