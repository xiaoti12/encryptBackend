package com.networkflow.backendspringboot3.model.request;

import lombok.Data;


@Data
public class AdminRequest {
    private Integer id;
    private String name;
    private String password;
    private String filepath;
}
