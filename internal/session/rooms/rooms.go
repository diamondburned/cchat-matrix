package rooms

import (
	"context"

	"github.com/chanbakjsd/gotrix"
	"github.com/chanbakjsd/gotrix/event"
	"github.com/chanbakjsd/gotrix/matrix"
	"github.com/diamondburned/cchat"
	"github.com/diamondburned/cchat/text"
	"github.com/diamondburned/cchat/utils/empty"
)

type Room struct {
	empty.Server

	client *gotrix.Client
	id     matrix.RoomID
}

var _ cchat.Server = (*Room)(nil)

func NewRoom(c *gotrix.Client, id matrix.RoomID) *Room {
	return &Room{
		client: c,
		id:     id,
	}
}

func (r *Room) ID() cchat.ID { return cchat.ID(r.id) }

func (r *Room) Name(ctx context.Context, labeler cchat.LabelContainer) (func(), error) {
	client := r.client.WithContext(ctx)

	name, _ := client.RoomName(r.id)
	if name == "" {
		name = string(r.id)
	}

	var segments []text.Segment

	avatarEv, _ := client.RoomState(r.id, event.TypeRoomAvatar, "")
	if avatar, ok := avatarEv.(event.RoomAvatarEvent); ok {
		segments = []text.Segment{roomNameSegment{
			strlen: len(name),
			avatar: newRoomAvatarSegment(avatar),
		}}
	}

	labeler.SetLabel(text.Rich{
		Content:  name,
		Segments: segments,
	})

	// TODO: add event handlers.
	return func() {}, nil
}

type roomNameSegment struct {
	empty.TextSegment

	strlen int
	avatar roomAvatarSegment
}

func (seg roomNameSegment) Bounds() (int, int) { return 0, seg.strlen }

func (seg roomNameSegment) AsAvatarer() text.Avatarer { return seg.avatar }

type roomAvatarSegment struct {
	url  string
	w, h int
}

func newRoomAvatarSegment(event event.RoomAvatarEvent) roomAvatarSegment {
	return roomAvatarSegment{
		url: event.URL,
		w:   event.Info.Width,
		h:   event.Info.Height,
	}
}

func (seg roomAvatarSegment) Avatar() string {
	return seg.url
}

func (seg roomAvatarSegment) AvatarText() string { return "Room Avatar" }

func (seg roomAvatarSegment) AvatarSize() int {
	if seg.w > seg.h {
		return seg.w
	}
	return seg.h
}
