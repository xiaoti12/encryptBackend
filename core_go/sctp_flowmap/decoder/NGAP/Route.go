package NGAP

import (
	"github.com/Freddy/sctp_flowmap/flowMap"
	"github.com/free5gc/ngap/ngapType"
)

// 对NGAP包进行初步的分类

func RouteNGAP(DecResult *ngapType.NGAPPDU,packet *flowMap.Packet) {
	//fmt.Println("===============RouteNGAP begins working.===============")
	switch DecResult.Present {
	case ngapType.NGAPPDUPresentInitiatingMessage:
		if DecResult.InitiatingMessage == nil {
			//fmt.Println("===============There is no initiatingMessage in this package.===============")
		}
		packet.NgapRoute = "InitiatingMessage"
		RouteInitiatingMessage(DecResult,packet)


	case ngapType.NGAPPDUPresentSuccessfulOutcome:
		if DecResult.SuccessfulOutcome == nil {
			//fmt.Println("===============There is no SuccessfulOutcome in this package.===============")
		}
		packet.NgapRoute = "SuccessfulOutcome"
		RouteSuccessfulOutcome(DecResult,packet)

	case ngapType.NGAPPDUPresentUnsuccessfulOutcome:
		if DecResult.UnsuccessfulOutcome == nil {
			//fmt.Println("===============There is no UnsuccessfulOutcome in this package.============")
		}
		packet.NgapRoute = "UnsuccessfulOutcome"
		RouteUnsuccessfulOutcome(DecResult,packet)

		//case ngapType.NGAPPDUPresentNothing:
	}
}


