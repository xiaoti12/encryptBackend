package com.networkflow.backendspringboot3.model.domain;

import com.baomidou.mybatisplus.annotation.TableId;
import com.baomidou.mybatisplus.annotation.TableName;
import lombok.Data;
import org.apache.commons.lang3.builder.ToStringBuilder;

import java.time.LocalDateTime;

@TableName(value = "TimeFlow")
@Data
public class TimeFlow {
    @TableId
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
    @Override
    public String toString() {
        return ToStringBuilder.reflectionToString(this);
    }
}
