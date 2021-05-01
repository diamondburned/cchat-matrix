package matrix

import (
	"context"

	"github.com/diamondburned/cchat"
	"github.com/diamondburned/cchat/services"
	"github.com/diamondburned/cchat/text"
	"github.com/diamondburned/cchat/utils/empty"

	"github.com/chanbakjsd/cchat-matrix/internal/auth"
	"github.com/chanbakjsd/cchat-matrix/internal/rich"
)

func init() {
	services.RegisterService(Service{})
}

const IconSource = "https://raw.githubusercontent.com/vector-im/logos/master/matrix/matrix-favicon-white.png"

type Service struct {
	empty.Service
}

func (Service) ID() string {
	return "com.github.chanbakjsd.cchat-matrix"
}

func (Service) Name(_ context.Context, labeler cchat.LabelContainer) (func(), error) {
	labeler.SetLabel(text.Rich{
		Content: "Matrix",
		Segments: []text.Segment{
			rich.IconSegment{Pos: 0, URL: IconSource, Text: "Matrix logo", Size: 42},
		},
	})

	return func() {}, nil
}

func (Service) Authenticate() []cchat.Authenticator {
	return []cchat.Authenticator{auth.HomeServer{}}
}

func (Service) AsSessionRestorer() cchat.SessionRestorer {
	return SessionRestorer{}
}
