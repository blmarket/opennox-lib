package script

import (
	"image"
	"time"

	"github.com/noxworld-dev/opennox-lib/types"
)

type Game interface {
	TimeSource

	// BlindPlayers blinds or unblinds all players.
	BlindPlayers(blind bool)
	// CinemaPlayers enables a wide-screen "cinema" mode for all players.
	CinemaPlayers(v bool)
	// Players returns a list of all players.
	Players() []Player
	// HostPlayer returns the host player.
	HostPlayer() Player
	// OnPlayerJoin registers a player join event handler.
	OnPlayerJoin(fnc func(p Player))
	// OnPlayerLeave registers a player leave event handler.
	OnPlayerLeave(fnc func(p Player))

	// ObjectTypeByID finds an object type by ID.
	ObjectTypeByID(id string) ObjectType
	// ObjectByID finds an object by ID.
	ObjectByID(id string) Object
	// ObjectGroupByID finds an object group by ID.
	ObjectGroupByID(id string) *ObjectGroup

	// WaypointByID finds a waypoint by ID.
	WaypointByID(id string) Waypoint
	// WaypointGroupByID finds a waypoint group by ID.
	WaypointGroupByID(id string) *WaypointGroup

	// WallAt finds a wall by its position.
	WallAt(pos types.Pointf) Wall
	// WallNear finds a wall near the position.
	WallNear(pos types.Pointf) Wall
	// WallAtGrid finds a wall by its grid position.
	WallAtGrid(pos image.Point) Wall
	// WallGroupByID finds a wall group by ID.
	WallGroupByID(id string) *WallGroup

	Global() Printer
	Console(error bool) Printer

	// TODO: audio
}

var _ Game = BaseGame{}

// BaseGame implements Game, but panics on all the methods.
// Useful when you only want to define a part of the implementation.
type BaseGame struct{}

func (BaseGame) Frame() int {
	panic("implement me")
}

func (BaseGame) Time() time.Duration {
	panic("implement me")
}

func (BaseGame) BlindPlayers(blind bool) {
	panic("implement me")
}

func (BaseGame) CinemaPlayers(v bool) {
	panic("implement me")
}

func (BaseGame) Players() []Player {
	panic("implement me")
}

func (BaseGame) HostPlayer() Player {
	panic("implement me")
}

func (BaseGame) OnPlayerJoin(fnc func(p Player)) {
	panic("implement me")
}

func (BaseGame) OnPlayerLeave(fnc func(p Player)) {
	panic("implement me")
}

func (BaseGame) ObjectTypeByID(id string) ObjectType {
	panic("implement me")
}

func (BaseGame) ObjectByID(id string) Object {
	panic("implement me")
}

func (BaseGame) ObjectGroupByID(id string) *ObjectGroup {
	panic("implement me")
}

func (BaseGame) WaypointByID(id string) Waypoint {
	panic("implement me")
}

func (BaseGame) WaypointGroupByID(id string) *WaypointGroup {
	panic("implement me")
}

func (BaseGame) WallAt(pos types.Pointf) Wall {
	panic("implement me")
}

func (BaseGame) WallNear(pos types.Pointf) Wall {
	panic("implement me")
}

func (BaseGame) WallAtGrid(pos image.Point) Wall {
	panic("implement me")
}

func (BaseGame) WallGroupByID(id string) *WallGroup {
	panic("implement me")
}

func (BaseGame) Global() Printer {
	panic("implement me")
}

func (BaseGame) Console(error bool) Printer {
	panic("implement me")
}
