package etalpmet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseTemplate(t *testing.T) {
	template := ReverseTemplate(
		[]byte("<div>test<br>help</div>"),
		[]byte("<div>hello<br>order</div>"),
		[]byte("<div>trolo<br>lolo</div>"))
	assert.Equal(t, [][]byte{[]byte("<div>"), []byte("<br>"), []byte("</div>")}, template)
	template = ReverseTemplate(
		[]byte("a<div>test<br> help</div>"),
		[]byte("<div>hello<br> order</div>"),
		[]byte("<div>trolo<br> lolo</div>b"))
	assert.Equal(t, [][]byte{nil, []byte("<div>"), []byte("<br>"), []byte("</div>"), nil}, template)
	template = ReverseTemplate(
		[]byte("a bc de"),
		[]byte("a xy de gh"),
		[]byte("!! a rty de ty"),
		[]byte("?a xxx defq"))
	assert.Equal(t, [][]byte{nil, []byte("de"), nil}, template)
	template = ReverseTemplate(
		[]byte("IMG_20180930_171704.jpg"),
		[]byte("IMG_20181001_150308.jpg"),
		[]byte("IMG_20181001_190338.jpg"),
		[]byte("IMG_20181021_122346.jpg"))
	assert.Equal(t, [][]byte{[]byte("IMG_2018"), []byte("_1"), []byte(".jpg")}, template)
}

func TestReverseTemplateWithParameters(t *testing.T) {
	template := ReverseTemplateWithParameters(
		5,     // min block length
		false, // trim
		[]byte("<b> spam and eggs </b>"),
		[]byte("<b> ham and spam </b>"),
		[]byte("<b> white and black </b>"))
	assert.Equal(t, [][]byte{nil, []byte(" and "), []byte(" </b>")}, template)
}
