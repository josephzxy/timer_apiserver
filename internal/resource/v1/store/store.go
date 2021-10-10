// Package store provides low-level interfaces for data storage operations.
package store

// Router is an interface that provides routes to stores
// dedicated to a specific scope of RESTful resource.
type Router interface {
	Timer() TimerStore
}
