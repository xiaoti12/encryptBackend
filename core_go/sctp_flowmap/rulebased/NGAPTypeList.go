package rulebased

import (
	"container/list"
	"github.com/Freddy/sctp_flowmap/flowMap"
)


func ListGenerate(flowList *list.List) *list.List {
	TP := list.New()

	for flow := flowList.Front(); flow != nil;flow = flow.Next(){
		flowInfo := flow.Value.(*flowMap.FlowInfo)
		NGAP_List := list.New()
		ProcedureGenerate(flowInfo,NGAP_List)

		TP_next := &TypeList{
			FlowID:    flowInfo.FlowID,
			Flow:      flowInfo,
			NGAP_List: NGAP_List,
		}
		TP.PushBack(TP_next)
	}
	return TP
}

func ProcedureGenerate(flow *flowMap.FlowInfo,NGAP_list *list.List)  {
	for cur := flow.PacketList.Front(); cur != nil; cur = cur.Next(){
		packet := cur.Value.(*flowMap.Packet)
		message := dic_status[packet.NgapRoute]
		pro := dict_pro[packet.NgapProcedureCode]
		switch pro {
		case 3:
			NGAP_NODE :=&NGAPProcedureInfo{
			Procedure: packet.NgapProcedureCode,
			Status:    Success,
			}
			NGAP_list.PushBack(NGAP_NODE)
		case 4:
			NGAP_NODE :=&NGAPProcedureInfo{
				Procedure: packet.NgapProcedureCode,
				Status:    Failure,
			}
			NGAP_list.PushBack(NGAP_NODE)
		case 1:
			if message == 0 {
				NGAP_NODE :=&NGAPProcedureInfo{
					Procedure: packet.NgapProcedureCode,
					Status:    NoResponse,
				}
				NGAP_list.PushBack(NGAP_NODE)
			} else {
				for info:=NGAP_list.Front() ;info != nil; info = info.Next(){
					if info.Value.(*NGAPProcedureInfo).Procedure == packet.NgapProcedureCode &&
						info.Value.(*NGAPProcedureInfo).Status == NoResponse{
						info.Value.(*NGAPProcedureInfo).Status = Success
						break
					}
				}
			}
		case 2:
			if message == 0 {
				NGAP_NODE :=&NGAPProcedureInfo{
					Procedure: packet.NgapProcedureCode,
					Status:    NoResponse,
				}
				NGAP_list.PushBack(NGAP_NODE)
			}else if message == 1 {
				for info:=NGAP_list.Front() ;info != nil; info = info.Next(){
					if info.Value.(*NGAPProcedureInfo).Procedure == packet.NgapProcedureCode &&
						info.Value.(*NGAPProcedureInfo).Status == NoResponse{
						info.Value.(*NGAPProcedureInfo).Status = Success
						break
					}
				}
			}else {
				for info:=NGAP_list.Front() ;info != nil; info = info.Next(){
					if info.Value.(*NGAPProcedureInfo).Procedure == packet.NgapProcedureCode &&
						info.Value.(*NGAPProcedureInfo).Status == NoResponse{
						info.Value.(*NGAPProcedureInfo).Status = Failure
						break
					}
				}
			}
		}
	}

}


