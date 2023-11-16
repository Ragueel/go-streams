package messagestream

type Topic int

const (
	Conversion Topic = iota
	ConversionFailed
	Resize
	ResizeFailed
)

func (topic Topic) String() string {
	return []string{"conversion.started", "conversion.failed", "resize.started", "resize.failed"}[topic]
}
