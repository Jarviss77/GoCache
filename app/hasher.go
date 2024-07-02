package main

// polynomial rolling hash function https://cp-algorithms.com/string/string-hashing.html
func hasher(key string) int{
	const p = 31
	const m = 1e9 + 9
	hashedKey := 0
	pow_p := 1
	for i := 0; i < len(key); i++ {
		hashedKey = (hashedKey + (int(key[i])) * pow_p) % m
		pow_p = (pow_p * p) % m
	}
	return hashedKey
}