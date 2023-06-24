package com.networkflow.backendspringboot3.model.request;

import lombok.Data;

import java.time.LocalDateTime;

@Data
public class TimeFlowRequest {
    private Integer FlowId;
    private Integer RanUeNgapId;
    private Integer TotalNum;
    private LocalDateTime BeginTime;
    private LocalDateTime LatestTime;
    private Integer VerificationTag;
    private String SrcIP;
    private String DstIP;
    private Integer StatusFlow;
    private String TaskID;
}
