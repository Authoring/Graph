package node

import (
	"github.com/Authoring/Graph/engine"
	"github.com/Authoring/Graph/logger"
	"github.com/google/uuid"
)

// Node defines the basic node structure
type Node struct {
	ID     string
	Name   string
	Engine *engine.Engine
}

// Actions defines the node actions
type Actions interface {
	Save() error
}

// NewNode create a new node
func NewNode(e *engine.Engine) *Node {
	return &Node{
		ID:     generateNodeID(),
		Engine: e,
	}
}

// Save saves the node to the database
func (n *Node) Save() error {
	logger.L.Infof("Saving %s", n.Name)
	return nil
}

func generateNodeID() string {
	return uuid.New().String()
}
