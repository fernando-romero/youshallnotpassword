package main

type SliceOfStrings []string

func (s SliceOfStrings) Len() int {
	return len(s)
}
func (s SliceOfStrings) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s SliceOfStrings) Less(i, j int) bool {
	if len(s[i]) != len(s[j]) {
		return len(s[i]) > len(s[j])
	}
	return s[i] > s[j]
}