func RouteInitiatingMessage(DecResult *ngapType.NGAPPDU,packet *flowMap.Packet) {
	InitiatingMessage := DecResult.InitiatingMessage
	switch InitiatingMessage.ProcedureCode.Value {
	case ngapType.ProcedureCodeNGSetup:
		//fmt.Println("==================== NGSetup ====================")
		packet.NgapProcedureCode="NGSetup"
		packet.NgapType="NGSetupRequest"
		NGSetupRequestDecoder(InitiatingMessage.Value,packet)
	case ngapType.ProcedureCodeInitialUEMessage:
		//fmt.Println("==================== InitialUEMessage ====================")
		packet.NgapProcedureCode="InitialUEMessage"
		packet.NgapType="InitialUEMessage"
		InitialUEMessageDecoder(InitiatingMessage.Value,packet)
	case ngapType.ProcedureCodeDownlinkNASTransport:
		//fmt.Println("==================== DownlinkNASTransport ====================")
		packet.NgapProcedureCode="DownlinkNASTransport"
		packet.NgapType="DownlinkNASTransport"
		DownlinkNASTransportDecoder(InitiatingMessage.Value,packet)
	case ngapType.ProcedureCodeAMFConfigurationUpdate:
		//fmt.Println("==================== AMFConfigurationUpdate ====================")
		packet.NgapProcedureCode="AMFConfigurationUpdate"
		packet.NgapType="AMFConfigurationUpdate"
	case ngapType.ProcedureCodeAMFStatusIndication:
		//fmt.Println("==================== AMFStatusIndication ====================")
		packet.NgapProcedureCode="AMFStatusIndication"
		packet.NgapType="AMFStatusIndication"
	case ngapType.ProcedureCodeDeactivateTrace:
		//fmt.Println("==================== DeactivateTrace ====================")
		packet.NgapProcedureCode="DeactivateTrace"
		packet.NgapType="DeactivateTrace"
		DeactivateTrace(InitiatingMessage.Value,packet)
	case ngapType.ProcedureCodeCellTrafficTrace:
		//fmt.Println("==================== CellTrafficTrace ====================")
		packet.NgapProcedureCode=""
	case ngapType.ProcedureCodeDownlinkNonUEAssociatedNRPPaTransport:
		//fmt.Println("==================== DownlinkNonUEAssociatedNRPPaTransport ====================")
		packet.NgapProcedureCode="DownlinkNonUEAssociatedNRPPaTransport"
		packet.NgapType="DownlinkNonUEAssociatedNRPPaTransport"
	case ngapType.ProcedureCodeDownlinkRANConfigurationTransfer:
		//fmt.Println("==================== DownlinkRANConfigurationTransfer ====================")
		packet.NgapProcedureCode="DownlinkRANConfigurationTransfer"
		packet.NgapType="DownlinkRANConfigurationTransfer"
	case ngapType.ProcedureCodeErrorIndication:
		//fmt.Println("==================== ErrorIndication ====================")
		packet.NgapProcedureCode="ErrorIndication"
		packet.NgapType="ErrorIndication"
		ErrorIndication(InitiatingMessage.Value,packet)
	case ngapType.ProcedureCodeDownlinkRANStatusTransfer:
		//fmt.Println("==================== DownLinkRANStatusTransfer ====================")
		packet.NgapProcedureCode="DownLinkRANStatusTransfer"
		packet.NgapType="DownLinkRANStatusTransfer"
		DownlinkRANStatusTransfer(InitiatingMessage.Value,packet)
	case ngapType.ProcedureCodeHandoverCancel:
		//fmt.Println("==================== HandoverCancel ====================")
		packet.NgapProcedureCode="HandoverCancel"
		packet.NgapType="HandoverCancel"
		HandoverCancel(InitiatingMessage.Value,packet)
	case ngapType.ProcedureCodeInitialContextSetup:
		//fmt.Println("==================== InitialContextSetup ====================")
		packet.NgapProcedureCode="InitialContextSetup"
		packet.NgapType="InitialContextSetupRequest"
		InitialContextSetupDecoder(InitiatingMessage.Value,packet)
	case ngapType.ProcedureCodePathSwitchRequest:
		//fmt.Println("==================== PathSwitchRequest ====================")
		packet.NgapProcedureCode="PathSwitchRequest"
		packet.NgapType="PathSwitchRequest"
		PathSwitchRequest(InitiatingMessage.Value,packet)
	case ngapType.ProcedureCodePDUSessionResourceModify:
		//fmt.Println("==================== PDUSessionResourceModify ====================")
		packet.NgapProcedureCode="PDUSessionResourceModify"
		packet.NgapType="PDUSessionResourceModifyRequest"
		PDUSessionResourceModifyDecoder(InitiatingMessage.Value,packet)

	case ngapType.ProcedureCodePDUSessionResourceModifyIndication:
		//fmt.Println("==================== PDUSessionResourceModifyIndication ====================")
		packet.NgapProcedureCode="PDUSessionResourceModifyIndication"
		packet.NgapType="PDUSessionResourceModifyIndication"
		PDUSessionResourceModifyIndication(InitiatingMessage.Value,packet)
	case ngapType.ProcedureCodePDUSessionResourceRelease:
		//fmt.Println("==================== PDUSessionResourceRelease ====================")
		packet.NgapProcedureCode="PDUSessionResourceRelease"
		packet.NgapType="PDUSessionResourceReleaseCommand"
		PDUSessionResourcedReleaseDecoder(InitiatingMessage.Value,packet)
	case ngapType.ProcedureCodePDUSessionResourceSetup:
		//fmt.Println("==================== PDUSessionResourceSetup ====================")
		packet.NgapProcedureCode="PDUSessionResourceSetup"
		packet.NgapType="PDUSessionResourceSetupRequest"
		PDUSessionResourceSetupDecoder(InitiatingMessage.Value,packet)
	case ngapType.ProcedureCodePWSCancel:
		//fmt.Println("==================== PWSCancel ====================")
		packet.NgapProcedureCode="PWSCancel"
		packet.NgapType="PWSCancelRequest"
	case ngapType.ProcedureCodeRANConfigurationUpdate:
		//fmt.Println("==================== RANConfigurationUpdate ====================")
		packet.NgapProcedureCode="RANConfigurationUpdate"
		packet.NgapType="RANConfigurationUpdate"
		RANConfigurationUpdateDecoder(InitiatingMessage.Value,packet)
	case ngapType.ProcedureCodeUEContextModification:
		//fmt.Println("==================== UEContextModification ====================")
		packet.NgapProcedureCode="UEContextModification"
		packet.NgapType="UEContextModificationRequest"
		UEContextModificationRequest(InitiatingMessage.Value,packet)
	case ngapType.ProcedureCodeUEContextReleaseRequest:
		//fmt.Println("==================== UEContextReleaseRequest ====================")
		packet.NgapProcedureCode="UEContextReleaseRequest"
		packet.NgapType="UEContextReleaseRequest"
		UEContextReleaseRequest(InitiatingMessage.Value,packet)
	case ngapType.ProcedureCodeUEContextRelease:
		//fmt.Println("==================== UEContextRelease ====================")
		packet.NgapProcedureCode="UEContextRelease"
		packet.NgapType="UEContextReleaseCommand"
	case ngapType.ProcedureCodeUERadioCapabilityCheck:
		//fmt.Println("==================== UERadioCapabilityCheck ====================")
		packet.NgapProcedureCode="UERadioCapabilityCheck"
		packet.NgapType="UERadioCapabilityCheckRequest"
		UERadioCapabilityCheckRequest(InitiatingMessage.Value,packet)
	case ngapType.ProcedureCodeDownlinkUEAssociatedNRPPaTransport:
		//fmt.Println("==================== DownlinkUEAssociatedNRPPaTransport ====================")
		packet.NgapProcedureCode="DownlinkUEAssociatedNRPPaTransport"
		packet.NgapType="DownlinkUEAssociatedNRPPaTransport"
		DownlinkUEAssociatedNRPPaTransport(InitiatingMessage.Value,packet)
	case ngapType.ProcedureCodeHandoverNotification:
		//fmt.Println("==================== HandoverNotification ====================")
		packet.NgapProcedureCode="HandoverNotification"
		packet.NgapType="HandoverNotification"
		HandoverNotify(InitiatingMessage.Value,packet)
	case ngapType.ProcedureCodeLocationReport:
		//fmt.Println("==================== LocationReport ====================")
		packet.NgapProcedureCode="LocationReport"
		packet.NgapType="LocationReport"
		LocationReport(InitiatingMessage.Value,packet)
	case ngapType.ProcedureCodeLocationReportingControl:
		//fmt.Println("==================== LocationReportingControl ====================")
		packet.NgapProcedureCode="LocationReportingControl"
		packet.NgapType="LocationReportingControl"
		LocationReportingControl(InitiatingMessage.Value,packet)
	case ngapType.ProcedureCodeLocationReportingFailureIndication:
		//fmt.Println("==================== LocationReportingFailureIndication ====================")
		packet.NgapProcedureCode="LocationReportingFailureIndication"
		packet.NgapType="LocationReportingFailureIndication"
		LocationReportingFailureIndication(InitiatingMessage.Value,packet)
	case ngapType.ProcedureCodeNASNonDeliveryIndication:
		//fmt.Println("==================== NASNonDeliveryIndication ====================")
		packet.NgapProcedureCode="NASNonDeliveryIndication"
		packet.NgapType="NASNonDeliveryIndication"
		NASNonDeliveryIndication(InitiatingMessage.Value,packet)
	case ngapType.ProcedureCodeOverloadStart:
		//fmt.Println("==================== OverloadStart ====================")
		packet.NgapProcedureCode="OverloadStart"
		packet.NgapType="OverloadStart"
	case ngapType.ProcedureCodeOverloadStop:
		//fmt.Println("==================== OverloadStop ====================")
		packet.NgapProcedureCode="OverloadStop"
		packet.NgapType="OverloadStop"
	case ngapType.ProcedureCodePaging:
		//fmt.Println("==================== Paging ====================")
		packet.NgapProcedureCode="Paging"
		packet.NgapType="Paging"
	case ngapType.ProcedureCodePDUSessionResourceNotify:
		//fmt.Println("==================== PDUSessionResourceNotify ====================")
		packet.NgapProcedureCode="PDUSessionResourceNotify"
		packet.NgapType="PDUSessionResourceNotify"
		PDUSessionResourceNotifyDecoder(InitiatingMessage.Value,packet)
	case ngapType.ProcedureCodePrivateMessage:
		//fmt.Println("==================== PrivateMessage ====================")
		packet.NgapProcedureCode="PrivateMessage"
		packet.NgapType="PrivateMessage"
	case ngapType.ProcedureCodePWSFailureIndication:
		//fmt.Println("==================== PWSFailureIndication ====================")
		packet.NgapProcedureCode="PWSFailureIndication"
		packet.NgapType="PWSFailureIndication"
	case ngapType.ProcedureCodePWSRestartIndication:
		//fmt.Println("==================== PWSRestartIndication ====================")
		packet.NgapProcedureCode="PWSRestartIndication"
		packet.NgapType="PWSRestartIndication"
	case ngapType.ProcedureCodeRerouteNASRequest:
		//fmt.Println("==================== RerouteNASRequest ====================")
		packet.NgapProcedureCode="RerouteNASRequest"
		packet.NgapType="RerouteNASRequest"
		RerouteNASRequest(InitiatingMessage.Value,packet)
	case ngapType.ProcedureCodeRRCInactiveTransitionReport:
		//fmt.Println("==================== RRCInactiveTransitionReport ====================")
		packet.NgapProcedureCode="RRCInactiveTransitionReport"
		packet.NgapType="RRCInactiveTransitionReport"
		RRCInactiveTransitionReport(InitiatingMessage.Value,packet)
	case ngapType.ProcedureCodeSecondaryRATDataUsageReport:
		//fmt.Println("==================== SecondaryRATDataUsageReport ====================")
		packet.NgapProcedureCode="SecondaryRATDataUsageReport"
		packet.NgapType="SecondaryRATDataUsageReport"
		SecondaryRATDataUsageReport(InitiatingMessage.Value,packet)
	case ngapType.ProcedureCodeTraceFailureIndication:
		//fmt.Println("==================== TraceFailureIndication ====================")
		packet.NgapProcedureCode="TraceFailureIndication"
		packet.NgapType="TraceFailureIndication"
		TraceFailureIndication(InitiatingMessage.Value,packet)
	case ngapType.ProcedureCodeTraceStart:
		//fmt.Println("==================== TraceStart ====================")
		packet.NgapProcedureCode="TraceStart"
		packet.NgapType="TraceStart"
		TraceStart(InitiatingMessage.Value,packet)
	case ngapType.ProcedureCodeUERadioCapabilityInfoIndication:
		//fmt.Println("==================== UERadioCapabilityInfoIndication ====================")
		packet.NgapProcedureCode="UERadioCapabilityInfoIndication"
		packet.NgapType="UERadioCapabilityInfoIndication"
		UERadioCapabilityInfoIndication(InitiatingMessage.Value,packet)
	case ngapType.ProcedureCodeUETNLABindingRelease:
	//	fmt.Println("==================== UETNLABindingRelease ====================")
		packet.NgapProcedureCode="UETNLABindingRelease"
		packet.NgapType="UETNLABindingReleaseRequest"
		UETNLABindingReleaseRequest(InitiatingMessage.Value,packet)
	case ngapType.ProcedureCodeUplinkNASTransport:
		//fmt.Println("==================== UplinkNASTransport ====================")
		packet.NgapProcedureCode="UplinkNASTransport"
		packet.NgapType="UplinkNASTransport"
		UplinkNASTransportDecoder(InitiatingMessage.Value,packet)
	case ngapType.ProcedureCodeUplinkNonUEAssociatedNRPPaTransport:
		//fmt.Println("==================== UplinkNonUEAssociateNRPPaTransport ====================")
		packet.NgapProcedureCode="UplinkNonUEAssociateNRPPaTransport"
		packet.NgapType="UplinkNonUEAssociateNRPPaTransport"
	case ngapType.ProcedureCodeUplinkRANConfigurationTransfer:
		//fmt.Println("==================== UplinkRANConfigurationTransfer ====================")
		packet.NgapProcedureCode="UplinkRANConfigurationTransfer"
		packet.NgapType="UplinkRANConfigurationTransfer"
	case ngapType.ProcedureCodeUplinkRANStatusTransfer:
		//fmt.Println("==================== UplinkRANStatusTransfer ====================")
		packet.NgapProcedureCode="UplinkRANStatusTransfer"
		packet.NgapType="UplinkRANStatusTransfer"
		UplinkRANStatusTransfer(InitiatingMessage.Value,packet)
	case ngapType.ProcedureCodeUplinkUEAssociatedNRPPaTransport:
		//fmt.Println("==================== UplinkUEAssociatedNRPPaTransport ====================")
		packet.NgapProcedureCode="UplinkUEAssociatedNRPPaTransport"
		packet.NgapType="UplinkUEAssociatedNRPPaTransport"
		UplinkUEAssociatedNRPPaTransport(InitiatingMessage.Value,packet)
	case ngapType.ProcedureCodeHandoverPreparation:
		//fmt.Println("==================== HandoverPreparation ====================")
		packet.NgapProcedureCode="HandoverPreparation"
		packet.NgapType="HandoverRequired"
		HandoverRequired(InitiatingMessage.Value,packet)
	case ngapType.ProcedureCodeHandoverResourceAllocation:
		//fmt.Println("==================== HandoverResourceAllocation ====================")
		packet.NgapProcedureCode="HandoverResourceAllocation"
		packet.NgapType="HandoverRequest"
	case ngapType.ProcedureCodeNGReset:
		//fmt.Println("==================== NGReset ====================")
		packet.NgapProcedureCode="NGReset"
		packet.NgapType="NGReset"
	case ngapType.ProcedureCodeWriteReplaceWarning:
		//fmt.Println("==================== WriteReplaceWarning ====================")
		packet.NgapProcedureCode="WriteReplaceWarning"
		packet.NgapType="WriteReplaceWarningRequest"
	}
}




