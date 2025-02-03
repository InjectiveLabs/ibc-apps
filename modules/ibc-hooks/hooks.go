package ibc_hooks

import (
	"context"
	// external libraries
	sdk "github.com/cosmos/cosmos-sdk/types"
	ibcclienttypes "github.com/cosmos/ibc-go/v9/modules/core/02-client/types"
	// ibc-go
	channeltypes "github.com/cosmos/ibc-go/v9/modules/core/04-channel/types"
	ibcexported "github.com/cosmos/ibc-go/v9/modules/core/exported"
)

type Hooks interface{}

type OnChanOpenInitOverrideHooks interface {
	OnChanOpenInitOverride(im IBCMiddleware, ctx context.Context, order channeltypes.Order, connectionHops []string, portID string, channelID string, counterparty channeltypes.Counterparty, version string) (string, error)
}
type OnChanOpenInitBeforeHooks interface {
	OnChanOpenInitBeforeHook(ctx context.Context, order channeltypes.Order, connectionHops []string, portID string, channelID string, counterparty channeltypes.Counterparty, version string)
}
type OnChanOpenInitAfterHooks interface {
	OnChanOpenInitAfterHook(ctx context.Context, order channeltypes.Order, connectionHops []string, portID string, channelID string, counterparty channeltypes.Counterparty, version string, finalVersion string, err error)
}

// OnChanOpenTry Hooks
type OnChanOpenTryOverrideHooks interface {
	OnChanOpenTryOverride(im IBCMiddleware, ctx context.Context, order channeltypes.Order, connectionHops []string, portID, channelID string, counterparty channeltypes.Counterparty, counterpartyVersion string) (string, error)
}
type OnChanOpenTryBeforeHooks interface {
	OnChanOpenTryBeforeHook(ctx context.Context, order channeltypes.Order, connectionHops []string, portID, channelID string, counterparty channeltypes.Counterparty, counterpartyVersion string)
}
type OnChanOpenTryAfterHooks interface {
	OnChanOpenTryAfterHook(ctx context.Context, order channeltypes.Order, connectionHops []string, portID, channelID string, counterparty channeltypes.Counterparty, counterpartyVersion string, version string, err error)
}

// OnChanOpenAck Hooks
type OnChanOpenAckOverrideHooks interface {
	OnChanOpenAckOverride(im IBCMiddleware, ctx context.Context, portID, channelID string, counterpartyChannelID string, counterpartyVersion string) error
}
type OnChanOpenAckBeforeHooks interface {
	OnChanOpenAckBeforeHook(ctx context.Context, portID, channelID string, counterpartyChannelID string, counterpartyVersion string)
}
type OnChanOpenAckAfterHooks interface {
	OnChanOpenAckAfterHook(ctx context.Context, portID, channelID string, counterpartyChannelID string, counterpartyVersion string, err error)
}

// OnChanOpenConfirm Hooks
type OnChanOpenConfirmOverrideHooks interface {
	OnChanOpenConfirmOverride(im IBCMiddleware, ctx context.Context, portID, channelID string) error
}
type OnChanOpenConfirmBeforeHooks interface {
	OnChanOpenConfirmBeforeHook(ctx context.Context, portID, channelID string)
}
type OnChanOpenConfirmAfterHooks interface {
	OnChanOpenConfirmAfterHook(ctx context.Context, portID, channelID string, err error)
}

// OnChanCloseInit Hooks
type OnChanCloseInitOverrideHooks interface {
	OnChanCloseInitOverride(im IBCMiddleware, ctx context.Context, portID, channelID string) error
}
type OnChanCloseInitBeforeHooks interface {
	OnChanCloseInitBeforeHook(ctx context.Context, portID, channelID string)
}
type OnChanCloseInitAfterHooks interface {
	OnChanCloseInitAfterHook(ctx context.Context, portID, channelID string, err error)
}

// OnChanCloseConfirm Hooks
type OnChanCloseConfirmOverrideHooks interface {
	OnChanCloseConfirmOverride(im IBCMiddleware, ctx context.Context, portID, channelID string) error
}
type OnChanCloseConfirmBeforeHooks interface {
	OnChanCloseConfirmBeforeHook(ctx context.Context, portID, channelID string)
}
type OnChanCloseConfirmAfterHooks interface {
	OnChanCloseConfirmAfterHook(ctx context.Context, portID, channelID string, err error)
}

