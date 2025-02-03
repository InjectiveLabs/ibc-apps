package mocks

import (
	"context"

	ibcclienttypes "github.com/cosmos/ibc-go/v9/modules/core/02-client/types"
	porttypes "github.com/cosmos/ibc-go/v9/modules/core/05-port/types"
	"github.com/cosmos/ibc-go/v9/modules/core/exported"
)

var _ porttypes.ICS4Wrapper = &ICS4WrapperMock{}

type ICS4WrapperMock struct{}

func (m *ICS4WrapperMock) SendPacket(
	ctx context.Context,
	sourcePort string,
	sourceChannel string,
	timeoutHeight ibcclienttypes.Height,
	timeoutTimestamp uint64,
	data []byte,
) (sequence uint64, err error) {
	return 1, nil
}

func (m *ICS4WrapperMock) WriteAcknowledgement(
	ctx context.Context,
	packet exported.PacketI,
	ack exported.Acknowledgement,
) error {
	return nil
}

func (m *ICS4WrapperMock) GetAppVersion(
	ctx context.Context,
	portID,
	channelID string,
) (string, bool) {
	return "", false
}
