package list

import (
	"fmt"

	"github.com/pkg/errors"
)

type List struct {
	firstNode *node
	len       int64
}

func NewList() (list *List) {
	list = &List{
		firstNode: nil,
		len:       0,
	}
	return
}

func (l *List) AddData(input int64) {
	// add
	if l.len == 0 {
		l.firstNode = &node{
			nextPoint: nil,
			data:      input,
		}
		l.len++
		return
	}
	n := l.firstNode
	for {
		if n.nextPoint == nil {
			break
		}
		n = n.nextPoint
	}
	n.nextPoint = &node{
		nextPoint: nil,
		data:      input,
	}
	l.len++
}

func (l *List) PrintAll() {
	n := l.firstNode
	for {
		if n.nextPoint == nil {
			fmt.Println(n.data)
			return
		}
		fmt.Println(n.data)
		n = n.nextPoint
	}
}

func (l *List) Len() int64 {
	return l.len
}

// GetDataByIndex выдает значение по индексу, если есть. Иначе ошибка
func (l *List) GetDataByIndex(index int64) (data int64, err error) {
	if l.len <= index || index < 0 {
		err = errors.New("Error: Out of range")
		return
	}
	n := l.firstNode
	i := int64(0)
	for {
		if i == index {
			break
		}
		n = n.nextPoint
		i++
	}
	return n.data, err
}

// DeleteDataByIndex удаляет значение по индексу, если есть. Иначе ошибка
func (l *List) DeleteDataByIndex(index int64) (ok bool, err error) {
	if l.len <= index || index < 0 {
		err = errors.New("Error: Out of range")
		return
	}
	indexbefore := index - 1
	indexafter := index + 1
	n := l.firstNode
	var (
		nBefore  *node
		nCurrent *node
		nAfter   *node
	)
	i := int64(0)
	for {
		switch i {
		case indexbefore:
			nBefore = n
		case index:
			nCurrent = n
		case indexafter:
			nAfter = n
		}
		if nAfter != nil {
			break
		}
		n = n.nextPoint
		if n == nil {
			break
		}
		i++
	}
	nBefore.nextPoint = nAfter
	nCurrent.nextPoint = nil
	ok = true
	return
}

// Sort сортирует значения по возростанию, если increase = true, иначе по убыванию. Если сортировать нечего, то ошибка
func (l *List) Sort(increase bool) (ok bool, err error) {
	if l.len < 2 {
		err = errors.New("Список меньше 2")
		return
	}
	n := l.firstNode
	var (
		nCurrent *node
		nAfter   *node
	)
	if increase == true {
		for {
			if nAfter != nil {
				break
			}

		}
	}
	return
}
