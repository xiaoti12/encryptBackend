package com.networkflow.backendspringboot3.controller;

import com.networkflow.backendspringboot3.common.R;
import com.networkflow.backendspringboot3.service.TLSFlowService;
import io.swagger.v3.oas.annotations.Operation;
import io.swagger.v3.oas.annotations.tags.Tag;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/api/tlsflow")
@Tag(name = "TLS表接口")
public class TLSFlowController {
    @Autowired
    private TLSFlowService tlsFlowService;
    @Operation(summary = "获取所有用户信息")
    @GetMapping("/getAllTLSFlow")
    public R getAllTask() {
        return tlsFlowService.allTLSFlow();
    }

    @Operation(summary = "获取所有TLS聚合流信息")
    @GetMapping("/getTLSFlowByTaskId")
    public R getTLSFlowByTaskId(@RequestParam("taskId") String taskId) {
        int index = taskId.indexOf('?');
        if (index != -1) {
            taskId = taskId.substring(0, index);
        }
        return tlsFlowService.getTLSFlowByTaskId(taskId);
    }
}
