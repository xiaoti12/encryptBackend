package com.networkflow.backendspringboot3.controller;

import com.networkflow.backendspringboot3.common.R;
import com.networkflow.backendspringboot3.service.TimeFlowService;
import io.swagger.v3.oas.annotations.Operation;
import io.swagger.v3.oas.annotations.tags.Tag;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/api/timeflow")
@Tag(name = "时间聚合表接口")
public class TimeFlowController {
    @Autowired
    private TimeFlowService timeFlowService;
    @Operation(summary = "获取所有时间聚合流信息")
    @GetMapping("/getAllTimeFlow")
    public R getAllTask() {
        return timeFlowService.allTimeFlow();
    }
}
