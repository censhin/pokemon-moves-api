package main

func MovesService() (Moves) {
	return MovesDao()
}

func MoveService(name string) (Move) {
	return MoveDao(name)
}
