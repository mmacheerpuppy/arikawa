package discord

type User struct {
	ID            Snowflake `json:"id,string"`
	Username      string    `json:"username"`
	Discriminator string    `json:"discriminator"`
	Avatar        Hash      `json:"avatar"`

	// These fields may be omitted

	Bot bool `json:"bot,omitempty"`
	MFA bool `json:"mfa_enabled,omitempty"`

	DiscordSystem bool `json:"system,omitempty"`
	EmailVerified bool `json:"verified,omitempty"`

	Locale string `json:"locale,omitempty"`
	Email  string `json:"email,omitempty"`

	Flags UserFlags `json:"flags,omitempty"`
	Nitro UserNitro `json:"premium_type,omitempty"`
}

type UserFlags uint16

const (
	NoFlag UserFlags = 0

	DiscordEmployee UserFlags = 1 << iota
	DiscordPartner
	HypeSquadEvents
	BugHunter
	HouseBravery
	HouseBrilliance
	HouseBalance
	EarlySupporter
	TeamUser
	System
)

type UserNitro uint8

const (
	NoUserNitro UserNitro = iota
	NitroClassic
	NitroFull
)