package com.networkflow.backendspringboot3.service.impl;

import com.baomidou.mybatisplus.core.conditions.query.QueryWrapper;
import com.baomidou.mybatisplus.extension.service.impl.ServiceImpl;
import com.networkflow.backendspringboot3.common.R;
import com.networkflow.backendspringboot3.mapper.UEFlowMapper;
import com.networkflow.backendspringboot3.model.domain.UEFlow;
import com.networkflow.backendspringboot3.service.UEFlowService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class UEFlowServiceImpl extends ServiceImpl<UEFlowMapper, UEFlow> implements UEFlowService {
    @Autowired
    private UEFlowMapper ueFlowMapper;

    @Override
    public R allUEFlow() {
        QueryWrapper<UEFlow> queryWrapper = new QueryWrapper<>();
        queryWrapper.lambda().orderByAsc(UEFlow::getBeginTime);
        return R.success(null, ueFlowMapper.selectList(queryWrapper));
    }

    @Override
    public R deleteUEFlowsByTaskId(String taskId) {
        QueryWrapper<UEFlow> queryWrapper = new QueryWrapper<>();
        queryWrapper.lambda().eq(UEFlow::getTaskID, taskId);
        return R.success(null, ueFlowMapper.delete(queryWrapper));
    }

    @Override
    public R getUEFlowByTaskId(String taskId) {
        QueryWrapper<UEFlow> queryWrapper = new QueryWrapper<>();
        queryWrapper.lambda().eq(UEFlow::getTaskID, taskId).orderByAsc(UEFlow::getBeginTime);
        return R.success(null, ueFlowMapper.selectList(queryWrapper));
    }
}
