package main

type StringIntMap struct {
	data map[string]int
}

// 1. Добавление элемента: Метод Add(key string, value int), который добавляет новую пару "ключ-значение" в карту.
func (s *StringIntMap) Add(key string, value int) {
	s.data[key] = value
}

// 2. Удаление элемента: Метод Remove(key string), который удаляет элемент по ключу из карты.
func (s *StringIntMap) Remove(key string) {
	if s.Exists(key) {
		delete(s.data, key)
	}
}

// 3. Копирование карты: Метод Copy() map[string]int, который возвращает новую карту, содержащую все элементы текущей карты.
func (s *StringIntMap) Copy() map[string]int {
	data := make(map[string]int)
	for k, v := range s.data {
		data[k] = v
	}
	return data
}

// 4. Проверка наличия ключа: Метод Exists(key string) bool, который проверяет, существует ли ключ в карте.
func (s *StringIntMap) Exists(key string) bool {
	if _, ok := s.data[key]; ok {
		return true
	}
	return false
}

// 5. Получение значения: Метод Get(key string) (int, bool), который возвращает значение по ключу и булевый флаг, указывающий на успешность операции.
func (s *StringIntMap) Get(key string) (int, bool) {
	if s.Exists(key) {
		return s.data[key], true
	}
	return 0, false
}
