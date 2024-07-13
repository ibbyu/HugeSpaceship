// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: user.sql

package sqlc

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const getUserByName = `-- name: GetUserByName :one
SELECT
    users.id, users.username, users.avatar_hash, users.psn_uid, users.rpcn_uid, users.bio, users.comments_enabled, users.location_x, users.location_y, users.entitled_slots, users.planet_lbp2, users.planet_lbp3, users.planet_lbp_vita, users.planet_cc, users.boo_icon, users.meh_icon, users.yay_icon, users.level_visibility, users.profile_visibility,
    users.entitled_slots - COUNT(s) AS free_slots,
    COUNT(s) AS used_slots
FROM users LEFT JOIN slots AS s ON s.uploader = users.id
WHERE username = $1 GROUP BY users.id LIMIT 1
`

type GetUserByNameRow struct {
	ID                uuid.UUID
	Username          pgtype.Text
	AvatarHash        pgtype.Text
	PsnUid            string
	RpcnUid           string
	Bio               string
	CommentsEnabled   bool
	LocationX         int32
	LocationY         int32
	EntitledSlots     int32
	PlanetLbp2        pgtype.Text
	PlanetLbp3        pgtype.Text
	PlanetLbpVita     pgtype.Text
	PlanetCc          pgtype.Text
	BooIcon           pgtype.Text
	MehIcon           pgtype.Text
	YayIcon           pgtype.Text
	LevelVisibility   NullPrivacy
	ProfileVisibility NullPrivacy
	FreeSlots         int32
	UsedSlots         int64
}

func (q *Queries) GetUserByName(ctx context.Context, name pgtype.Text) (GetUserByNameRow, error) {
	row := q.db.QueryRow(ctx, getUserByName, name)
	var i GetUserByNameRow
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.AvatarHash,
		&i.PsnUid,
		&i.RpcnUid,
		&i.Bio,
		&i.CommentsEnabled,
		&i.LocationX,
		&i.LocationY,
		&i.EntitledSlots,
		&i.PlanetLbp2,
		&i.PlanetLbp3,
		&i.PlanetLbpVita,
		&i.PlanetCc,
		&i.BooIcon,
		&i.MehIcon,
		&i.YayIcon,
		&i.LevelVisibility,
		&i.ProfileVisibility,
		&i.FreeSlots,
		&i.UsedSlots,
	)
	return i, err
}
