package com.networkflow.backendspringboot3.model.domain;

import com.baomidou.mybatisplus.annotation.TableId;
import com.baomidou.mybatisplus.annotation.TableName;
import lombok.Data;
import org.apache.commons.lang3.builder.ToStringBuilder;


@TableName(value = "admin")
@Data
public class Admin {
    @TableId
    private Integer id;
    private String name;
    private String password;
    private String filepath;
    @Override
    public String toString() {
        return ToStringBuilder.reflectionToString(this);
    }
}
