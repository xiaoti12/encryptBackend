package PacketDB

import (
	"container/list"
	_ "container/list"
	_ "fmt"
	"github.com/Freddy/sctp_flowmap/database"
	"log"
)


func init() {
	creatPacketTable()
}

func InsertPacket(PacketList *list.List) {

	var sqlStr string = insertPacketSQL

	tx, err := database.Connect.Begin()
	checkErr(err)
	stmt, err := tx.Prepare(sqlStr)
	checkErr(err)
	for info := PacketList.Front(); info != nil; info = info.Next(){
		packet_sql := info.Value.(*Packet)
		if _, err := stmt.Exec(
			//packet_sql.PacketId,
			packet_sql.NgapType,
			packet_sql.NgapProcedureCode,
			packet_sql.RanUeNgapId,
			packet_sql.PacketLen,
			packet_sql.ArriveTimeUs,
			packet_sql.ArriveTime,
			packet_sql.TimeInterval,
			packet_sql.VerificationTag,
			packet_sql.SrcIP,
			packet_sql.DstIP,
			packet_sql.DirSeq,
			packet_sql.FlowUEID,
			packet_sql.FlowTimeID,
			packet_sql.StatusPacket,
		); err != nil {
			log.Fatal(err)
		}

	}
	checkErr(tx.Commit())
}
/*
func QueryFlows() {
	var sqlStr string = queryPacketSQL

	rows, err := database.Connect.Query(sqlStr)
	checkErr(err)
	for rows.Next() {
		var packet Packet
		checkErr(rows.Scan(
			&packet.
			))

		fmt.Println(packet) //printf log
	}
}

*/

func creatPacketTable() {
	var sqlStr string = creatPacketTableSQL

	_, err := database.Connect.Exec(sqlStr)
	checkErr(err)
}

func dropflowLogTable() {
	if _, err := database.Connect.Exec(dropPacketTableSQL); err != nil {
		log.Fatal(err)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}



