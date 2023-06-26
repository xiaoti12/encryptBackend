package com.networkflow.backendspringboot3.service.impl;

import com.baomidou.mybatisplus.core.conditions.query.QueryWrapper;
import com.baomidou.mybatisplus.extension.service.impl.ServiceImpl;
import com.networkflow.backendspringboot3.common.R;
import com.networkflow.backendspringboot3.mapper.*;
import com.networkflow.backendspringboot3.model.domain.Abstract;
import com.networkflow.backendspringboot3.model.domain.Task;
import com.networkflow.backendspringboot3.model.domain.TimeFlow;
import com.networkflow.backendspringboot3.model.domain.UEFlow;
import com.networkflow.backendspringboot3.model.domain.TLSFlow;
import com.networkflow.backendspringboot3.service.AbstractService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.time.LocalDateTime;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.ArrayList;

@Service
public class AbstractServiceImpl extends ServiceImpl<AbstractMapper, Abstract> implements AbstractService {
    @Autowired
    private AbstractMapper abstractMapper;
    @Autowired
    private PacketMapper packetMapper;
    @Autowired
    private TaskMapper taskMapper;
    @Autowired
    private TimeFlowMapper timeFlowMapper;
    @Autowired
    private UEFlowMapper ueFlowMapper;
    @Autowired
    private TLSFlowMapper tlsFlowMapper;

