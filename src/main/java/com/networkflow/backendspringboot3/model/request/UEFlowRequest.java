package com.networkflow.backendspringboot3.model.request;

import lombok.Data;

import java.math.BigInteger;
import java.time.LocalDateTime;

@Data
public class UEFlowRequest {
    private BigInteger FlowId;
    private Integer RanUeNgapId;
    private Integer TotalNum;
    private Integer StartSecond;
    private Integer EndSecond;
    private LocalDateTime BeginTime;
    private LocalDateTime LatestTime;
    private Integer VerificationTag;
    private String SrcIP;
    private String DstIP;
    private Integer StatusFlow;
    private String TaskID;
}
