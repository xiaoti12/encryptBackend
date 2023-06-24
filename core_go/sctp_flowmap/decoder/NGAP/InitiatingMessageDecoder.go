package NGAP

import (
	"github.com/Freddy/sctp_flowmap/flowMap"
	"github.com/free5gc/ngap/ngapType"
	"github.com/free5gc/openapi/models"
)

type SupportedTAI struct {
	Tai        models.Tai
	SNssaiList []models.Snssai
}


func NGSetupRequestDecoder(NGSetup ngapType.InitiatingMessageValue,packet *flowMap.Packet) {

}

func InitialContextSetupDecoder(InitialContextSetup ngapType.InitiatingMessageValue,packet *flowMap.Packet) {
	initiatialContextSetup := InitialContextSetup.InitialContextSetupRequest
	//notice the Setup here, there's no need to distinguish the Req and Res, because Res is in the SuccessfulOutcome.
	if initiatialContextSetup != nil {
		for i := 0; i < len(initiatialContextSetup.ProtocolIEs.List); i++ {
			ie := initiatialContextSetup.ProtocolIEs.List[i]
			switch ie.Id.Value {
			case ngapType.ProtocolIEIDAMFUENGAPID:
				//AMF_UE_NGAP_ID := ie.Value.AMFUENGAPID
				//fmt.Println("[AMF-UE-NGAP-ID]:", AMF_UE_NGAP_ID)
			case ngapType.ProtocolIEIDRANUENGAPID:
				RAN_UE_NGAP_ID := ie.Value.RANUENGAPID.Value
				//fmt.Println("[RAN-UE-NGAP-ID]:", RAN_UE_NGAP_ID)
				packet.RAN_UE_NGAP_ID=RAN_UE_NGAP_ID
			}
		}
	}
}



func InitialUEMessageDecoder(InitialUEMessage ngapType.InitiatingMessageValue,packet *flowMap.Packet) {
	//fmt.Println("===============InitialUEMessageDecoder begins working.===============")
	initialUEMessage := InitialUEMessage.InitialUEMessage
	if initialUEMessage != nil {
		for i := 0; i < len(initialUEMessage.ProtocolIEs.List); i++ {
			ie := initialUEMessage.ProtocolIEs.List[i]
			switch ie.Id.Value {
			case ngapType.ProtocolIEIDRANUENGAPID:
				id_RAN_UE_NGAP_ID := ie.Value.RANUENGAPID.Value
				//fmt.Println("[id-RAN-UE-NGAP-ID]:", id_RAN_UE_NGAP_ID)
				packet.RAN_UE_NGAP_ID=id_RAN_UE_NGAP_ID
			}
		}
	}
}

func DownlinkNASTransportDecoder(DownlinkNASTransport ngapType.InitiatingMessageValue,packet *flowMap.Packet) {
	//fmt.Println("===============DownlinkNASTransportDecoder begins working.===============")
	downlinkNASTransport := DownlinkNASTransport.DownlinkNASTransport
	if downlinkNASTransport != nil {
		for i := 0; i < len(downlinkNASTransport.ProtocolIEs.List); i++ {
			ie := downlinkNASTransport.ProtocolIEs.List[i]
			switch ie.Id.Value {
			case ngapType.ProtocolIEIDAMFUENGAPID:
				//id_AMF_UE_NGAP_ID := ie.Value.AMFUENGAPID.Value
				//fmt.Println("[id-AMF-UE-NGAP-ID]:", id_AMF_UE_NGAP_ID)
			case ngapType.ProtocolIEIDRANUENGAPID:
				id_RAN_UE_NGAP_ID := ie.Value.RANUENGAPID.Value
				//fmt.Println("[id-RAN-UE-NGAP-ID]:", id_RAN_UE_NGAP_ID)
				packet.RAN_UE_NGAP_ID=id_RAN_UE_NGAP_ID
			}

		}
	}
}





