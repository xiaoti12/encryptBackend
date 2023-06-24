package com.networkflow.backendspringboot3.service;

import com.baomidou.mybatisplus.extension.service.IService;
import com.networkflow.backendspringboot3.common.R;
import com.networkflow.backendspringboot3.model.domain.TimeFlow;
import com.networkflow.backendspringboot3.model.request.TimeFlowRequest;

public interface TimeFlowService extends IService<TimeFlow> {
    R allTimeFlow();

    R deleteTimeFlowsByTaskId(String taskId);
}
