//
// Modified BSD 3-Clause Clear License
//
// Copyright (c) 2019 Insolar Technologies GmbH
//
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without modification,
// are permitted (subject to the limitations in the disclaimer below) provided that
// the following conditions are met:
//  * Redistributions of source code must retain the above copyright notice, this list
//    of conditions and the following disclaimer.
//  * Redistributions in binary form must reproduce the above copyright notice, this list
//    of conditions and the following disclaimer in the documentation and/or other materials
//    provided with the distribution.
//  * Neither the name of Insolar Technologies GmbH nor the names of its contributors
//    may be used to endorse or promote products derived from this software without
//    specific prior written permission.
//
// NO EXPRESS OR IMPLIED LICENSES TO ANY PARTY'S PATENT RIGHTS ARE GRANTED
// BY THIS LICENSE. THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS
// AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES,
// INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY
// AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL
// THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT,
// INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING,
// BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS
// OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
// ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
// Notwithstanding any other provisions of this license, it is prohibited to:
//    (a) use this software,
//
//    (b) prepare modifications and derivative works of this software,
//
//    (c) distribute this software (including without limitation in source code, binary or
//        object code form), and
//
//    (d) reproduce copies of this software
//
//    for any commercial purposes, and/or
//
//    for the purposes of making available this software to third parties as a service,
//    including, without limitation, any software-as-a-service, platform-as-a-service,
//    infrastructure-as-a-service or other similar online service, irrespective of
//    whether it competes with the products or services of Insolar Technologies GmbH.
//

package tests

import (
	"fmt"

	"github.com/insolar/insolar/network/consensus/adapters"
	"github.com/insolar/insolar/network/consensus/gcpv2/nodeset"
	"github.com/insolar/insolar/network/consensus/gcpv2/packets"

	common2 "github.com/insolar/insolar/network/consensus/gcpv2/common"

	"github.com/insolar/insolar/network/consensus/common"
)

type EmuPacketWrapper struct {
	parser packets.PacketParser
}

func UnwrapPacketParser(payload interface{}) packets.PacketParser {
	if v, ok := payload.(EmuPacketWrapper); ok {
		return v.parser
	}
	return nil
}

func WrapPacketParser(payload packets.PacketParser) interface{} {
	return EmuPacketWrapper{parser: payload}
}

func (v EmuPacketWrapper) String() string {
	return fmt.Sprintf("Wrap{%v}", v.parser)
}

var _ common.SignedEvidenceHolder = &basePacket{}

type basePacket struct {
	src       common.ShortNodeID
	tgt       common.ShortNodeID
	nodeCount uint16
	mp        common2.MembershipProfile

	sd common.SignedDigest
}

func (r *basePacket) GetRequestedPower() common2.MemberPower {
	return r.mp.RequestedPower
}

func (r *basePacket) IsLeaving() bool {
	return false
}

func (r *basePacket) GetLeaveReason() uint32 {
	return 0
}

func (r *basePacket) GetJoinerID() common.ShortNodeID {
	return 0
}

func (r *basePacket) GetJoinerAnnouncement() packets.JoinerAnnouncementReader {
	return nil
}

func (r *basePacket) GetNodeStateHashEvidence() common2.NodeStateHashEvidence {
	return r.mp.StateEvidence
}

func (r *basePacket) GetAnnouncementSignature() common2.MemberAnnouncementSignature {
	return r.mp.AnnounceSignature
}

func (r *basePacket) GetNodeID() common.ShortNodeID {
	return r.tgt
}

func (r *basePacket) GetNodeRank() common2.MembershipRank {
	return common2.NewMembershipRank(r.mp.Power, r.mp.Index, r.nodeCount, 0)
}

func (r *basePacket) GetAnnouncementReader() packets.MembershipAnnouncementReader {
	return r
}

func (r *basePacket) GetEvidence() common.SignedData {
	v := common.NewBits64(0)
	d := common.NewDigest(&v, "stub")
	s := common.NewSignature(&v, "stub")
	return common.NewSignedData(&v, d, s)
}

