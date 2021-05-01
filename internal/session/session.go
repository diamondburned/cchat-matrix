package session

import (
	"context"
	"sort"

	"github.com/chanbakjsd/cchat-matrix/internal/session/rooms"
	"github.com/chanbakjsd/gotrix"
	"github.com/diamondburned/cchat"
	"github.com/diamondburned/cchat/text"
	"github.com/diamondburned/cchat/utils/empty"
	"github.com/pkg/errors"
)

type Session struct {
	empty.Session
	*gotrix.Client
}

func New(cli *gotrix.Client) (cchat.Session, error) {
	s := &Session{
		Client: cli,
	}

	if err := s.Open(); err != nil {
		return nil, err
	}

	return s, nil
}

func (s *Session) ID() cchat.ID {
	return cchat.ID(s.UserID)
}

func (s *Session) Name(ctx context.Context, labeler cchat.LabelContainer) (func(), error) {
	labeler.SetLabel(text.Plain(string(s.UserID)))
	return func() {}, nil
}

func (s *Session) Servers(c cchat.ServersContainer) (func(), error) {
	roomIDs, err := s.Rooms()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get rooms")
	}

	sort.Slice(roomIDs, func(i, j int) bool {
		return roomIDs[i] < roomIDs[j]
	})

	roomList := make([]cchat.Server, len(roomIDs))
	for i, roomID := range roomIDs {
		roomList[i] = rooms.NewRoom(s.Client, roomID)
	}

	c.SetServers(roomList)

	return func() {}, nil
}

func (s *Session) Columnate() bool { return false }

func (s *Session) Disconnect() error {
	return s.Close()
}

func (s *Session) AsSessionSaver() cchat.SessionSaver {
	return s
}

func (s *Session) SaveSession() map[string]string {
	return map[string]string{
		"homeserver":  s.HomeServer,
		"accessToken": s.AccessToken,
		"deviceID":    string(s.DeviceID),
		"userID":      string(s.UserID),
	}
}
