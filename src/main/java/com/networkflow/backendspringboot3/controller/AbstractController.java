package com.networkflow.backendspringboot3.controller;

import com.networkflow.backendspringboot3.common.R;
import com.networkflow.backendspringboot3.service.AbstractService;
import io.swagger.v3.oas.annotations.Operation;
import io.swagger.v3.oas.annotations.tags.Tag;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/api/abstract")
@Tag(name = "任务表接口")
public class AbstractController {
    @Autowired
    private AbstractService abstractService;
    @Operation(summary = "获取所有数据包信息")
    @GetMapping("/getAllAbstract")
    public R getAllAbstract() {
        return abstractService.allAbstract();
    }
}
