package chat

import (
	"github.com/asdine/storm"
	"github.com/google/uuid"
)

type PPLXChatStore interface {
	CreateThread(t *Thread) error
	GetThread(ID uuid.UUID) (*Thread, error)
	ListThreads() ([]*Thread, error)
	UpdateThread(ID uuid.UUID, t *Thread) error
	DeleteThread(ID uuid.UUID) error

	CreateTopic(t *ChatTopic) error
	GetTopic(ID uuid.UUID) (*ChatTopic, error)
	ListTopics() ([]*ChatTopic, error)
	DeleteTopic(ID uuid.UUID) error
	UpdateTopic(ID uuid.UUID, t *ChatTopic) error

	SaveChatResponse(cr *ChatResponse) error
	GetChatResponse(ID string) (*ChatResponse, error)
	ListChatResponses() ([]*ChatResponse, error)
	ClearEntries() error
	ClearEntriesByTopic(name string) error
}

type StromRepo struct {
	db *storm.DB
}

func NewStromRepo(db *storm.DB) *StromRepo {
	return &StromRepo{db}
}

func (sr *StromRepo) CreateThread(t *Thread) error {
	return sr.db.Save(t)
}

func (sr *StromRepo) GetThread(ID uuid.UUID) (*Thread, error) {
	var t Thread
	err := sr.db.One("ID", ID, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (sr *StromRepo) ListThreads() ([]*Thread, error) {
	var threads []*Thread
	err := sr.db.All(&threads)
	if err != nil {
		return nil, err
	}
	return threads, nil
}

func (sr *StromRepo) UpdateThread(ID uuid.UUID, t *Thread) error {
	return sr.db.Update(t)
}

func (sr *StromRepo) DeleteThread(ID uuid.UUID) error {
	var t Thread
	err := sr.db.One("ID", ID, &t)
	if err != nil {
		return err
	}
	return sr.db.DeleteStruct(&t)

}

func (sr *StromRepo) CreateTopic(t *ChatTopic) error {
	return sr.db.Save(t)
}

func (sr *StromRepo) GetTopic(ID uuid.UUID) (*ChatTopic, error) {
	var t ChatTopic
	err := sr.db.One("ID", ID, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (sr *StromRepo) ListTopics() ([]*ChatTopic, error) {
	var topics []*ChatTopic
	err := sr.db.All(&topics)
	if err != nil {
		return nil, err
	}
	return topics, nil
}

func (sr *StromRepo) DeleteTopic(ID uuid.UUID) error {
	var t ChatTopic
	err := sr.db.One("ID", ID, &t)
	if err != nil {
		return err
	}
	return sr.db.DeleteStruct(&t)
}

func (sr *StromRepo) UpdateTopic(ID uuid.UUID, t *ChatTopic) error {
	return sr.db.Update(t)
}

func (sr *StromRepo) SaveChatResponse(cr *ChatResponse) error {
	return sr.db.Save(cr)
}

func (sr *StromRepo) GetChatResponse(ID string) (*ChatResponse, error) {
	var cr ChatResponse
	err := sr.db.One("ID", ID, &cr)
	if err != nil {
		return nil, err
	}
	return &cr, nil
}

func (sr *StromRepo) ListChatResponses() ([]*ChatResponse, error) {
	var crs []*ChatResponse
	err := sr.db.All(&crs)
	if err != nil {
		return nil, err
	}
	return crs, nil
}

func (sr *StromRepo) ClearEntries() error {
	// delete messages
	err := sr.db.Drop(&ChatResponse{})
	if err != nil {
		return err
	}
	// delete threads
	err = sr.db.Drop(&Thread{})
	if err != nil {
		return err
	}
	// delete topics
	err = sr.db.Drop(&ChatTopic{})
	if err != nil {
		return err
	}
	return nil
}

func (sr *StromRepo) ClearEntriesByTopic(name string) error {
	var topic ChatTopic
	err := sr.db.One("Name", name, &topic)
	if err != nil {
		return err
	}
	// delete messages
	err = sr.db.Drop(&ChatResponse{})
	if err != nil {
		return err
	}
	// delete threads
	err = sr.db.Drop(&Thread{})
	if err != nil {
		return err
	}
	// delete topic
	err = sr.db.DeleteStruct(&topic)
	if err != nil {
		return err
	}
	return nil
}