func (r *basePacket) GetSourceId() common.ShortNodeID {
	return r.src
}

func (r *basePacket) GetReceiverId() common.ShortNodeID {
	return r.tgt
}

func (r *basePacket) GetTargetID() common.ShortNodeID {
	return r.tgt
}

func (r *basePacket) GetPacketSignature() common.SignedDigest {
	return r.sd
}

func (*basePacket) IsPulsePacket() bool {
	return false
}

func (r *basePacket) GetPulsePacket() packets.PulsePacketReader {
	return nil
}

func (r *basePacket) AsPhase0Packet() packets.Phase0PacketReader {
	return nil
}

func (r *basePacket) AsPhase1Packet() packets.Phase1PacketReader {
	return nil
}

func (r *basePacket) AsPhase2Packet() packets.Phase2PacketReader {
	return nil
}

func (r *basePacket) AsPhase3Packet() packets.Phase3PacketReader {
	return nil
}

func (r *basePacket) GetEvidenceSignature() common.SignedDigest {
	return r.sd
}

func (r *basePacket) GetPulseDataEvidence() common.SignedEvidenceHolder {
	return r
}

func (r *basePacket) String() string {
	return fmt.Sprintf("s:%v, t:%v", r.src, r.tgt)
}

var _ packets.Phase0PacketReader = &EmuPhase0NetPacket{}
var _ packets.MemberPacketReader = &EmuPhase0NetPacket{}
var _ packets.PacketParser = &EmuPhase0NetPacket{}
var _ emuPackerCloner = &EmuPhase0NetPacket{}

type EmuPhase0NetPacket struct {
	basePacket
	pulsePacket common2.OriginalPulsarPacket
	pn          common.PulseNumber
}

func (r *EmuPhase0NetPacket) GetPacketType() packets.PacketType {
	return packets.PacketPhase0
}

func (r *EmuPhase0NetPacket) GetMemberPacket() packets.MemberPacketReader {
	return r
}

func (r *EmuPhase0NetPacket) AsPhase0Packet() packets.Phase0PacketReader {
	return r
}

func (r *EmuPhase0NetPacket) GetPulseNumber() common.PulseNumber {
	if r.pulsePacket == nil {
		return r.pn
	}
	return r.pulsePacket.(*adapters.PulsePacketReader).GetPulseData().PulseNumber
}

func (r *EmuPhase0NetPacket) GetEmbeddedPulsePacket() packets.PulsePacketReader {
	return r.pulsePacket.(*adapters.PulsePacketReader)
}

func (r *EmuPhase0NetPacket) String() string {
	return fmt.Sprintf("ph:0 %v, pulsePkt: {%v}", r.basePacket.String(), r.pulsePacket)
}

var _ packets.Phase1PacketReader = &EmuPhase1NetPacket{}
var _ packets.MemberPacketReader = &EmuPhase1NetPacket{}
var _ packets.PacketParser = &EmuPhase1NetPacket{}

type EmuPhase1NetPacket struct {
	EmuPhase0NetPacket
	isRequest bool
	// packetType uint8 // to reuse this type for Phase1 and Phase1Req
}

func (r *EmuPhase1NetPacket) GetCloudIntroduction() packets.CloudIntroductionReader {
	panic("implement me")
}

func (r *EmuPhase1NetPacket) GetFullIntroduction() packets.FullIntroductionReader {
	panic("implement me")
}

func (r *EmuPhase1NetPacket) GetNodeClaimsSignature() common2.MemberAnnouncementSignature {
	return r.mp.AnnounceSignature
}

func (r *EmuPhase1NetPacket) String() string {
	prefix := ""
	if r.isRequest {
		prefix = "rq"
	}
	return fmt.Sprintf("ph:1%s %s pulsePkt:{%v} mp:{%v} nc:%d", prefix, r.basePacket.String(), r.pulsePacket, r.mp, r.nodeCount)
}

func (r *EmuPhase1NetPacket) GetPacketType() packets.PacketType {
	if r.isRequest {
		return packets.PacketReqPhase1
	} else {
		return packets.PacketPhase1
	}
}

