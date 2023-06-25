package com.networkflow.backendspringboot3.model.request;

import lombok.Data;

import java.math.BigInteger;
import java.time.LocalDateTime;

@Data
public class TLSFlowRequest {
    private String TaskID;
    private String FlowId;
    private LocalDateTime BeginTime;
    private String SrcIP;
    private String DstIP;
    private Integer SrcPort;
    private Integer DstPort;
    private String Issuer;
    private String CommonName;
    private Integer Validity;
    private Float WhiteProb;
    private Float BlackProb;
    private String Classification;
}
