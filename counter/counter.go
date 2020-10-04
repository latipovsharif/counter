package counter

import (
	"sync"
)

// NewCounter возвращает экземпляр класса
// Counter с установленным максимальным значением
// для счетчика, если значение установлено в 0
// то оно будет равно максимальному значению для типа uint
func NewCounter(maxPossibleValue uint) *Counter {
	mv := ^uint(0)

	if maxPossibleValue != 0 {
		mv = maxPossibleValue
	}

	return &Counter{
		maxPossibleValue: mv,
	}
}

// Counter представляет класс счетчика.
type Counter struct {
	value            uint
	maxPossibleValue uint
	mutex            sync.RWMutex
}

// Increment при вызове метода свойство value
// увеличивается на единицу до того как значение value
// не достигает значения maxPossibleValue.
// Если значение value равняется maxPossibleValue
// то значение value обнуляется
func (c *Counter) Increment() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	// Если не использовать данную проверку
	// то есть вероятность вызова метода
	// при неустановленном значении maxPossibleValue.
	// Есть возможность сделать саму структуру Counter{}
	// приватным, но в этом случае использование данной
	// реализации может быть неудобным.
	if c.maxPossibleValue == 0 {
		c.maxPossibleValue = ^uint(0)
	}

	if c.value == c.maxPossibleValue {
		c.reset()

		return
	}

	c.value++
}

// Value возвращает текущее значение счетчика.
func (c *Counter) Value() uint {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return c.value
}

// SetMaximumValue устанавливает максимально
// возможное значение для счетчика, если
// установленное число меньше текущего значения value
// то значение value устанавливается в 0.
func (c *Counter) SetMaximumValue(v uint) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if v < c.value {
		c.reset()
	}

	c.maxPossibleValue = v
}

// reset сбрасывает значение счетчика.
func (c *Counter) reset() {
	c.value = 0
}
