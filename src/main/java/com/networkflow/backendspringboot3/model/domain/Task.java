package com.networkflow.backendspringboot3.model.domain;

import com.baomidou.mybatisplus.annotation.IdType;
import com.baomidou.mybatisplus.annotation.TableId;
import com.baomidou.mybatisplus.annotation.TableName;
import lombok.Data;
import org.apache.commons.lang3.builder.ToStringBuilder;

import java.time.LocalDateTime;

@TableName(value = "Task")
@Data
public class Task {
    @TableId(type = IdType.ASSIGN_UUID)
    private String taskId;

    private LocalDateTime createTime;
    private LocalDateTime startTime;
    private LocalDateTime endTime;

    // mode: 0 为离线流量检测，1 为在线流量检测
    private Integer mode;
    // model: 检测模型
    private String model;
    // PCAP包存储位置
    private String pcapPath;
    // 正常数据包数
    private Integer normal;
    // 异常数据包数
    private Integer abnormal;
    // 总数据包数
    private Integer total;
    // 网口名称
    private String netcard;

    // status：0 未启动；1 等待解析中；2 正在解析和检测中；3 完成解析聚合，正在检测中；4 正在汇总中；5 已完成。
    private Integer status;
    @Override
    public String toString() {
        return ToStringBuilder.reflectionToString(this);
    }

}
