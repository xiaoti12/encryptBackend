package com.networkflow.backendspringboot3.model.domain;

import com.baomidou.mybatisplus.annotation.TableId;
import com.baomidou.mybatisplus.annotation.TableName;
import lombok.Data;
import org.apache.commons.lang3.builder.ToStringBuilder;

import java.time.LocalDateTime;

@TableName(value = "Packet")
@Data
public class Packet {
    private String NgapType;
    private String NgapProcedureCode;
    private Integer RanUeNgapId;
    private Integer PacketLen;
    private Integer ArriveTimeUs;
    private LocalDateTime ArriveTime;
    private Integer TimeInterval;
    private Integer VerificationTag;
    private String SrcIP;
    private String DstIP;
    private Integer DirSeq;
    private Integer FlowUEID;
    private Integer FlowTimeID;
    private Integer StatusPacket;

    @Override
    public String toString() {
        return ToStringBuilder.reflectionToString(this);
    }
}
