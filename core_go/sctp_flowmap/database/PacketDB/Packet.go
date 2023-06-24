package PacketDB

import (
	"fmt"
	"time"
)

type Packet struct {
	//PacketId          uint64 //包哈希id
	NgapType          string
	NgapProcedureCode string
	RanUeNgapId       uint64
	PacketLen         uint32
	ArriveTimeUs      uint64
	ArriveTime        time.Time
	TimeInterval      uint64
	VerificationTag   uint64
	SrcIP             string
	DstIP             string
	DirSeq            uint16
	FlowUEID          uint64
	FlowTimeID        uint64
	StatusPacket      uint16
}

var PacketTable = "SCTP.Packet"

var insertPacketSQL = `
		INSERT INTO ` + PacketTable +
	`
		(NgapType, NgapProcedureCode, RanUeNgapId, PacketLen, 
		ArriveTimeUs, ArriveTime, TimeInterval,
		VerificationTag, SrcIP, DstIP, DirSeq, FlowUEID, FlowTimeID, StatusPacket) 
		values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

var queryPacketSQL = `
		SELECT NgapType, NgapProcedureCode, RanUeNgapId, PacketLen, 
		ArriveTimeUs, ArriveTime, TimeInterval,
		VerificationTag, SrcIP, DstIP, DirSeq, FlowUEID, FlowTimeID, StatusPacket 
		FROM ` + PacketTable

var creatPacketTableSQL = `
	CREATE TABLE IF NOT EXISTS ` + PacketTable + ` (
	    NgapType String ,
	    NgapProcedureCode String ,
	    RanUeNgapId UInt64 ,
    	PacketLen UInt32 ,
    	ArriveTimeUs UInt64 ,
        ArriveTime DateTime64(6) ,
    	TimeInterval UInt64 ,
    	VerificationTag UInt64 ,
	    SrcIP String ,
	    DstIP String ,
	    DirSeq UInt16 ,
	    FlowUEID UInt64 ,
     	FlowTimeID UInt64 ,
	    StatusPacket UInt16 ,

		INDEX i_RanUeNgapId (RanUeNgapId) TYPE minmax GRANULARITY 4, 
		INDEX i_PacketLen (PacketLen) TYPE minmax GRANULARITY 4, 
		INDEX i_ArriveTimeUs (ArriveTimeUs) TYPE minmax GRANULARITY 4, 
		INDEX i_ArriveTime (ArriveTime) TYPE minmax GRANULARITY 4, 
		INDEX i_TimeInterval (TimeInterval) TYPE minmax GRANULARITY 4, 
		INDEX i_SrcIP (SrcIP) TYPE minmax GRANULARITY 4, 
		INDEX i_DstIP (DstIP) TYPE minmax GRANULARITY 4, 
		INDEX i_DirSeq (DirSeq) TYPE minmax GRANULARITY 4, 
		INDEX i_FlowUEID (FlowUEID) TYPE minmax GRANULARITY 4, 
		INDEX i_FlowTimeID (FlowTimeID) TYPE minmax GRANULARITY 4
		)  
		ENGINE = MergeTree()
        PARTITION BY toYYYYMMDD(ArriveTime)
		ORDER BY (ArriveTime)`

var dropPacketTableSQL = "DROP TABLE " + PacketTable

func (fl *Packet) initFlowLog() {
	//TODO init
}

func (fl Packet) String() string {
	return fmt.Sprintf(`
            NgapType: %s,
            NgapProcedureCode: %s,
		    RanUeNgapId: %u,
            PacketLen: %u,
            ArriveTimeUs: %u,
            ArriveTime: %s,
            TimeInterval: %u,
		    VerificationTag: %u,
		    SrcIP: %s, 
		    DstIP: %s,
            DirSeq: %u,
            FlowUEID: %u,
            FlowTimeID: %u,
		    StatusFlow: %u
		`, fl.NgapType, fl.NgapProcedureCode, fl.RanUeNgapId, fl.PacketLen,
		fl.ArriveTimeUs, fl.ArriveTime, fl.TimeInterval, fl.VerificationTag, fl.SrcIP, fl.DstIP,
		fl.DirSeq, fl.FlowUEID, fl.FlowTimeID, fl.StatusPacket)
}
