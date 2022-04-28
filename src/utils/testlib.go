package utils

func AssertEquals(expected, actual interface{}) {
	if expected != actual {
		panic(actual)
	}
}
