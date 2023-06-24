package TimeFlow

import (
	"fmt"
	"time"
)

type TimeFlow struct {
	FlowId          uint64 //流哈希id
	RanUeNgapId     uint64 //包哈希id
	TotalNum        uint32
	BeginTime       time.Time
	LatestTime      time.Time
	VerificationTag uint64
	SrcIP           string
	DstIP           string
	//TimeID          uint64
	StatusFlow uint16
	TaskID     string
}

var TimeFlowTable = "SCTP.TimeFlow"

var insertTimeFlowSQL = `
		INSERT INTO ` + TimeFlowTable +
	`
		(FlowId, RanUeNgapId, TotalNum, 
		BeginTime, LatestTime, VerificationTag, SrcIP, 
		DstIP,  StatusFlow, TaskID) 
		values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

var queryTimeFlowLogSQL = `
		SELECT FlowId, RanUeNgapId, TotalNum, 
		BeginTime, LatestTime, VerificationTag, SrcIP, 
		DstIP, StatusFlow ,TaskID
		FROM ` + TimeFlowTable

var creatTimeFlowTableSQL = `
	CREATE TABLE IF NOT EXISTS ` + TimeFlowTable + ` (
		FlowId UInt64,
		RanUeNgapId UInt64,
        TotalNum UInt32,
		BeginTime DateTime64(6), 
		LatestTime DateTime64(6), 
		VerificationTag UInt64,
		SrcIP String, 
		DstIP String,
		StatusFlow UInt16, 
        TaskID String,
		
		INDEX i_FlowId (FlowId) TYPE minmax GRANULARITY 4, 
		INDEX i_RanUeNgapId (RanUeNgapId) TYPE minmax GRANULARITY 4, 
		INDEX i_TotalNum (TotalNum) TYPE minmax GRANULARITY 4, 
		INDEX i_BeginTime (BeginTime) TYPE minmax GRANULARITY 4, 
		INDEX i_SrcIP (SrcIP) TYPE minmax GRANULARITY 4, 
		INDEX i_DstIP (DstIP) TYPE minmax GRANULARITY 4,
        INDEX i_TaskID (TaskID) TYPE minmax GRANULARITY 4
		)  
		ENGINE = MergeTree() 
		PARTITION BY toYYYYMMDD(BeginTime)
		ORDER BY (BeginTime)`

var dropTimeFlowTableSQL = "DROP TABLE " + TimeFlowTable

func (fl *TimeFlow) initFlowLog() {
	//TODO init
}

func (fl TimeFlow) String() string {
	return fmt.Sprintf(`
            FlowId: %u,
		    RanUeNgapId: %u,
            TotalNum: %u ,
	    	BeginTime: %s , 
		    LatestTime: %s , 
		    VerificationTag: %u ,
		    SrcIP: %s , 
		    DstIP: %s , 
		    StatusFlow: %u ,
            TaskID: %s,
		`, fl.FlowId, fl.RanUeNgapId, fl.TotalNum, fl.BeginTime,
		fl.LatestTime, fl.VerificationTag, fl.SrcIP, fl.DstIP, fl.StatusFlow, fl.TaskID)
}
