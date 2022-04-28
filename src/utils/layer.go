package utils

/**
The layer (starting at 1) the label (starting at 1) is located
if the labels [1 .. label] are arranged in an infinite binary tree
*/
func BinaryTreeLayer(label int) int {
	layer := 0
	for {
		if label < 1<<layer {
			return layer
		}
		layer++
	}
}

func ParentLabel(label int) int {
	if label <= 1 {
		return 0
	}
	layer := BinaryTreeLayer(label)
	if layer&1 == 0 {
		offset := 1 << (layer - 1)
		return (3*offset - label - 1) >> 1
	}
	layer--
	parent := label >> 1
	offset := 1 << (layer - 1)
	return 3*offset - parent - 1
}
