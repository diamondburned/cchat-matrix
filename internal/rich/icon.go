package rich

import (
	"github.com/diamondburned/cchat/text"
	"github.com/diamondburned/cchat/utils/empty"
)

type IconSegment struct {
	empty.TextSegment

	Pos  int
	URL  string
	Text string
	Size int
}

var _ text.Segment = (*IconSegment)(nil)

func (seg IconSegment) Bounds() (int, int) { return seg.Pos, seg.Pos }

func (seg IconSegment) AsAvatarer() text.Avatarer { return icon(seg) }

type icon IconSegment

var _ text.Avatarer = (*icon)(nil)

// AvatarText returns the underlying text of the image. Frontends could use this
// for hovering or displaying the text instead of the image.
func (i icon) AvatarText() string { return i.Text }

// AvatarSize returns the requested dimension for the image. This function could
// return (0, 0), which the frontend should use the avatar's dimensions.
func (i icon) AvatarSize() (size int) { return i.Size }

// Avatar returns the URL for the image.
func (i icon) Avatar() (url string) { return i.URL }
