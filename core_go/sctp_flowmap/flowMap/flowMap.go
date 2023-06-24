package flowMap

import (
	"container/list"
	"fmt"
	"github.com/Freddy/sctp_flowmap/database/PacketDB"
	"github.com/Freddy/sctp_flowmap/database/TimeFlow"
	"github.com/Freddy/sctp_flowmap/database/UEFlow"
)

var (
	FlowTable_UE   []*Flow
	FlowTable_Time []*Flow
	flowCount      uint64 = 0
	TimeMax        int64  = 300000000
	min_intervl    int64  = 1
	TimeChip       int64  = 20000000
)

func init() {
	FlowTable_UE = make([]*Flow, TABLE_SIZE)
	FlowTable_Time = make([]*Flow, TABLE_SIZE)
}

func Count_UE_ID(packet *Packet) (uint64, bool) {
	RAN_UE_NGAP_ID := packet.RAN_UE_NGAP_ID
	if RAN_UE_NGAP_ID == -1 {
		return 0, false
	}
	//flowID := FastFourHash(string(packet.SrcIP), string(packet.DstIP), string(RAN_UE_NGAP_ID),string(packet.VerificationTag))
	flowID := FastTwoHash([]byte(string(RAN_UE_NGAP_ID)), []byte(string(packet.VerificationTag)))
	//flowID := FnvHash([]byte(string(RAN_UE_NGAP_ID)))
	return flowID, true

}

func Count_Time_ID(packet *Packet, TimeFirst int64) (uint64, int64) {
	Time := packet.ArriveTimeUs - TimeFirst
	Time = Time / TimeChip
	return FastTwoHash([]byte(string(Time)), []byte(string(packet.VerificationTag))), Time
}

func Put(packet *Packet, flowTable []*Flow, flowID uint64, taskid string) bool {
	var flowInfo *FlowInfo
	var first = false // 是否流的首包
	flowInfo, isExist := loadFlow(flowID, flowTable)
	if isExist {
		packet.TimeInterval = packet.ArriveTimeUs - flowInfo.EndTimeUs + min_intervl
		flowInfo.EndTimeUs = packet.ArriveTimeUs
		flowInfo.PacketList.PushBack(packet)
		flowInfo.TotalNum = flowInfo.TotalNum + 1
		if flowInfo.SrcIP == packet.SrcIP {
			packet.DirSeq = 1
		} else {
			packet.DirSeq = -1
		}

	} else {
		// 首次接收，创建流info
		flowInfo = &FlowInfo{
			FlowID:          flowID,
			RAN_UE_NGAP_ID:  packet.RAN_UE_NGAP_ID,
			FlowType:        NGAPType,
			TotalNum:        1,
			VerificationTag: packet.VerificationTag,
			SrcIP:           packet.SrcIP,
			DstIP:           packet.DstIP,
			TimeID:          packet.TimeID,
			BeginTime:       packet.ArriveTime,
			TaskID:          taskid,
		}
		packet.DirSeq = 1
		packet.TimeInterval = min_intervl
		flowInfo.EndTime = packet.ArriveTime
		flowInfo.EndTimeUs = packet.ArriveTimeUs
		flowInfo.PacketList = list.List{}
		flowInfo.PacketList.PushBack(packet)
		storeFlow(flowID, flowInfo, flowTable)
		flowCount++
		first = true
	}
	return first
}

func ItFlowMap(flowTable []*Flow) {
	for _, flow := range flowTable {
		if flow == nil {
			continue
		}
		for cur := flow; cur != nil; cur = cur.next {
			fmt.Println("FlowID:", cur.info.FlowID)
			//fmt.Println("Totalnum",cur.info.TotalNum)
			//fmt.Println("TimeID",cur.info.TimeID)
			//fmt.Println("RAN_UE_ID:",cur.info.RAN_UE_NGAP_ID)
			//i:=cur.info.TotalNum
			fmt.Println("TotalNum:", cur.info.TotalNum)

			for pac := cur.info.PacketList.Front(); pac != nil; pac = pac.Next() {
				parse := pac.Value.(*Packet)
				//fmt.Println("packetnum:")
				//fmt.Println("RAN_UE_NGAP_ID",parse.RAN_UE_NGAP_ID)
				fmt.Println("NgapType:", parse.NgapType)
				//fmt.Println("NgapProc:",parse.NgapProcedureCode)

				//fmt.Println("packetlen:",parse.PacketLen)
				//fmt.Println("Arrivetime",parse.ArriveTime)
				//fmt.Println("timeinter",parse.TimeInterval)
			}

		}
	}

}