func (r *EmuPhase1NetPacket) AsPhase0Packet() packets.Phase0PacketReader {
	return nil
}

func (r *EmuPhase1NetPacket) AsPhase1Packet() packets.Phase1PacketReader {
	return r
}

func (r *EmuPhase1NetPacket) GetNodeStateHashEvidence() common2.NodeStateHashEvidence {
	return r.mp.StateEvidence
}

func (r *EmuPhase1NetPacket) HasPulseData() bool {
	return r.pulsePacket != nil
}

func (r *EmuPhase1NetPacket) GetMemberPacket() packets.MemberPacketReader {
	return r
}

var _ packets.Phase2PacketReader = &EmuPhase2NetPacket{}
var _ packets.MemberPacketReader = &EmuPhase2NetPacket{}
var _ packets.PacketParser = &EmuPhase2NetPacket{}

type EmuPhase2NetPacket struct {
	basePacket
	pulseNumber   common.PulseNumber
	neighbourhood []packets.MembershipAnnouncementReader
}

func (r *EmuPhase2NetPacket) GetBriefIntroduction() packets.BriefIntroductionReader {
	panic("implement me")
}

func (r *EmuPhase2NetPacket) String() string {
	return fmt.Sprintf("ph:2 %s pn:%v mp:{%v} nc:%d ngbh:%v", r.basePacket.String(), r.pulseNumber, r.mp, r.nodeCount, r.neighbourhood)
}

func (r *EmuPhase2NetPacket) GetNeighbourhood() []packets.MembershipAnnouncementReader {
	return r.neighbourhood
}

func (r *EmuPhase2NetPacket) GetPacketType() packets.PacketType {
	return packets.PacketPhase2
}

func (r *EmuPhase2NetPacket) AsPhase2Packet() packets.Phase2PacketReader {
	return r
}

func (r *EmuPhase2NetPacket) GetPulseNumber() common.PulseNumber {
	return r.pulseNumber
}

func (r *EmuPhase2NetPacket) GetMemberPacket() packets.MemberPacketReader {
	return r
}

var _ packets.Phase3PacketReader = &EmuPhase3NetPacket{}
var _ packets.MemberPacketReader = &EmuPhase3NetPacket{}
var _ packets.PacketParser = &EmuPhase3NetPacket{}

type EmuPhase3NetPacket struct {
	basePacket
	pulseNumber common.PulseNumber
	bitset      nodeset.NodeBitset
	gshTrusted  common2.GlobulaStateHash
	gshDoubted  common2.GlobulaStateHash
}

func (r *EmuPhase3NetPacket) String() string {
	return fmt.Sprintf("ph:3 %s, pn:%v, set:%v, gshT:%v, gshD:%v", r.basePacket.String(), r.pulseNumber,
		r.bitset, r.gshTrusted, r.gshDoubted)
}

func (r *EmuPhase3NetPacket) GetBitset() nodeset.NodeBitset {
	return r.bitset
}

func (r *EmuPhase3NetPacket) GetTrustedGsh() common2.GlobulaStateHash {
	return r.gshTrusted
}

func (r *EmuPhase3NetPacket) GetDoubtedGsh() common2.GlobulaStateHash {
	return r.gshDoubted
}

func (r *EmuPhase3NetPacket) GetTrustedCshEvidence() common.SignedEvidenceHolder {
	return r
}

func (r *EmuPhase3NetPacket) GetDoubtedCshEvidence() common.SignedEvidenceHolder {
	return r
}

func (r *EmuPhase3NetPacket) GetPacketType() packets.PacketType {
	return packets.PacketPhase3
}

func (r *EmuPhase3NetPacket) AsPhase3Packet() packets.Phase3PacketReader {
	return r
}

func (r *EmuPhase3NetPacket) GetPulseNumber() common.PulseNumber {
	return r.pulseNumber
}

func (r *EmuPhase3NetPacket) GetMemberPacket() packets.MemberPacketReader {
	return r
}
