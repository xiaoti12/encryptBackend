package com.networkflow.backendspringboot3.service;

import com.baomidou.mybatisplus.extension.service.IService;
import com.networkflow.backendspringboot3.common.R;
import com.networkflow.backendspringboot3.model.domain.UEFlow;

public interface UEFlowService extends IService<UEFlow> {
    R allUEFlow();

    R deleteUEFlowsByTaskId(String taskId);

    R getUEFlowByTaskId(String taskId);
}
