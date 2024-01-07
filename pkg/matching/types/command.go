package types

import (
	"HugeSpaceship/pkg/matching/types/commands"
)

type MatchCommand string

const (
	CreateRoom   MatchCommand = "CreateRoomCommand"
	FindBestRoom MatchCommand = "FindBestRoom"
)

type RawMatchCommand interface {
	commands.CreateRoom
}
