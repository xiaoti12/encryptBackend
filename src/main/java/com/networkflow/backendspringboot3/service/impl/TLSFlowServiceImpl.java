package com.networkflow.backendspringboot3.service.impl;

import com.baomidou.mybatisplus.core.conditions.query.QueryWrapper;
import com.baomidou.mybatisplus.extension.service.impl.ServiceImpl;
import com.networkflow.backendspringboot3.common.R;
import com.networkflow.backendspringboot3.mapper.TLSFlowMapper; // change import here
import com.networkflow.backendspringboot3.model.domain.TLSFlow;
import com.networkflow.backendspringboot3.service.TLSFlowService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class TLSFlowServiceImpl extends ServiceImpl<TLSFlowMapper, TLSFlow> implements TLSFlowService {
    @Autowired
    private TLSFlowMapper tlsFlowMapper;

    @Override
    public R allTLSFlow() {
        QueryWrapper<TLSFlow> queryWrapper = new QueryWrapper<>();
        queryWrapper.lambda().orderByAsc(TLSFlow::getBeginTime);
        return R.success(null, tlsFlowMapper.selectList(queryWrapper));
    }

    @Override
    public R deleteTLSFlowsByTaskId(String taskId) {
        QueryWrapper<TLSFlow> queryWrapper = new QueryWrapper<>();
        queryWrapper.lambda().eq(TLSFlow::getTaskID, taskId);
        return R.success(null, tlsFlowMapper.delete(queryWrapper));
    }

    @Override
    public R getTLSFlowByTaskId(String taskId) {
        QueryWrapper<TLSFlow> queryWrapper = new QueryWrapper<>();
        queryWrapper.lambda().eq(TLSFlow::getTaskID, taskId).orderByAsc(TLSFlow::getBeginTime);
        return R.success(null, tlsFlowMapper.selectList(queryWrapper));
    }
}