func PDUSessionResourceModifyDecoder(PDUSessionResourceModify ngapType.InitiatingMessageValue,packet *flowMap.Packet) {
	pDUSessionResourceModify := PDUSessionResourceModify.PDUSessionResourceModifyRequest
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



func PDUSessionResourcedReleaseDecoder(PDUSessionResourcedRelease ngapType.InitiatingMessageValue,packet *flowMap.Packet) {
	pduSessionResourcedRelease := PDUSessionResourcedRelease.PDUSessionResourceReleaseCommand
	if pduSessionResourcedRelease != nil {
		for i := 0; i < len(pduSessionResourcedRelease.ProtocolIEs.List); i++ {
			ie := pduSessionResourcedRelease.ProtocolIEs.List[i]
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

func PDUSessionResourceSetupDecoder(PDUSessionResourceSetup ngapType.InitiatingMessageValue,packet *flowMap.Packet) {
	pDUSeesionReusourceSetup := PDUSessionResourceSetup.PDUSessionResourceSetupRequest
	if pDUSeesionReusourceSetup != nil {
		for i := 0; i < len(pDUSeesionReusourceSetup.ProtocolIEs.List); i++ {
			ie := pDUSeesionReusourceSetup.ProtocolIEs.List[i]
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

func RANConfigurationUpdateDecoder(RANConfigurationUpdate ngapType.InitiatingMessageValue,packet *flowMap.Packet) {
	//fmt.Println("===============RANConfigurationUpdateDecoder begins working.===============")
	rANConfigurationUpdate := RANConfigurationUpdate.RANConfigurationUpdate
	if rANConfigurationUpdate != nil {
		for i := 0; i < len(rANConfigurationUpdate.ProtocolIEs.List); i++ {
			ie := rANConfigurationUpdate.ProtocolIEs.List[i]
			switch ie.Id.Value {
			case ngapType.ProtocolIEIDGlobalRANNodeID:
				//GlobalRANNodeID := ie.Value.GlobalRANNodeID.GlobalGNBID.GNBID.GNBID.Bytes
				//fmt.Println("[GlobalRANNodeID]:", GlobalRANNodeID)

			}
		}
	}
}




func UplinkNASTransportDecoder(UplinkNASTransport ngapType.InitiatingMessageValue,packet *flowMap.Packet) {
	//fmt.Println("===============UplinkNASTransportDecoder begins working.===============")
	uplinkNASTransport := UplinkNASTransport.UplinkNASTransport
	if uplinkNASTransport != nil {
		for i := 0; i < len(uplinkNASTransport.ProtocolIEs.List); i++ {
			ie := uplinkNASTransport.ProtocolIEs.List[i]
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


func PDUSessionResourceNotifyDecoder(PDUSessionResourceNotify ngapType.InitiatingMessageValue,packet *flowMap.Packet) {
	//fmt.Println("===============UplinkNASTransportDecoder begins working.===============")
	pDUSessionResourceNotify := PDUSessionResourceNotify.PDUSessionResourceNotify
	if pDUSessionResourceNotify != nil {
		for i := 0; i < len(pDUSessionResourceNotify.ProtocolIEs.List); i++ {
			ie := pDUSessionResourceNotify.ProtocolIEs.List[i]
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


func PDUSessionResourceModifyIndication(PDUSessionResourceModifyIndication ngapType.InitiatingMessageValue,packet *flowMap.Packet) {
	//fmt.Println("===============UplinkNASTransportDecoder begins working.===============")
	pDUSessionResourceModifyIndication := PDUSessionResourceModifyIndication.PDUSessionResourceModifyIndication
	if pDUSessionResourceModifyIndication != nil {
		for i := 0; i < len(pDUSessionResourceModifyIndication.ProtocolIEs.List); i++ {
			ie := pDUSessionResourceModifyIndication.ProtocolIEs.List[i]
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


func UEContextReleaseRequest(UEContextReleaseRequest ngapType.InitiatingMessageValue,packet *flowMap.Packet) {
	//fmt.Println("===============UplinkNASTransportDecoder begins working.===============")
	uEContextReleaseRequest := UEContextReleaseRequest.UEContextReleaseRequest
	if uEContextReleaseRequest != nil {
		for i := 0; i < len(uEContextReleaseRequest.ProtocolIEs.List); i++ {
			ie := uEContextReleaseRequest.ProtocolIEs.List[i]
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


func UEContextModificationRequest(UEContextModificationRequest ngapType.InitiatingMessageValue,packet *flowMap.Packet) {
	//fmt.Println("===============UplinkNASTransportDecoder begins working.===============")
	uEContextModificationRequest := UEContextModificationRequest.UEContextModificationRequest
	if uEContextModificationRequest != nil {
		for i := 0; i < len(uEContextModificationRequest.ProtocolIEs.List); i++ {
			ie := uEContextModificationRequest.ProtocolIEs.List[i]
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

func RRCInactiveTransitionReport(RRCInactiveTransitionReport ngapType.InitiatingMessageValue,packet *flowMap.Packet) {
	//fmt.Println("===============UplinkNASTransportDecoder begins working.===============")
	rRCInactiveTransitionReport := RRCInactiveTransitionReport.RRCInactiveTransitionReport
	if rRCInactiveTransitionReport != nil {
		for i := 0; i < len(rRCInactiveTransitionReport.ProtocolIEs.List); i++ {
			ie := rRCInactiveTransitionReport.ProtocolIEs.List[i]
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


func HandoverRequired(HandoverRequired ngapType.InitiatingMessageValue,packet *flowMap.Packet) {
	//fmt.Println("===============UplinkNASTransportDecoder begins working.===============")
	handoverRequired := HandoverRequired.HandoverRequired
	if handoverRequired != nil {
		for i := 0; i < len(handoverRequired.ProtocolIEs.List); i++ {
			ie := handoverRequired.ProtocolIEs.List[i]
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


func HandoverNotify(HandoverNotify ngapType.InitiatingMessageValue,packet *flowMap.Packet) {
	//fmt.Println("===============UplinkNASTransportDecoder begins working.===============")
	handoverNotify := HandoverNotify.HandoverNotify
	if handoverNotify != nil {
		for i := 0; i < len(handoverNotify.ProtocolIEs.List); i++ {
			ie := handoverNotify.ProtocolIEs.List[i]
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


func PathSwitchRequest(PathSwitchRequest ngapType.InitiatingMessageValue,packet *flowMap.Packet) {
	//fmt.Println("===============UplinkNASTransportDecoder begins working.===============")
	pathSwitchRequest := PathSwitchRequest.PathSwitchRequest
	if pathSwitchRequest != nil {
		for i := 0; i < len(pathSwitchRequest.ProtocolIEs.List); i++ {
			ie := pathSwitchRequest.ProtocolIEs.List[i]
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


func HandoverCancel(HandoverCancel ngapType.InitiatingMessageValue,packet *flowMap.Packet) {
	//fmt.Println("===============UplinkNASTransportDecoder begins working.===============")
	handoverCancel := HandoverCancel.HandoverCancel
	if handoverCancel != nil {
		for i := 0; i < len(handoverCancel.ProtocolIEs.List); i++ {
			ie := handoverCancel.ProtocolIEs.List[i]
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


func UERadioCapabilityCheckRequest(UERadioCapabilityCheckRequest ngapType.InitiatingMessageValue,packet *flowMap.Packet) {
	//fmt.Println("===============UplinkNASTransportDecoder begins working.===============")
	uERadioCapabilityCheckRequest := UERadioCapabilityCheckRequest.UERadioCapabilityCheckRequest
	if uERadioCapabilityCheckRequest != nil {
		for i := 0; i < len(uERadioCapabilityCheckRequest.ProtocolIEs.List); i++ {
			ie := uERadioCapabilityCheckRequest.ProtocolIEs.List[i]
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


func UplinkRANStatusTransfer(UplinkRANStatusTransfer ngapType.InitiatingMessageValue,packet *flowMap.Packet) {
	//fmt.Println("===============UplinkNASTransportDecoder begins working.===============")
	uplinkRANStatusTransfer := UplinkRANStatusTransfer.UplinkRANStatusTransfer
	if uplinkRANStatusTransfer != nil {
		for i := 0; i < len(uplinkRANStatusTransfer.ProtocolIEs.List); i++ {
			ie := uplinkRANStatusTransfer.ProtocolIEs.List[i]
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

func DownlinkRANStatusTransfer(DownlinkRANStatusTransfer ngapType.InitiatingMessageValue,packet *flowMap.Packet) {
	//fmt.Println("===============UplinkNASTransportDecoder begins working.===============")
	downlinkRANStatusTransfer := DownlinkRANStatusTransfer.DownlinkRANStatusTransfer
	if downlinkRANStatusTransfer != nil {
		for i := 0; i < len(downlinkRANStatusTransfer.ProtocolIEs.List); i++ {
			ie := downlinkRANStatusTransfer.ProtocolIEs.List[i]
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

func NASNonDeliveryIndication(NASNonDeliveryIndication ngapType.InitiatingMessageValue,packet *flowMap.Packet) {
	//fmt.Println("===============UplinkNASTransportDecoder begins working.===============")
	nASNonDeliveryIndication := NASNonDeliveryIndication.NASNonDeliveryIndication
	if nASNonDeliveryIndication != nil {
		for i := 0; i < len(nASNonDeliveryIndication.ProtocolIEs.List); i++ {
			ie := nASNonDeliveryIndication.ProtocolIEs.List[i]
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


func RerouteNASRequest(RerouteNASRequest ngapType.InitiatingMessageValue,packet *flowMap.Packet) {
	//fmt.Println("===============UplinkNASTransportDecoder begins working.===============")
	rerouteNASRequest := RerouteNASRequest.RerouteNASRequest
	if rerouteNASRequest != nil {
		for i := 0; i < len(rerouteNASRequest.ProtocolIEs.List); i++ {
			ie := rerouteNASRequest.ProtocolIEs.List[i]
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

func ErrorIndication(ErrorIndication ngapType.InitiatingMessageValue,packet *flowMap.Packet) {
	//fmt.Println("===============UplinkNASTransportDecoder begins working.===============")
	errorIndication := ErrorIndication.ErrorIndication
	if errorIndication != nil {
		for i := 0; i < len(errorIndication.ProtocolIEs.List); i++ {
			ie := errorIndication.ProtocolIEs.List[i]
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

func DownlinkUEAssociatedNRPPaTransport(DownlinkUEAssociatedNRPPaTransport ngapType.InitiatingMessageValue,packet *flowMap.Packet) {
	//fmt.Println("===============UplinkNASTransportDecoder begins working.===============")
	downlinkUEAssociatedNRPPaTransport := DownlinkUEAssociatedNRPPaTransport.DownlinkUEAssociatedNRPPaTransport
	if downlinkUEAssociatedNRPPaTransport != nil {
		for i := 0; i < len(downlinkUEAssociatedNRPPaTransport.ProtocolIEs.List); i++ {
			ie := downlinkUEAssociatedNRPPaTransport.ProtocolIEs.List[i]
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

func UplinkUEAssociatedNRPPaTransport(UplinkUEAssociatedNRPPaTransport ngapType.InitiatingMessageValue,packet *flowMap.Packet) {
	//fmt.Println("===============UplinkNASTransportDecoder begins working.===============")
	uplinkUEAssociatedNRPPaTransport := UplinkUEAssociatedNRPPaTransport.UplinkUEAssociatedNRPPaTransport
	if uplinkUEAssociatedNRPPaTransport != nil {
		for i := 0; i < len(uplinkUEAssociatedNRPPaTransport.ProtocolIEs.List); i++ {
			ie := uplinkUEAssociatedNRPPaTransport.ProtocolIEs.List[i]
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


func TraceStart(TraceStart ngapType.InitiatingMessageValue,packet *flowMap.Packet) {
	//fmt.Println("===============UplinkNASTransportDecoder begins working.===============")
	traceStart := TraceStart.TraceStart
	if traceStart != nil {
		for i := 0; i < len(traceStart.ProtocolIEs.List); i++ {
			ie := traceStart.ProtocolIEs.List[i]
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

func TraceFailureIndication(TraceFailureIndication ngapType.InitiatingMessageValue,packet *flowMap.Packet) {
	//fmt.Println("===============UplinkNASTransportDecoder begins working.===============")
	traceFailureIndication := TraceFailureIndication.TraceFailureIndication
	if traceFailureIndication != nil {
		for i := 0; i < len(traceFailureIndication.ProtocolIEs.List); i++ {
			ie := traceFailureIndication.ProtocolIEs.List[i]
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

func DeactivateTrace(DeactivateTrace ngapType.InitiatingMessageValue,packet *flowMap.Packet) {
	//fmt.Println("===============UplinkNASTransportDecoder begins working.===============")
	deactivateTrace := DeactivateTrace.DeactivateTrace
	if deactivateTrace != nil {
		for i := 0; i < len(deactivateTrace.ProtocolIEs.List); i++ {
			ie := deactivateTrace.ProtocolIEs.List[i]
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


func LocationReportingControl(LocationReportingControl ngapType.InitiatingMessageValue,packet *flowMap.Packet) {
	//fmt.Println("===============UplinkNASTransportDecoder begins working.===============")
	locationReportingControl := LocationReportingControl.LocationReportingControl
	if locationReportingControl != nil {
		for i := 0; i < len(locationReportingControl.ProtocolIEs.List); i++ {
			ie := locationReportingControl.ProtocolIEs.List[i]
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

func LocationReportingFailureIndication(LocationReportingFailureIndication ngapType.InitiatingMessageValue,packet *flowMap.Packet) {
	//fmt.Println("===============UplinkNASTransportDecoder begins working.===============")
	locationReportingFailureIndication := LocationReportingFailureIndication.LocationReportingFailureIndication
	if locationReportingFailureIndication != nil {
		for i := 0; i < len(locationReportingFailureIndication.ProtocolIEs.List); i++ {
			ie := locationReportingFailureIndication.ProtocolIEs.List[i]
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

func LocationReport(LocationReport ngapType.InitiatingMessageValue,packet *flowMap.Packet) {
	//fmt.Println("===============UplinkNASTransportDecoder begins working.===============")
	locationReport := LocationReport.LocationReport
	if locationReport != nil {
		for i := 0; i < len(locationReport.ProtocolIEs.List); i++ {
			ie := locationReport.ProtocolIEs.List[i]
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

func UETNLABindingReleaseRequest(UETNLABindingReleaseRequest ngapType.InitiatingMessageValue,packet *flowMap.Packet) {
	//fmt.Println("===============UplinkNASTransportDecoder begins working.===============")
	uETNLABindingReleaseRequest := UETNLABindingReleaseRequest.UETNLABindingReleaseRequest
	if uETNLABindingReleaseRequest != nil {
		for i := 0; i < len(uETNLABindingReleaseRequest.ProtocolIEs.List); i++ {
			ie := uETNLABindingReleaseRequest.ProtocolIEs.List[i]
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


func UERadioCapabilityInfoIndication(UERadioCapabilityInfoIndication ngapType.InitiatingMessageValue,packet *flowMap.Packet) {
	//fmt.Println("===============UplinkNASTransportDecoder begins working.===============")
	uERadioCapabilityInfoIndication := UERadioCapabilityInfoIndication.UERadioCapabilityInfoIndication
	if uERadioCapabilityInfoIndication != nil {
		for i := 0; i < len(uERadioCapabilityInfoIndication.ProtocolIEs.List); i++ {
			ie := uERadioCapabilityInfoIndication.ProtocolIEs.List[i]
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



func SecondaryRATDataUsageReport(SecondaryRATDataUsageReport ngapType.InitiatingMessageValue,packet *flowMap.Packet) {
	//fmt.Println("===============UplinkNASTransportDecoder begins working.===============")
	secondaryRATDataUsageReport := SecondaryRATDataUsageReport.SecondaryRATDataUsageReport
	if secondaryRATDataUsageReport != nil {
		for i := 0; i < len(secondaryRATDataUsageReport.ProtocolIEs.List); i++ {
			ie := secondaryRATDataUsageReport.ProtocolIEs.List[i]
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









