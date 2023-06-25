package com.networkflow.backendspringboot3.service;

import com.baomidou.mybatisplus.extension.service.IService;
import com.networkflow.backendspringboot3.common.R;
import com.networkflow.backendspringboot3.model.domain.TLSFlow; //change import here

public interface UEFlowService extends IService<TLSFlow> {
    R allTLSFlow();

    R deleteTLSFlowsByTaskId(String taskId);

    R getTLSFlowByTaskId(String taskId);
}
