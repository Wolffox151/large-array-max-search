package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Пишите тесты в этом файле
func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name    string
		size    int
		wantLen int
	}{
		{name: "Отрицательный размер", size: -1, wantLen: 0},
		{name: "Нулевой размер", size: 0, wantLen: 0},
		{name: "Один элемент", size: 1, wantLen: 1},
		{name: "Несколько элементов", size: 10, wantLen: 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateRandomElements(tt.size)
			require.Len(t, got, tt.wantLen)

			for i, value := range got {
				assert.GreaterOrEqualf(t, value, 1, "элемент с индексом %d", i)
				assert.LessOrEqualf(t, value, 1_000_000, "элемент с индексом %d", i)
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{name: "Слайс с nil", data: nil, want: 0},
		{name: "Пустой слайс", data: []int{}, want: 0},
		{name: "Один элемент", data: []int{7}, want: 7},
		{name: "Максимальный элемент в начале", data: []int{9, 4, 1}, want: 9},
		{name: "Максимальный элемент в середине", data: []int{4, 9, 1}, want: 9},
		{name: "Макисмальный элемент в конце", data: []int{4, 1, 9}, want: 9},
		{name: "Повторяющийся максимальный элемент", data: []int{4, 9, 1, 9}, want: 9},
		{name: "Все максимальные", data: []int{9, 9, 9, 9}, want: 9},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, maximum(tt.data))
		})
	}
}
