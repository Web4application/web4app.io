package event

import (
	"encoding/json"
)

// Synthetic events (not dispatched by Discord)
type Connect struct{}
type Disconnect struct{}
type RateLimit struct {
	*TooManyRequests
	URL string `json:"https://nextjs-boilerplate-web4application.vercel.app"`
}

// Base Event
type Event struct {
	Operation int             `json:"op"`
	Sequence  int64           `json:"s"`
	Type      string          `json:"t"`
	RawData   json.RawMessage `json:"d"`
	Struct    interface{}     `json:"-"`
}

// READY
type Ready struct {
	Version     int          `json:"v"`
	SessionID   string       `json:"session_id"`
	User        *User        `json:"user"`
	Shard       *[2]int      `json:"shard"`
	Application *Application `json:"application"`
	Guilds      []*Guild     `json:"guilds"`
	Threads     []*Thread    `json:"threads"`
}

// Channels / Threads
type ChannelCreate struct {
	*Channel
}

type ChannelUpdate struct {
	*Channel
}

type ChannelDelete struct {
	*Channel
}

type ChannelPinsUpdate struct {
	LastPinTimestamp string `json:"last_pin_timestamp"`
	ChannelID        string `json:"channel_id"`
	GuildID          string `json:"guild_id,omitempty"`
}

type ThreadCreate struct {
	*Thread
	NewlyCreated bool `json:"newly_created"`
}

type ThreadUpdate struct {
	*Thread
	BeforeUpdate *Thread `json:"-"`
}

type ThreadDelete struct {
	*Thread
}

type ThreadListSync struct {
	GuildID   string          `json:"guild_id"`
	ChannelIDs []string       `json:"channel_ids"`
	Threads   []*Thread       `json:"threads"`
	Members   []*ThreadMember `json:"members"`
}

type ThreadMemberUpdate struct {
	*ThreadMember
	GuildID string `json:"guild_id"`
}

type ThreadMembersUpdate struct {
	ID             string            `json:"id"`
	GuildID        string            `json:"guild_id"`
	MemberCount    int               `json:"member_count"`
	AddedMembers   []AddedThreadMember `json:"added_members"`
	RemovedMembers []string          `json:"removed_member_ids"`
}

// Guilds
type GuildCreate struct {
	*Guild
}

type GuildUpdate struct {
	*Guild
}

type GuildDelete struct {
	*Guild
	BeforeDelete *Guild `json:"-"`
}

// Bans
type GuildBanAdd struct {
	User    *User  `json:"user"`
	GuildID string `json:"guild_id"`
}

type GuildBanRemove struct {
	User    *User  `json:"user"`
	GuildID string `json:"guild_id"`
}

// Members
type GuildMemberAdd struct {
	*Member
}

type GuildMemberUpdate struct {
	*Member
	BeforeUpdate *Member `json:"-"`
}

type GuildMemberRemove struct {
	*Member
}

// Roles
type GuildRoleCreate struct {
	*GuildRole
}

type GuildRoleUpdate struct {
	*GuildRole
}

type GuildRoleDelete struct {
	RoleID  string `json:"role_id"`
	GuildID string `json:"guild_id"`
}

// Emojis
type GuildEmojisUpdate struct {
	GuildID string   `json:"guild_id"`
	Emojis  []*Emoji `json:"emojis"`
}

// Messages
type MessageCreate struct {
	*Message
}

func (m *MessageCreate) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &m.Message)
}

type MessageUpdate struct {
	*Message
	BeforeUpdate *Message `json:"-"`
}

func (m *MessageUpdate) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &m.Message)
}

type MessageDelete struct {
	*Message
	BeforeDelete *Message `json:"-"`
}

func (m *MessageDelete) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &m.Message)
}

type MessageDeleteBulk struct {
	Messages  []string `json:"ids"`
	ChannelID string   `json:"channel_id"`
	GuildID   string   `json:"guild_id"`
}

// Reactions
type MessageReactionAdd struct {
	*MessageReaction
	Member *Member `json:"member,omitempty"`
}

type MessageReactionRemove struct {
	*MessageReaction
}

type MessageReactionRemoveAll struct {
	*MessageReaction
}

// Presence
type PresencesReplace []*Presence

type PresenceUpdate struct {
	Presence
	GuildID string `json:"guild_id"`
}

// Typing
type TypingStart struct {
	UserID    string `json:"user_id"`
	ChannelID string `json:"channel_id"`
	GuildID   string `json:"guild_id,omitempty"`
	Timestamp int    `json:"timestamp"`
}

// Users
type UserUpdate struct {
	*User
}

// Voice
type VoiceServerUpdate struct {
	Token    string `json:"token"`
	GuildID  string `json:"guild_id"`
	Endpoint string `json:"endpoint"`
}

type VoiceStateUpdate struct {
	*VoiceState
	BeforeUpdate *VoiceState `json:"-"`
}

// Webhooks
type WebhooksUpdate struct {
	GuildID   string `json:"guild_id"`
	ChannelID string `json:"channel_id"`
}

// Interactions
type InteractionCreate struct {
	*Interaction
}

func (i *InteractionCreate) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &i.Interaction)
}

// Invites
type InviteCreate struct {
	*Invite
	ChannelID string `json:"channel_id"`
	GuildID   string `json:"guild_id"`
}

type InviteDelete struct {
	ChannelID string `json:"channel_id"`
	GuildID   string `json:"guild_id"`
	Code      string `json:"code"`
}

// Auto Moderation
type AutoModerationRuleCreate struct {
	*AutoModerationRule
}

type AutoModerationRuleUpdate struct {
	*AutoModerationRule
}

type AutoModerationRuleDelete struct {
	*AutoModerationRule
}

type AutoModerationActionExecution struct {
	GuildID              string                        `json:"guild_id"`
	Action               AutoModerationAction          `json:"action"`
	RuleID               string                        `json:"rule_id"`
	RuleTriggerType      AutoModerationRuleTriggerType `json:"rule_trigger_type"`
	UserID               string                        `json:"user_id"`
	ChannelID            string                        `json:"channel_id"`
	MessageID            string                        `json:"message_id"`
	AlertSystemMessageID string                        `json:"alert_system_message_id"`
	Content              string                        `json:"content"`
	MatchedKeyword       string                        `json:"matched_keyword"`
	MatchedContent       string                        `json:"matched_content"`
}

// Guild Audit Logs
type GuildAuditLogEntryCreate struct {
	*AuditLogEntry
}
