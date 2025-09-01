package main

// Пишите тесты в этом файле
import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_GenerateRandomElements_WhenOk(t *testing.T) {
	size := 123
	result, err := generateRandomElements(size)
	require.NoError(t, err)
	assert.Equal(t, size, len(result))
}
func Test_GenerateRandomElements_WhenNotOk(t *testing.T) {
	size := 0
	_, err := generateRandomElements(size)
	require.Error(t, err)

	size = -15
	_, err = generateRandomElements(size)
	require.Error(t, err)
}
func Test_Maximum_WhenOk(t *testing.T) {
	slice := []int{12, 234, -2134, 234, 832}
	sliceCopy := slice
	result := maxChunks(slice)
	assert.Equal(t, 832, result)
	assert.Equal(t, sliceCopy, slice)

	slice = []int{-234, -2134, -12, -234, -1234, -832}
	sliceCopy = slice
	result = maxChunks(slice)
	assert.Equal(t, -12, result)
	assert.Equal(t, sliceCopy, slice)

	slice = []int{5}
	sliceCopy = slice
	result = maxChunks(slice)
	assert.Equal(t, 5, result)
	assert.Equal(t, sliceCopy, slice)

	slice = []int{0, 0, 0}
	sliceCopy = slice
	result = maxChunks(slice)
	assert.Equal(t, 0, result)
	assert.Equal(t, sliceCopy, slice)
}
func Test_Maximum_WhenNotOk(t *testing.T) {
	slice := []int{}
	sliceCopy := slice
	result := maximum(slice)
	assert.Equal(t, 0, result)
	assert.Equal(t, sliceCopy, slice)
}
func Test_MaxChunks_WhenOk(t *testing.T) {
	slice := []int{12, 234, -2134, 234, 832}
	sliceCopy := slice
	result := maxChunks(slice)
	assert.Equal(t, 832, result)
	assert.Equal(t, sliceCopy, slice)

	slice = []int{-234, -2134, -12, -234, -1234, -832}
	sliceCopy = slice
	result = maxChunks(slice)
	assert.Equal(t, -12, result)
	assert.Equal(t, sliceCopy, slice)

	slice = []int{5}
	sliceCopy = slice
	result = maxChunks(slice)
	assert.Equal(t, 5, result)
	assert.Equal(t, sliceCopy, slice)

	slice = []int{0, 0, 0}
	sliceCopy = slice
	result = maxChunks(slice)
	assert.Equal(t, 0, result)
	assert.Equal(t, sliceCopy, slice)
}
func Test_MaxChunks_WhenNotOk(t *testing.T) {
	slice := []int{}
	sliceCopy := slice
	result := maximum(slice)
	assert.Equal(t, 0, result)
	assert.Equal(t, sliceCopy, slice)
}
