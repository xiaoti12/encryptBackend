package NGAP

import (
	"github.com/Freddy/sctp_flowmap/flowMap"

	//"github.com/free5gc/ngap"
	"github.com/free5gc/ngap/ngapType"
	//"github.com/google/gopacket"
	//"github.com/google/gopacket/pcap"
)

func NGSetupResponseDecoder(NGSetup ngapType.SuccessfulOutcomeValue,packet *flowMap.Packet) {
	//fmt.Println("===============NGSetupResponseDecoder begins working.==============")
	//fmt.Println("===============Nothing Need in the NGSetupResponse.===============")
}
func RANConfigurationUpdateAcknowledgeDecoder() {
	//fmt.Println("===============RANConfigurationUpdateAcknowledgeDecoder begins working.===============")
	//fmt.Println("===============Nothing Need in the RANConfigurationUpdateAcknowledge.===============")
}

func InitialContextSetupResDecoder(InitialContextSetupRes ngapType.SuccessfulOutcomeValue,packet *flowMap.Packet) {
	initiatingContextSetupRes := InitialContextSetupRes.InitialContextSetupResponse
	if initiatingContextSetupRes != nil {
		for i := 0; i < len(initiatingContextSetupRes.ProtocolIEs.List); i++ {
			ie := initiatingContextSetupRes.ProtocolIEs.List[i]
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

func PDUSessionResourceSetupResDecoder(PDUSessionResourceSetupRes ngapType.SuccessfulOutcomeValue,packet *flowMap.Packet) {
	pDUSessionResourceSetupRes := PDUSessionResourceSetupRes.PDUSessionResourceSetupResponse
	if pDUSessionResourceSetupRes != nil {
		for i := 0; i < len(pDUSessionResourceSetupRes.ProtocolIEs.List); i++ {
			ie := pDUSessionResourceSetupRes.ProtocolIEs.List[i]
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
func PDUSessionResourceReleaseDecoder(PDUSessionResourceRelease ngapType.SuccessfulOutcomeValue,packet *flowMap.Packet)  {
	pDUSessionResourceRelease := PDUSessionResourceRelease.PDUSessionResourceReleaseResponse
	if pDUSessionResourceRelease != nil {
		for i := 0; i < len(pDUSessionResourceRelease.ProtocolIEs.List); i++ {
			ie := pDUSessionResourceRelease.ProtocolIEs.List[i]
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

func PDUSessionResourceModifySucDecoder(PDUSessionResourceModify ngapType.SuccessfulOutcomeValue,packet *flowMap.Packet)  {
	pDUSessionResourceModify := PDUSessionResourceModify.PDUSessionResourceModifyResponse
	if pDUSessionResourceModify != nil {
		for i := 0; i < len(pDUSessionResourceModify.ProtocolIEs.List); i++ {
			ie := pDUSessionResourceModify.ProtocolIEs.List[i]
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


func PDUSessionResourceModifyConfirm(PDUSessionResourceModifyConfirm ngapType.SuccessfulOutcomeValue,packet *flowMap.Packet) {
	//fmt.Println("===============UplinkNASTransportDecoder begins working.===============")
	pDUSessionResourceModifyConfirm := PDUSessionResourceModifyConfirm.PDUSessionResourceModifyConfirm
	if pDUSessionResourceModifyConfirm != nil {
		for i := 0; i < len(pDUSessionResourceModifyConfirm.ProtocolIEs.List); i++ {
			ie := pDUSessionResourceModifyConfirm.ProtocolIEs.List[i]
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


func UEContextReleaseComplete(UEContextReleaseComplete ngapType.SuccessfulOutcomeValue,packet *flowMap.Packet) {
	//fmt.Println("===============UplinkNASTransportDecoder begins working.===============")
	uEContextReleaseComplete := UEContextReleaseComplete.UEContextReleaseComplete
	if uEContextReleaseComplete != nil {
		for i := 0; i < len(uEContextReleaseComplete.ProtocolIEs.List); i++ {
			ie := uEContextReleaseComplete.ProtocolIEs.List[i]
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



func UEContextModificationResponse(UEContextModificationResponse ngapType.SuccessfulOutcomeValue,packet *flowMap.Packet) {
	//fmt.Println("===============UplinkNASTransportDecoder begins working.===============")
	uEContextModificationResponse := UEContextModificationResponse.UEContextModificationResponse
	if uEContextModificationResponse != nil {
		for i := 0; i < len(uEContextModificationResponse.ProtocolIEs.List); i++ {
			ie := uEContextModificationResponse.ProtocolIEs.List[i]
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


func HandoverCommand(HandoverCommand ngapType.SuccessfulOutcomeValue,packet *flowMap.Packet) {
	//fmt.Println("===============UplinkNASTransportDecoder begins working.===============")
	handoverCommand := HandoverCommand.HandoverCommand
	if handoverCommand != nil {
		for i := 0; i < len(handoverCommand.ProtocolIEs.List); i++ {
			ie := handoverCommand.ProtocolIEs.List[i]
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

func HandoverRequestAcknowledge(HandoverRequestAcknowledge ngapType.SuccessfulOutcomeValue,packet *flowMap.Packet) {
	//fmt.Println("===============UplinkNASTransportDecoder begins working.===============")
	handoverRequestAcknowledge := HandoverRequestAcknowledge.HandoverRequestAcknowledge
	if handoverRequestAcknowledge != nil {
		for i := 0; i < len(handoverRequestAcknowledge.ProtocolIEs.List); i++ {
			ie := handoverRequestAcknowledge.ProtocolIEs.List[i]
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


func PathSwitchRequestAcknowledge(PathSwitchRequestAcknowledge ngapType.SuccessfulOutcomeValue,packet *flowMap.Packet) {
	//fmt.Println("===============UplinkNASTransportDecoder begins working.===============")
	pathSwitchRequestAcknowledge := PathSwitchRequestAcknowledge.PathSwitchRequestAcknowledge
	if pathSwitchRequestAcknowledge != nil {
		for i := 0; i < len(pathSwitchRequestAcknowledge.ProtocolIEs.List); i++ {
			ie := pathSwitchRequestAcknowledge.ProtocolIEs.List[i]
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

func HandoverCancelAcknowledge(HandoverCancelAcknowledge ngapType.SuccessfulOutcomeValue,packet *flowMap.Packet) {
	//fmt.Println("===============UplinkNASTransportDecoder begins working.===============")
	handoverCancelAcknowledge := HandoverCancelAcknowledge.HandoverCancelAcknowledge
	if handoverCancelAcknowledge != nil {
		for i := 0; i < len(handoverCancelAcknowledge.ProtocolIEs.List); i++ {
			ie := handoverCancelAcknowledge.ProtocolIEs.List[i]
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

func UERadioCapabilityCheckResponse(UERadioCapabilityCheckResponse ngapType.SuccessfulOutcomeValue,packet *flowMap.Packet) {
	//fmt.Println("===============UplinkNASTransportDecoder begins working.===============")
	uERadioCapabilityCheckResponse := UERadioCapabilityCheckResponse.UERadioCapabilityCheckResponse
	if uERadioCapabilityCheckResponse != nil {
		for i := 0; i < len(uERadioCapabilityCheckResponse.ProtocolIEs.List); i++ {
			ie := uERadioCapabilityCheckResponse.ProtocolIEs.List[i]
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