// OnRecvPacket Hooks
type OnRecvPacketOverrideHooks interface {
	OnRecvPacketOverride(im IBCMiddleware, ctx context.Context, channelVersion string, packet channeltypes.Packet, relayer sdk.AccAddress) ibcexported.Acknowledgement
}
type OnRecvPacketBeforeHooks interface {
	OnRecvPacketBeforeHook(ctx context.Context, channelVersion string, packet channeltypes.Packet, relayer sdk.AccAddress)
}
type OnRecvPacketAfterHooks interface {
	OnRecvPacketAfterHook(ctx context.Context, channelVersion string, packet channeltypes.Packet, relayer sdk.AccAddress, ack ibcexported.Acknowledgement)
}

// OnAcknowledgementPacket Hooks
type OnAcknowledgementPacketOverrideHooks interface {
	OnAcknowledgementPacketOverride(im IBCMiddleware, ctx context.Context, channelVersion string, packet channeltypes.Packet, acknowledgement []byte, relayer sdk.AccAddress) error
}
type OnAcknowledgementPacketBeforeHooks interface {
	OnAcknowledgementPacketBeforeHook(ctx context.Context, channelVersion string, packet channeltypes.Packet, acknowledgement []byte, relayer sdk.AccAddress)
}
type OnAcknowledgementPacketAfterHooks interface {
	OnAcknowledgementPacketAfterHook(ctx context.Context, channelVersion string, packet channeltypes.Packet, acknowledgement []byte, relayer sdk.AccAddress, err error)
}

// OnTimeoutPacket Hooks
type OnTimeoutPacketOverrideHooks interface {
	OnTimeoutPacketOverride(im IBCMiddleware, ctx context.Context, channelVersion string, packet channeltypes.Packet, relayer sdk.AccAddress) error
}
type OnTimeoutPacketBeforeHooks interface {
	OnTimeoutPacketBeforeHook(ctx context.Context, channelVersion string, packet channeltypes.Packet, relayer sdk.AccAddress)
}
type OnTimeoutPacketAfterHooks interface {
	OnTimeoutPacketAfterHook(ctx context.Context, channelVersion string, packet channeltypes.Packet, relayer sdk.AccAddress, err error)
}

// SendPacket Hooks
type SendPacketOverrideHooks interface {
	SendPacketOverride(i ICS4Middleware, ctx context.Context, sourcePort string, sourceChannel string, timeoutHeight ibcclienttypes.Height, timeoutTimestamp uint64, data []byte) (sequence uint64, err error)
}
type SendPacketBeforeHooks interface {
	SendPacketBeforeHook(ctx context.Context, sourcePort string, sourceChannel string, timeoutHeight ibcclienttypes.Height, timeoutTimestamp uint64, data []byte)
}
type SendPacketAfterHooks interface {
	SendPacketAfterHook(ctx context.Context, sourcePort string, sourceChannel string, timeoutHeight ibcclienttypes.Height, timeoutTimestamp uint64, data []byte, err error)
}

// WriteAcknowledgement Hooks
type WriteAcknowledgementOverrideHooks interface {
	WriteAcknowledgementOverride(i ICS4Middleware, ctx context.Context, packet ibcexported.PacketI, ack ibcexported.Acknowledgement) error
}
type WriteAcknowledgementBeforeHooks interface {
	WriteAcknowledgementBeforeHook(ctx context.Context, packet ibcexported.PacketI, ack ibcexported.Acknowledgement)
}
type WriteAcknowledgementAfterHooks interface {
	WriteAcknowledgementAfterHook(ctx context.Context, packet ibcexported.PacketI, ack ibcexported.Acknowledgement, err error)
}

// GetAppVersion Hooks
type GetAppVersionOverrideHooks interface {
	GetAppVersionOverride(i ICS4Middleware, ctx sdk.Context, portID, channelID string) (string, bool)
}
type GetAppVersionBeforeHooks interface {
	GetAppVersionBeforeHook(ctx sdk.Context, portID, channelID string)
}
type GetAppVersionAfterHooks interface {
	GetAppVersionAfterHook(ctx sdk.Context, portID, channelID string, result string, success bool)
}
