package store

type StoreRouter interface {
	Timer() TimerStore
}
