package com.networkflow.backendspringboot3.service.impl;

import com.baomidou.mybatisplus.core.conditions.query.QueryWrapper;
import com.baomidou.mybatisplus.extension.service.impl.ServiceImpl;
import com.networkflow.backendspringboot3.common.R;
import com.networkflow.backendspringboot3.mapper.TimeFlowMapper;
import com.networkflow.backendspringboot3.model.domain.TimeFlow;
import com.networkflow.backendspringboot3.service.TimeFlowService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class TimeFlowServiceImpl extends ServiceImpl<TimeFlowMapper, TimeFlow> implements TimeFlowService {
    @Autowired
    private TimeFlowMapper timeFlowMapper;

    @Override
    public R allTimeFlow() {
        QueryWrapper<TimeFlow> queryWrapper = new QueryWrapper<>();
        queryWrapper.lambda().orderByAsc(TimeFlow::getBeginTime);
        return R.success(null, timeFlowMapper.selectList(queryWrapper));
    }

    @Override
    public R deleteTimeFlowsByTaskId(String taskId) {
        QueryWrapper<TimeFlow> queryWrapper = new QueryWrapper<>();
        queryWrapper.lambda().eq(TimeFlow::getTaskID, taskId);
        return R.success(null, timeFlowMapper.delete(queryWrapper));
    }
}
