package saver

//Экспортируемые функции приватных полей для тестов

//IsBufferEmpty возвращает true, если буфер saver пустой
func (s *saver) IsBufferEmpty() bool {
	s.Lock()
	defer s.Unlock()
	return len(s.buffer) == 0
}

//IsInit возвращает true, если запущен saver
func (s *saver) IsInit() bool {
	s.configGuard.Lock()
	defer s.configGuard.Unlock()
	return s.config.isInit
}

//IsClosed возвращает true, если закрыт канал closeCh
func (s *saver) IsClosed() bool {
	s.configGuard.Lock()
	defer s.configGuard.Unlock()
	return s.config.isClosed
}
