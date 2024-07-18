package calendar

type Client struct {
}

// Cal
func (c Client) ReadEvents() ([]Event, error) {
	return []Event{}, nil
}

type Event struct {
}
