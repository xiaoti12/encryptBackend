package rulebased

import (
	"container/list"
	"github.com/Freddy/sctp_flowmap/flowMap"
)

type TypeList struct {
	FlowID    uint64
	Flow      *flowMap.FlowInfo
	NGAP_List *list.List
	count     int
}


type TreeNode struct {
	NGAP_Procedure string
	FlowID         int64
	fatherNode     *TreeNode
	childrenNodes  *list.List
}

type NGAPProcedureInfo struct {
	Procedure string
	Status    uint
}

const (
	NoResponse    uint=0
	Success       uint=1
	Failure       uint=2
)

var dic_status = map[string]int{
	"InitiatingMessage":0,
	"SuccessfulOutcome":1,
	"UnsuccessfulOutcome":2}

var dict_pro = map[string]int{
	"PDUSessionResourceSetup":1,
	"PDUSessionResourceRelease":1,
	"PDUSessionResourceModify":1,
	"PDUSessionResourceModifyIndication":1,
	"UEContextRelease":1,
	"HandoverCancel":1,
	"UERadioCapabilityCheck":1,
	"InitialContextSetup":2,
	"UEContextModification":2,
	"HandoverPreparation":2,
	"PathSwitch":2,
	"PDUSessionResourceNotify":3,
	"UEContextReleaseRequest":3,
	"RRCInactiveTransitionReport":3,
	"HandoverRequestAcknowledge":3,
	"HandoverNotify":3,
	"UplinkRANStatusTransfer":3,
	"DownlinkRANStatusTransfer":3,
	"InitialUEMessage":3,
	"DownlinkNASTransport":3,
	"UplinkNASTransport":3,
	"NASNonDeliveryIndication":3,
	"RerouteNASRequest":3,
	"DownlinkUEAssociatedNRPPaTransport":3,
	"UplinkUEAssociatedNRPPaTransport":3,
	"TraceStart":3,
	"DeactivateTrace":3,
	"CellTrafficTrace":3,
	"LocationReportingControl":3,
	"LocationReportingFailureIndication":3,
	"LocationReport":3,
	"UETNLABindingReleaseRequest":3,
	"UERadioCapabilityInfoIndication":3,
	"SecondaryRATDataUsageReport":3,
	"ErrorIndication":4,
	"TraceFailureIndication":4}
