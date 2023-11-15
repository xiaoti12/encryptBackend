package com.networkflow.backendspringboot3.model.domain;

import com.baomidou.mybatisplus.annotation.TableId;
import com.baomidou.mybatisplus.annotation.TableName;
import lombok.Data;
import org.apache.commons.lang3.builder.ToStringBuilder;

@Data
public class Admin {
    @TableId
    private String name;
    private String avatar;
    private String access;

    @Override
    public String toString() {
        return ToStringBuilder.reflectionToString(this);
    }
}
