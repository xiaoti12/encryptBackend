package main

import (
	"flag"
	"fmt"
	"github.com/Freddy/sctp_flowmap/database"
	"github.com/Freddy/sctp_flowmap/decoder"
	"github.com/Freddy/sctp_flowmap/flowMap"
)

func main() {
	pcapPath := flag.String("pcap_path", "", "Path to the pcap file")
	taskID := flag.String("taskid", "", "The task ID")

	flag.Parse()
	_, err := database.Connect.Begin()
	database.CheckErr(err)

	if *pcapPath == "" {
		fmt.Println("pcap_path参数未提供")
		return
	}
	if *taskID == "" {
		fmt.Println("pcap_path参数未提供")
		return
	}
	decoder.Decode(*pcapPath, *taskID)
	//flowMap.ItFlowMap(flowMap.FlowTable_Time)
	/*
		List := flowMap.UEFlowMapToStore()


		TP := rulebased.ListGenerate(List)

		for tp:=TP.Front() ;tp!=nil;tp = tp.Next() {
			fmt.Print("UEId:")
			flowtp := tp.Value.(*rulebased.TypeList)
			fmt.Println(flowtp.Flow.RAN_UE_NGAP_ID)

			for ngap_list:=flowtp.NGAP_List.Front();ngap_list!=nil;ngap_list=ngap_list.Next(){
				fmt.Print(ngap_list.Value.(*rulebased.NGAPProcedureInfo).Procedure)
				fmt.Print("  status:")
				fmt.Println(ngap_list.Value.(*rulebased.NGAPProcedureInfo).Status)
			}

			packet_list := flowtp.Flow.PacketList
			for cur:=packet_list.Front();cur!=nil;cur=cur.Next(){
				packet :=cur.Value.(*flowMap.Packet)
				fmt.Print(packet.NgapProcedureCode)
				fmt.Print("  ")
				fmt.Println(packet.NgapRoute)
			}



		}

	*/

	//flowMap.ItFlowMap(flowMap.FlowTable_UE)
	//flowMap.ItFlowMap(flowMap.FlowTable_Time)
	flowMap.UEFlowMapToStore()
	flowMap.TimeFlowMapToStore()
	//flowMap.ItFlowMap(flowMap.FlowTable_Time)

}
