package event

type EventConfig struct {
	Events map[EventType]*Event
}

type Event struct {
	EventType  EventType
	JsFunction string
}

type EventType string

const (
	MouseClick EventType = "click"
)

func NewEventConfig() *EventConfig {
	return &EventConfig{
		Events: make(map[EventType]*Event),
	}
}

func (ec *EventConfig) On(eventType EventType, Js string) {
	ec.Events[eventType] = &Event{
		EventType:  eventType,
		JsFunction: Js,
	}
}
