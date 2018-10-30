//go:generate go get github.com/alvaroloes/enumer
//go:generate enumer -type=Type

package change

// Type is an enumeration for data change type
type Type uint

const (
	// Updated is type of change when data was updated
	Updated Type = iota
	// Deleted is type of change when data was deleted
	Deleted Type = iota
)