func UEFlowMapToStore() *list.List {
	var rubbishList = list.New()

	for _, flow := range FlowTable_UE {
		if flow == nil {
			continue
		}
		for cur := flow; cur != nil; cur = cur.next {
			flowInfo := cur.info
			rubbishList.PushBack(flowInfo)
			deleteFlow(flowInfo.FlowID, FlowTable_UE)
			flowCount--

		}
	}
	//fmt.Println("UEFlow")
	UEflowStore(rubbishList)
	return rubbishList
}

func UEflowStore(rubbishList *list.List) {
	var UeFlowList = list.New()
	for info := rubbishList.Front(); info != nil; info = info.Next() {
		flowInfo := info.Value.(*FlowInfo)
		fl := &UEFlow.UeFlow{
			FlowId:          uint64(flowInfo.FlowID),
			RanUeNgapId:     uint64(flowInfo.RAN_UE_NGAP_ID),
			TotalNum:        uint32(flowInfo.TotalNum),
			BeginTime:       flowInfo.BeginTime,
			LatestTime:      flowInfo.EndTime,
			VerificationTag: uint64(flowInfo.VerificationTag),
			SrcIP:           flowInfo.SrcIP,
			DstIP:           flowInfo.DstIP,
			//TimeID          uint64
			StatusFlow: 0,
			TaskID:     flowInfo.TaskID,
		}
		UeFlowList.PushBack(fl)

	}
	//fmt.Println("UEList")
	UEFlow.InsertUeFlow(UeFlowList)

}

func TimeFlowMapToStore() {
	var rubbishList = list.New()
	for _, flow := range FlowTable_Time {
		if flow == nil {
			continue
		}
		for cur := flow; cur != nil; cur = cur.next {
			flowInfo := cur.info
			rubbishList.PushBack(flowInfo)
			deleteFlow(flowInfo.FlowID, FlowTable_Time)
			flowCount--

		}
	}
	//fmt.Println("TimeFlow")
	TimeflowStore(rubbishList)
}

func TimeflowStore(rubbishList *list.List) {
	var TimeFlowList = list.New()
	var PacketList = list.New()
	for info := rubbishList.Front(); info != nil; info = info.Next() {
		flowInfo := info.Value.(*FlowInfo)
		fl := &TimeFlow.TimeFlow{
			FlowId:          uint64(flowInfo.FlowID),
			RanUeNgapId:     uint64(flowInfo.RAN_UE_NGAP_ID),
			TotalNum:        uint32(flowInfo.TotalNum),
			BeginTime:       flowInfo.BeginTime,
			LatestTime:      flowInfo.EndTime,
			VerificationTag: uint64(flowInfo.VerificationTag),
			SrcIP:           flowInfo.SrcIP,
			DstIP:           flowInfo.DstIP,
			//TimeID          uint64
			StatusFlow: 0,
			TaskID:     flowInfo.TaskID,
		}
		TimeFlowList.PushBack(fl)
		for cur := flowInfo.PacketList.Front(); cur != nil; cur = cur.Next() {
			parse := cur.Value.(*Packet)
			packet := &PacketDB.Packet{
				//PacketId: FnvHash([]byte(string(parse.ArriveTimeUs))),
				NgapType:          parse.NgapType,
				NgapProcedureCode: parse.NgapProcedureCode,
				RanUeNgapId:       uint64(parse.RAN_UE_NGAP_ID),
				PacketLen:         uint32(parse.PacketLen),
				ArriveTimeUs:      uint64(parse.ArriveTimeUs),
				ArriveTime:        parse.ArriveTime,
				TimeInterval:      uint64(parse.TimeInterval),
				VerificationTag:   uint64(parse.VerificationTag),
				SrcIP:             parse.SrcIP,
				DstIP:             parse.DstIP,
				DirSeq:            uint16(parse.DirSeq),
				FlowUEID:          uint64(parse.TimeID),
				FlowTimeID:        uint64(parse.FlowID),
				StatusPacket:      0,
			}
			PacketList.PushBack(packet)
		}
	}
	//fmt.Println("TimeList")
	TimeFlow.InsertTimeFlow(TimeFlowList)
	PacketDB.InsertPacket(PacketList)

}
