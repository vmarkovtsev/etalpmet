package etalpmet

import (
	"bytes"
	"sort"

	"gopkg.in/vmarkovtsev/go-lcss.v1"
)

// ReverseTemplate infers the ordered list of common blocks in the given list of strings.
// The first element of that list is `nil` if there are leading bytes in at least one string;
// likewise, the last element is `nil` if there are trailing bytes.
//
// Complexity:
// * time: `sum(n_i*log(n_i))*T` where `n_i` is the length of each string and `T` is the number of found blocks.
// * space: `sum(n_i)`.
func ReverseTemplate(strs ...[]byte) [][]byte {
	return ReverseTemplateWithParameters(2, true, strs...)
}

// ReverseTemplateWithParameters infers the ordered list of common blocks in the given list of strings.
// The first element of that list is `nil` if there are leading bytes in at least one string;
// likewise, the last element is `nil` if there are trailing bytes.
// `minBlockLength` sets the minimum block length to be included into the template.
// `trimSpace` removes whitespace around the blocks. `minBlockLength` is checked after trimming.
//
// Complexity:
// * time: `sum(n_i*log(n_i))*T` where `n_i` is the length of each string and `T` is the number of found blocks.
// * space: `sum(n_i)`.
func ReverseTemplateWithParameters(minBlockLength int, trimSpace bool, strs ...[]byte) [][]byte {
	var template []anchor
	// we recursively find the longest common substring of the strings and split them into two parts
	parts := []chunk{{position{0, 1}, strs}}
	for len(parts) > 0 {
		part := parts[len(parts)-1]
		parts = parts[:len(parts)-1]
		lcs := lcss.LongestCommonSubstring(part.strs...)
		efflcs := lcs
		if trimSpace {
			efflcs = bytes.TrimSpace(efflcs)
		}
		if len(efflcs) < minBlockLength {
			continue
		}
		template = append(template, anchor{position{part.start*2 + 1, part.fraction * 2}, efflcs})
		// we split each string into left and right part relative to `lcs`
		newStrsLeft := make([][]byte, len(strs))
		newStrsRight := make([][]byte, len(strs))
		var hasLeft, hasRight bool
		for i, str := range part.strs {
			lr := bytes.SplitN(str, lcs, 2)
			left, right := lr[0], lr[1]
			if len(left) > 0 {
				hasLeft = true
				newStrsLeft[i] = left
			}
			if len(right) > 0 {
				hasRight = true
				newStrsRight[i] = right
			}
		}
		if hasLeft {
			parts = append(parts, chunk{position{part.start * 2, part.fraction * 2}, newStrsLeft})
		}
		if hasRight {
			parts = append(parts, chunk{position{part.start*2 + 1, part.fraction * 2}, newStrsRight})
		}
	}
	// we went depth-first so we need to sort
	sort.Slice(template, func(i, j int) bool {
		first, second := template[i], template[j]
		firstStart, secondStart := first.start, second.start
		if second.fraction > first.fraction {
			firstStart *= second.fraction / first.fraction
		} else {
			secondStart *= first.fraction / second.fraction
		}
		return firstStart < secondStart
	})
	// check the edges - do we need to insert nil-s?
	var hasLeft, hasRight int
	for _, str := range strs {
		if hasLeft == 0 && !bytes.HasPrefix(str, template[0].block) {
			hasLeft = 1
		}
		if hasRight == 0 && !bytes.HasSuffix(str, template[len(template)-1].block) {
			hasRight = 1
		}
	}
	result := make([][]byte, len(template)+hasLeft+hasRight)
	for i, t := range template {
		result[i+hasLeft] = t.block
	}
	if hasLeft == 1 {
		result[0] = nil
	}
	if hasRight == 1 {
		result[len(result)-1] = nil
	}
	return result
}

type position struct {
	start    int
	fraction int
}

type chunk struct {
	position
	strs [][]byte
}

type anchor struct {
	position
	block []byte
}
