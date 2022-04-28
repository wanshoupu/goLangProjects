package utils

import (
	"strconv"
	"testing"
)

func TestBinaryTreeLayer(t *testing.T) {
	test := 14
	result := BinaryTreeLayer(test)
	if result != 4 {
		panic(strconv.Itoa(result) + " not expected")
	}
}

func TestBinaryTreeLayerBoundary(t *testing.T) {
	test := 0
	result := BinaryTreeLayer(test)
	if result != 0 {
		panic(strconv.Itoa(result) + " not expected")
	}
}

func TestBinaryTreeLayerOne(t *testing.T) {
	test := 1
	result := BinaryTreeLayer(test)
	if result != 1 {
		panic(strconv.Itoa(result) + " not expected")
	}
}

func TestParentLabel(t *testing.T) {
	test := 14
	result := ParentLabel(test)
	if result != 4 {
		panic(strconv.Itoa(result) + " not expected")
	}
}

func TestParentLabel4(t *testing.T) {
	test := 4
	result := ParentLabel(test)
	if result != 3 {
		panic(strconv.Itoa(result) + " not expected")
	}
}
