package util

func AssertEquals(expected, actual interface{}) {
	if expected != actual {
		panic(actual)
	}
}
