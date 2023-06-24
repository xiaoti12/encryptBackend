package com.networkflow.backendspringboot3.service;

import com.baomidou.mybatisplus.extension.service.IService;
import com.networkflow.backendspringboot3.common.R;
import com.networkflow.backendspringboot3.model.domain.Task;
import com.networkflow.backendspringboot3.model.request.TaskRequest;
import org.springframework.web.multipart.MultipartFile;

public interface TaskService extends IService<Task> {
    // 获取所有任务
    R allTask();

    // 创建任务
    R createTask(TaskRequest createTaskRequest, MultipartFile file);

    // 更新任务状态
    R updateTaskStatus(String taskId, Integer status);

    // 更新任务
    R updateTask(TaskRequest createTaskRequest, MultipartFile uploadFile);

    // 删除任务
    R deleteTask(String[] taskIds);

    // 开始任务
    R startTask(String[] taskIds);

    // 根据Task修改
    boolean updateTaskByTask(Task task);

    // 定时任务处理Status
    void checkStatus();

}
