package com.networkflow.backendspringboot3.model.domain;

import io.swagger.v3.oas.models.security.SecurityScheme;
import lombok.Data;
import org.apache.commons.lang3.builder.ToStringBuilder;

@Data
public class Abstract {
    private Integer normal;
    private Integer abnormal;
    @Override
    public String toString() {
        return ToStringBuilder.reflectionToString(this);
    }
}
