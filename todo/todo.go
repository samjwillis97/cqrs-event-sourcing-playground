package todo

type Todo struct {
	title       string
	description string
	status      Status

	id      int
	version uint64
	//events []interface{}
}

type Status int

const (
	todo Status = iota
	inProgress
	done
)

func Create(id int, title string, description string) *Todo {
	item := Todo{}

	// TODO: Fire off event

	return &item
}

func NewFromEvents(events []interface{}) *Todo {
	state := &Todo{}
	for _, event := range events {
		state.On(event)
		state.version++
	}
	return state
}

func (state *Todo) On(event interface{}) {
	switch e := event.(type) {
	case TodoCreated:
		state.title = e.Title
		state.description = e.Descriptions
		state.status = todo
	case TodoTitleModified:
		state.title = e.Title
	case TodoDescriptionModified:
		state.description = e.Description
	case TodoInProgress:
		state.status = inProgress
	case TodoDone:
		state.status = done
	}
}
