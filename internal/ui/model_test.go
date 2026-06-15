package ui

import (
	"fmt"
	"testing"

	"github.com/talonlikeaclaw/godash/internal/models"
)

func TestUpdateHandlesDataRefreshMsg(t *testing.T) {
	m := NewModelWithFakeData()

	msg := dataRefreshMsg{
		node:       models.Node{Name: "pve", CPU: 12.5, Status: "online"},
		vms:        []models.VM{{ID: 1, Name: "test-vm", Status: "running"}},
		containers: []models.Container{{ID: 101, Name: "test-ct", Status: "running"}},
	}

	updated, _ := m.Update(msg)
	updatedModel := updated.(Model)

	if updatedModel.node.Name != "pve" {
		t.Errorf("node name: got %s, want pve", updatedModel.node.Name)
	}
	if len(updatedModel.vms) != 1 {
		t.Errorf("vms: got %d, want 1", len(updatedModel.vms))
	}
	if len(updatedModel.containers) != 1 {
		t.Errorf("containers: got %d, want 1", len(updatedModel.containers))
	}
}

func TestDetailViewTracksIDNotPosition(t *testing.T) {
	m := NewModelWithFakeData()

	// Simulate entering detail view on first VM (position 0)
	firstVM := m.vms[0]
	m.selectedID = firstVM.ID
	m.selectedIsVM = true
	m.currentView = ViewDetail

	// Simulate refresh where list order flips
	flipped := make([]models.VM, len(m.vms))
	copy(flipped, m.vms)
	flipped[0], flipped[len(flipped)-1] = flipped[len(flipped)-1], flipped[0]

	msg := dataRefreshMsg{
		node:       m.node,
		vms:        flipped,
		containers: m.containers,
	}

	updated, _ := m.Update(msg)
	updatedModel := updated.(Model)

	// Cursor should have moved to track the original VM's new position
	if updatedModel.vms[updatedModel.cursor].ID != firstVM.ID {
		t.Errorf("cursor should track VM ID %d after reorder, got ID %d at cursor %d",
			firstVM.ID, updatedModel.vms[updatedModel.cursor].ID, updatedModel.cursor)
	}
}

func TestUpdateHandlesDataRefreshMsgWithError(t *testing.T) {
	m := NewModelWithFakeData()

	msg := dataRefreshMsg{
		err: fmt.Errorf("connection refused"),
	}

	updated, _ := m.Update(msg)
	updatedModel := updated.(Model)

	if updatedModel.lastErr == nil {
		t.Error("expected lastErr to be set, got nil")
	}
}
