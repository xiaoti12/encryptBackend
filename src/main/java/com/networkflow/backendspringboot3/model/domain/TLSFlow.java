package com.networkflow.backendspringboot3.model.domain;

import com.baomidou.mybatisplus.annotation.TableId;
import com.baomidou.mybatisplus.annotation.TableName;
import lombok.Data;
import org.apache.commons.lang3.builder.ToStringBuilder;

import java.time.LocalDateTime;

@TableName(value = "TLSFlow")
@Data
public class TLSFlow {
    @TableId
    private String TaskID;
    private String FlowID;
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

    @Override
    public String toString() {
        return ToStringBuilder.reflectionToString(this);
    }
}
