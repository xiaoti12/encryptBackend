package flowMap

import (
	"container/list"
	"time"
)

type PacketType string

var (
	TABLE_SIZE        uint64 = 10240000
	PACKET_TABLE_SIZE uint   = 20000
)

const (
	NGAPType PacketType = "NGAP"
)

type Flow struct {
	info *FlowInfo
	next *Flow
}

type FlowInfo struct {
	RAN_UE_NGAP_ID int64
	FlowID         uint64
	FlowType       PacketType
	TotalNum       uint
	PacketList     list.List

	BeginTime       time.Time
	EndTime         time.Time
	EndTimeUs       int64
	VerificationTag uint32
	SrcIP           string
	DstIP           string
	TimeID          int64
	TaskID          string
}
type Packet struct {
	FlowID            uint64
	NgapType          string
	NgapProcedureCode string
	NgapRoute         string
	RAN_UE_NGAP_ID    int64
	ArriveTimeUs      int64
	ArriveTime        time.Time
	PacketLen         uint
	TimeInterval      int64
	NgapPayloadBytes  []byte
	PayloadBytes      []byte
	VerificationTag   uint32
	SrcIP             string
	DstIP             string
	TimeID            int64
	DirSeq            int
}

func loadFlow(flowId uint64, flowTable []*Flow) (*FlowInfo, bool) {
	cur := flowTable[flowId%TABLE_SIZE]
	if cur == nil {
		return nil, false
	}
	for ; cur != nil; cur = cur.next {
		if cur.info.FlowID == flowId {
			if cur.info.TotalNum == PACKET_TABLE_SIZE {
				return cur.info, false
			}
			return cur.info, true
		}
	}
	return nil, false
}

func storeFlow(flowId uint64, flowInfo *FlowInfo, flowTable []*Flow) {
	cur := flowTable[flowId%TABLE_SIZE]
	if cur == nil {
		flowTable[flowId%TABLE_SIZE] = &Flow{info: flowInfo, next: nil}
	} else {
		next := cur
		for next.next != cur {
			next = next.next
		}
		next.next = &Flow{info: flowInfo, next: nil}
	}
}

func deleteFlow(flowId uint64, flowTable []*Flow) bool {
	pre := flowTable[flowId%TABLE_SIZE]
	if pre.info.FlowID == flowId {
		flowTable[flowId%TABLE_SIZE] = pre.next
		return true
	}
	cur := pre.next
	for cur != nil && cur.info.FlowID != flowId {
		pre = pre.next
		cur = cur.next
	}
	if cur == nil {
		return false
	}
	pre.next = cur.next
	return true
}
