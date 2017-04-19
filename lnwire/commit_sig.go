package lnwire

import (
	"io"

	"github.com/roasbeef/btcd/btcec"
)

// CommitSig is sent by either side to stage any pending HTLC's in the
// receiver's pending set into a new commitment state.  Implicitly, the new
// commitment transaction constructed which has been signed by CommitSig
// includes all HTLC's in the remote node's pending set. A CommitSig message
// may be sent after a series of UpdateAddHTLC/UpdateFufillHTLC messages in
// order to batch add several HTLC's with a single signature covering all
// implicitly accepted HTLC's.
type CommitSig struct {
	// ChanID uniquely identifies to which currently active channel this
	// CommitSig applies to.
	ChanID ChannelID

	// CommitSig is Alice's signature for Bob's new commitment transaction.
	// Alice is able to send this signature without requesting any
	// additional data due to the piggybacking of Bob's next revocation
	// hash in his prior RevokeAndAck message, as well as the canonical
	// ordering used for all inputs/outputs within commitment transactions.
	CommitSig *btcec.Signature

	// TODO(roasbeef): add HTLC sigs after state machine is updated to
	// support that
}

// NewCommitSig creates a new empty CommitSig message.
func NewCommitSig() *CommitSig {
	return &CommitSig{}
}

// A compile time check to ensure CommitSig implements the lnwire.Message
// interface.
var _ Message = (*CommitSig)(nil)

// Decode deserializes a serialized CommitSig message stored in the
// passed io.Reader observing the specified protocol version.
//
// This is part of the lnwire.Message interface.
func (c *CommitSig) Decode(r io.Reader, pver uint32) error {
	return readElements(r,
		&c.ChanID,
		&c.CommitSig,
	)
}

// Encode serializes the target CommitSig into the passed io.Writer
// observing the protocol version specified.
//
// This is part of the lnwire.Message interface.
func (c *CommitSig) Encode(w io.Writer, pver uint32) error {
	return writeElements(w,
		c.ChanID,
		c.CommitSig,
	)
}

// Command returns the integer uniquely identifying this message type on the
// wire.
//
// This is part of the lnwire.Message interface.
func (c *CommitSig) Command() uint32 {
	return CmdCommitSig
}

// MaxPayloadLength returns the maximum allowed payload size for a
// CommitSig complete message observing the specified protocol version.
//
// This is part of the lnwire.Message interface.
func (c *CommitSig) MaxPayloadLength(uint32) uint32 {
	// 32 + 64
	return 96
}
