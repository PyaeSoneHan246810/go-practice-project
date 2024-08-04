package iomanager

type IOManager interface {
	ReadLines() ([]string, error)
	WriteData(data any) error
}
