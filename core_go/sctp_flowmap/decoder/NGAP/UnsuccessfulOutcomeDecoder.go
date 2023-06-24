package NGAP

import (
	"github.com/Freddy/sctp_flowmap/flowMap"
	"github.com/free5gc/ngap/ngapType"
)



func InitialContextSetupFailure(InitialContextSetupFailure ngapType.UnsuccessfulOutcomeValue,packet *flowMap.Packet) {
	initialContextSetupFailure := InitialContextSetupFailure.InitialContextSetupFailure
	if initialContextSetupFailure != nil {
		for i := 0; i < len(initialContextSetupFailure.ProtocolIEs.List); i++ {
			ie := initialContextSetupFailure.ProtocolIEs.List[i]
			switch ie.Id.Value {
			case ngapType.ProtocolIEIDAMFUENGAPID:
				//AMF_UE_NGAP_ID := ie.Value.AMFUENGAPID.Value
				//fmt.Println("[AMF-UE-NGAP-ID]:", AMF_UE_NGAP_ID)
			case ngapType.ProtocolIEIDRANUENGAPID:
				RAN_UE_NGAP_ID := ie.Value.RANUENGAPID.Value
				//fmt.Println("[RAN-UE-NGAP-ID]:", RAN_UE_NGAP_ID)
				packet.RAN_UE_NGAP_ID=RAN_UE_NGAP_ID
			}
		}
	}
}


func UEContextModificationFailure(UEContextModificationFailure ngapType.UnsuccessfulOutcomeValue,packet *flowMap.Packet) {
	uEContextModificationFailure := UEContextModificationFailure.UEContextModificationFailure
	if uEContextModificationFailure != nil {
		for i := 0; i < len(uEContextModificationFailure.ProtocolIEs.List); i++ {
			ie := uEContextModificationFailure.ProtocolIEs.List[i]
			switch ie.Id.Value {
			case ngapType.ProtocolIEIDAMFUENGAPID:
				//AMF_UE_NGAP_ID := ie.Value.AMFUENGAPID.Value
				//fmt.Println("[AMF-UE-NGAP-ID]:", AMF_UE_NGAP_ID)
			case ngapType.ProtocolIEIDRANUENGAPID:
				RAN_UE_NGAP_ID := ie.Value.RANUENGAPID.Value
				//fmt.Println("[RAN-UE-NGAP-ID]:", RAN_UE_NGAP_ID)
				packet.RAN_UE_NGAP_ID=RAN_UE_NGAP_ID
			}
		}
	}
}


func HandoverPreparationFailure(HandoverPreparationFailure ngapType.UnsuccessfulOutcomeValue,packet *flowMap.Packet) {
	handoverPreparationFailure := HandoverPreparationFailure.HandoverPreparationFailure
	if handoverPreparationFailure != nil {
		for i := 0; i < len(handoverPreparationFailure.ProtocolIEs.List); i++ {
			ie := handoverPreparationFailure.ProtocolIEs.List[i]
			switch ie.Id.Value {
			case ngapType.ProtocolIEIDAMFUENGAPID:
				//AMF_UE_NGAP_ID := ie.Value.AMFUENGAPID.Value
				//fmt.Println("[AMF-UE-NGAP-ID]:", AMF_UE_NGAP_ID)
			case ngapType.ProtocolIEIDRANUENGAPID:
				RAN_UE_NGAP_ID := ie.Value.RANUENGAPID.Value
				//fmt.Println("[RAN-UE-NGAP-ID]:", RAN_UE_NGAP_ID)
				packet.RAN_UE_NGAP_ID=RAN_UE_NGAP_ID
			}
		}
	}
}


func PathSwitchRequestFailure(PathSwitchRequestFailure ngapType.UnsuccessfulOutcomeValue,packet *flowMap.Packet) {
	pathSwitchRequestFailure := PathSwitchRequestFailure.PathSwitchRequestFailure
	if pathSwitchRequestFailure != nil {
		for i := 0; i < len(pathSwitchRequestFailure.ProtocolIEs.List); i++ {
			ie := pathSwitchRequestFailure.ProtocolIEs.List[i]
			switch ie.Id.Value {
			case ngapType.ProtocolIEIDAMFUENGAPID:
				//AMF_UE_NGAP_ID := ie.Value.AMFUENGAPID.Value
				//fmt.Println("[AMF-UE-NGAP-ID]:", AMF_UE_NGAP_ID)
			case ngapType.ProtocolIEIDRANUENGAPID:
				RAN_UE_NGAP_ID := ie.Value.RANUENGAPID.Value
				//fmt.Println("[RAN-UE-NGAP-ID]:", RAN_UE_NGAP_ID)
				packet.RAN_UE_NGAP_ID=RAN_UE_NGAP_ID
			}
		}
	}
}