func RouteSuccessfulOutcome(DecResult *ngapType.NGAPPDU,packet *flowMap.Packet) {
	SuccessfulOutcome := DecResult.SuccessfulOutcome
	switch SuccessfulOutcome.ProcedureCode.Value {
	case ngapType.ProcedureCodeNGSetup:
		//fmt.Println("==================== NGSetup ====================")
		packet.NgapProcedureCode="NGSetup"
		packet.NgapType="NGSetupResponse"
		NGSetupResponseDecoder(SuccessfulOutcome.Value,packet)
	case ngapType.ProcedureCodeNGReset:
		//fmt.Println("==================== NGReset ====================")
		packet.NgapProcedureCode="NGReset"
		packet.NgapType="NGResetAcknowledge"
	case ngapType.ProcedureCodeAMFConfigurationUpdate:
		//fmt.Println("==================== AMFConfigurationUpdate ====================")
		packet.NgapProcedureCode="AMFConfigurationUpdate"
		packet.NgapType="AMFConfigurationUpdateAcknowledge"
	case ngapType.ProcedureCodeHandoverCancel:
		//fmt.Println("==================== HandoverCancel ====================")
		packet.NgapProcedureCode="HandoverCancel"
		packet.NgapType="HandoverCancelAcknowledge"
		HandoverCancelAcknowledge(SuccessfulOutcome.Value,packet)
	case ngapType.ProcedureCodeHandoverPreparation:
		//fmt.Println("==================== HandoverPreparation ====================")
		packet.NgapProcedureCode="HandoverPreparation"
		packet.NgapType="HandoverCommand"
		HandoverCommand(SuccessfulOutcome.Value,packet)
	case ngapType.ProcedureCodeHandoverResourceAllocation:
		//fmt.Println("==================== HandoverResourceAllocation ====================")
		packet.NgapProcedureCode="HandoverResourceAllocation"
		packet.NgapType="HandoverRequestAcknowledge"
		HandoverRequestAcknowledge(SuccessfulOutcome.Value,packet)
	case ngapType.ProcedureCodeInitialContextSetup:
		//fmt.Println("==================== InitialContextSetup ====================")
		packet.NgapProcedureCode="InitialContextSetup"
		packet.NgapType="InitialContextSetupResponse"
		InitialContextSetupResDecoder(SuccessfulOutcome.Value,packet)
	case ngapType.ProcedureCodePathSwitchRequest:
		//fmt.Println("==================== PathSwitchRequest ====================")
		packet.NgapProcedureCode="PathSwitchRequest"
		packet.NgapType="PathSwitchRequestAcknowledge"
		PathSwitchRequestAcknowledge(SuccessfulOutcome.Value,packet)
	case ngapType.ProcedureCodePDUSessionResourceModify:
		//fmt.Println("==================== PDUSessionResourceModify ====================")
		packet.NgapProcedureCode="PDUSessionResourceModify"
		packet.NgapType="PDUSessionResourceModifyResponse"
		PDUSessionResourceModifySucDecoder(SuccessfulOutcome.Value,packet)
	case ngapType.ProcedureCodePDUSessionResourceModifyIndication:
		//fmt.Println("==================== PDUSessionResourceModifyIndication ====================")
		packet.NgapProcedureCode="PDUSessionResourceModifyIndication"
		packet.NgapType="PDUSessionResourceModifyConfirm"
		PDUSessionResourceModifyConfirm(SuccessfulOutcome.Value,packet)
	case ngapType.ProcedureCodePDUSessionResourceRelease:
		//fmt.Println("==================== PDUSessionResourceRelease ====================")
		packet.NgapProcedureCode="PDUSessionResourceRelease"
		packet.NgapType="PDUSessionResourceReleaseResponse"
		PDUSessionResourceReleaseDecoder(SuccessfulOutcome.Value,packet)
	case ngapType.ProcedureCodePDUSessionResourceSetup:
		//fmt.Println("==================== PDUSessionResourceSetup ====================")
		packet.NgapProcedureCode="PDUSessionResourceSetup"
		packet.NgapType="PDUSessionResourceSetupResponse"
		PDUSessionResourceSetupResDecoder(SuccessfulOutcome.Value,packet)
	case ngapType.ProcedureCodePWSCancel:
		//fmt.Println("==================== PWSCancel ====================")
		packet.NgapProcedureCode="PWSCancel"
		packet.NgapType="PWSCancelResponse"
	case ngapType.ProcedureCodeRANConfigurationUpdate:
		//fmt.Println("==================== RANConfigurationUpdateAcknowledge ====================")
		packet.NgapProcedureCode="ProcedureCodeRANConfiguration"
		packet.NgapType="RANConfigurationUpdateAcknowledge"
		RANConfigurationUpdateAcknowledgeDecoder()
	case ngapType.ProcedureCodeUEContextModification:
		//fmt.Println("==================== UEContextModification ====================")
		packet.NgapProcedureCode="UEContextModification"
		packet.NgapType="UEContextModificationResponse"
		UEContextModificationResponse(SuccessfulOutcome.Value,packet)
	case ngapType.ProcedureCodeUEContextRelease:
		//fmt.Println("==================== UEContextRelease ====================")
		packet.NgapProcedureCode="UEContextRelease"
		packet.NgapType="UEContextReleaseComplete"
		UEContextReleaseComplete(SuccessfulOutcome.Value,packet)
	case ngapType.ProcedureCodeUERadioCapabilityCheck:
		//fmt.Println("==================== UERadioCapabilityCheck ====================")
		packet.NgapProcedureCode="UERadioCapabilityCheck"
		packet.NgapType="UERadioCapabilityCheckResponse"
		UERadioCapabilityCheckResponse(SuccessfulOutcome.Value,packet)
	case ngapType.ProcedureCodeWriteReplaceWarning:
		//fmt.Println("==================== WriteReplaceWarning ====================")
		packet.NgapProcedureCode="WriteReplaceWarning"
		packet.NgapType="WriteReplaceWarningResponse"

	}
}