    @Override
    public R allAbstract() {
        // task_status
        // status: 0(未开始),1(待解析),2(解析中),3(待检测),4(检测中),5(检测完成),100(错误)
        // flow_classification
        // status_origin: 0(未检测),100(检测完成且为正常)，200(检测完成且为异常)
        // status: 0(正常),1(恶意)

        // 介绍栏
        // 活跃任务
        Map<String, Integer> activeTask = new HashMap<>();
        // 活跃任务——在线任务数(统计任务表中有多少mode为1的任务)
        Long online = taskMapper.selectCount(new QueryWrapper<Task>().eq("mode", 1));
        // 活跃任务——离线任务数(统计任务表中有多少mode为1的任务)
        Long offline = taskMapper.selectCount(new QueryWrapper<Task>().lambda().eq(Task::getMode, 0));
        activeTask.put("online", online.intValue());
        activeTask.put("offline", offline.intValue());

        // 已完成任务数(按每天计算)(数据库中endtime的时间精确到分, 以天为单位，返回每天进行了多少任务)
        Map<String, Integer> completedTask = new HashMap<>();
        // 异常流数(统计UEFlow和TimeFlow中共有多少status为200)
        Map<String, Long> Abnormal = new HashMap<>();
        // 正常流数(统计UEFlow和TimeFlow中共有多少status为100)
        Map<String, Long> Normal = new HashMap<>();
        List<Task> completedTasks = taskMapper.selectList(new QueryWrapper<Task>().lambda().eq(Task::getStatus, 5));
        for (Task task : completedTasks) {
            LocalDateTime createTime = task.getCreateTime();
            String day = createTime.toLocalDate().toString();
            String taskId = task.getTaskId();

            Long abnormalFlow = tlsFlowMapper.selectCount(new QueryWrapper<TLSFlow>().lambda().eq(TLSFlow::getClassification, 1).eq(TLSFlow::getTaskID, taskId));
            Long normalFlow = tlsFlowMapper.selectCount(new QueryWrapper<TLSFlow>().lambda().eq(TLSFlow::getClassification, 0).eq(TLSFlow::getTaskID, taskId));

            completedTask.put(day, completedTask.getOrDefault(day, 0) + 1);
            Abnormal.put(day, Abnormal.getOrDefault(day, 0L) + abnormalFlow);
            Normal.put(day, Normal.getOrDefault(day, 0L) + normalFlow);
        }

        // 转换completedTask
        List<Map<String, Object>> completedTaskList = new ArrayList<>();
        for (Map.Entry<String, Integer> entry : completedTask.entrySet()) {
            String dateByDay = entry.getKey();
            Integer count = entry.getValue();

            Map<String, Object> completedTaskItem = new HashMap<>();
            completedTaskItem.put("dateByDay", dateByDay);
            completedTaskItem.put("completeTaskByDay", count);

            completedTaskList.add(completedTaskItem);
        }

        // 转换n2Abnormal
        List<Map<String, Object>> AbnormalList = new ArrayList<>();
        for (Map.Entry<String, Long> entry : Abnormal.entrySet()) {
            String dateByDay = entry.getKey();
            Long n2AbnormalByDay = entry.getValue();

            Map<String, Object> AbnormalItem = new HashMap<>();
            AbnormalItem.put("dateByDay", dateByDay);
            AbnormalItem.put("n2AbnormalByDay", n2AbnormalByDay);

            AbnormalList.add(AbnormalItem);
        }

        // 转换n2Normal
        List<Map<String, Object>> NormalList = new ArrayList<>();
        for (Map.Entry<String, Long> entry : Normal.entrySet()) {
            String dateByDay = entry.getKey();
            Long n2NormalByDay = entry.getValue();

            Map<String, Object> NormalItem = new HashMap<>();
            NormalItem.put("dateByDay", dateByDay);
            NormalItem.put("n2NormalByDay", n2NormalByDay);

            NormalList.add(NormalItem);
        }

        // 所有流
        Long abnormalFlowAll = tlsFlowMapper.selectCount(new QueryWrapper<TLSFlow>().lambda().eq(TLSFlow::getClassification, 1));
        Long normalFlowAll = tlsFlowMapper.selectCount(new QueryWrapper<TLSFlow>().lambda().eq(TLSFlow::getClassification, 0));
        Map<String, Long> abnormalFlowBinary = new HashMap<>();
        abnormalFlowBinary.put("normal", normalFlowAll);
        abnormalFlowBinary.put("abnormal", abnormalFlowAll);

        Map<String, Long> abnormalFlowMulti = new HashMap<>();
        abnormalFlowMulti.put("normal", normalFlowAll);
        abnormalFlowMulti.put("abnormal", abnormalFlowAll);

        // 异常事件(返回UEFlow和TimeFlow中所有status为1的流，并以时间倒序排序)
        QueryWrapper<TLSFlow> queryWrapper = new QueryWrapper<>();
        // queryWrapper.lambda().eq(UEFlow::getStatusFlow, 1).orderByDesc(UEFlow::getLatestTime);
        queryWrapper.lambda().eq(TLSFlow::getClassification, 1).orderByDesc(TLSFlow::getBeginTime);
        List<TLSFlow> abnormalEvent = tlsFlowMapper.selectList(queryWrapper);


        // 活跃流数——已检测流(统计UEFlow和TimeFlow中共有多少status为0的流)
        // 未被使用
        // Long activeDetectedFlows = ueFlowMapper.selectCount(new QueryWrapper<UEFlow>().in("StatusFlow", 100, 200))
        //         + timeFlowMapper.selectCount(new QueryWrapper<TimeFlow>().in("StatusFlow", 100, 200));

        // 活跃流数——待检测流(统计UEFlow和TimeFlow中共有多少status为1的流)
        // 未被使用
        // Long activePendingFlows = ueFlowMapper.selectCount(new QueryWrapper<UEFlow>().eq("StatusFlow", 0))
        //         + timeFlowMapper.selectCount(new QueryWrapper<TimeFlow>().eq("StatusFlow", 0));


        Map<String, Object> introduce = new HashMap<>();
        introduce.put("activeTask", activeTask);
        introduce.put("completedTask", completedTaskList);
        introduce.put("Abnormal", AbnormalList);
        introduce.put("Normal", NormalList);

        Map<String, Object> result = new HashMap<>();
        result.put("introduce", introduce);
        result.put("abnormalFlowBinary", abnormalFlowBinary);
        result.put("abnormalFlowMulti", abnormalFlowMulti);
        result.put("abnormalEvent", abnormalEvent);

        return R.success("success", result);
    }

}