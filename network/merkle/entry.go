/*
 *    Copyright 2018 Insolar
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package merkle

import (
	"github.com/insolar/insolar/core"
	"github.com/insolar/insolar/core/utils"
)

type PulseEntry struct {
	Pulse *core.Pulse
}

func (pe *PulseEntry) hash() []byte {
	return pulseHash(pe.Pulse)
}

type GlobuleEntry struct {
	PulseEntry
	ProofSet      map[core.Node]*PulseProof
	PulseHash     []byte
	PrevCloudHash []byte
	GlobuleIndex  uint32
}

func (ge *GlobuleEntry) hash() []byte {
	return nil
}

type CloudEntry struct {
	ProofSet      []*GlobuleProof
	PrevCloudHash []byte
	// TODO: implement later
	// ProofSet map[core.Globule]*GlobuleProof
}

func (ce *CloudEntry) hash() []byte {
	var result [][]byte

	for _, proof := range ce.ProofSet {
		globuleInfoHash := globuleInfoHash(ce.PrevCloudHash, proof.GlobuleIndex, proof.NodeCount)
		globuleHash := globuleHash(globuleInfoHash, proof.NodeRoot)
		result = append(result, globuleHash)
	}

	if len(result)%2 == 1 {
		result = append(result, utils.UInt32ToBytes(reserved))
	}

	mt, err := fromList(result)
	if err != nil {
		panic(err.Error())
	}

	return mt.MerkleRoot()
}

type nodeEntry struct {
	PulseEntry
	PulseProof *PulseProof
	Node       core.Node
}

func (ne *nodeEntry) hash() []byte {
	pulseHash := ne.PulseEntry.hash()
	nodeInfoHash := nodeInfoHash(pulseHash, ne.PulseProof.StateHash)
	return nodeHash(ne.PulseProof.Signature, nodeInfoHash)
}

func nodeEntryByRole(entry *GlobuleEntry) map[core.NodeRole][]*nodeEntry {
	roleMap := make(map[core.NodeRole][]*nodeEntry)
	for node, pulseProof := range entry.ProofSet {
		role := node.Role()
		roleMap[role] = append(roleMap[role], &nodeEntry{
			Node:       node,
			PulseProof: pulseProof,
		})
	}
	return roleMap
}