func RouteUnsuccessfulOutcome(DecResult *ngapType.NGAPPDU,packet *flowMap.Packet) {
	UnsuccessfulOutcome := DecResult.UnsuccessfulOutcome
	switch UnsuccessfulOutcome.ProcedureCode.Value {
	case ngapType.ProcedureCodeAMFConfigurationUpdate:
		//fmt.Println("==================== AMFConfigurationUpdate ====================")
		packet.NgapProcedureCode="AMFConfigurationUpdate"
		packet.NgapType="AMFConfigurationUpdateFailure"
	case ngapType.ProcedureCodeHandoverPreparation:
		//fmt.Println("==================== HandoverPrepararion ====================")
		packet.NgapProcedureCode="HandoverPrepararion"
		packet.NgapType="HandoverPreparationFailure"
		HandoverPreparationFailure(UnsuccessfulOutcome.Value,packet)
	case ngapType.ProcedureCodeHandoverResourceAllocation:
		//fmt.Println("==================== HandoverNotification ====================")
		packet.NgapProcedureCode="HandoverResourceAllocation"
		packet.NgapType="HandoverFailure"
	case ngapType.ProcedureCodeInitialContextSetup:
		//fmt.Println("==================== InitialContextSetup ====================")
		packet.NgapProcedureCode="InitialContextSetup"
		packet.NgapType="InitialContextSetupFailure"
		InitialContextSetupFailure(UnsuccessfulOutcome.Value,packet)
	case ngapType.ProcedureCodeNGSetup:
		//fmt.Println("==================== NGSetup ====================")
		packet.NgapProcedureCode="NGSetup"
		packet.NgapType="NGSetupFailure"
	case ngapType.ProcedureCodePathSwitchRequest:
		//fmt.Println("==================== PathSwitchRequest ====================")
		packet.NgapProcedureCode="PathSwitchRequest"
		packet.NgapType="PathSwitchRequestFailure"
		PathSwitchRequestFailure(UnsuccessfulOutcome.Value,packet)
	case ngapType.ProcedureCodeRANConfigurationUpdate:
		//fmt.Println("==================== RANConfigurationUpdate ====================")
		packet.NgapProcedureCode="RANConfigurationUpdate"
		packet.NgapType="RANConfigurationUpdateFailure"
	case ngapType.ProcedureCodeUEContextModification:
		//fmt.Println("==================== UEContextModification ====================")
		packet.NgapProcedureCode="UEContextModification"
		packet.NgapType="UEContextModificationFailure"
		UEContextModificationFailure(UnsuccessfulOutcome.Value,packet)
	}
}
