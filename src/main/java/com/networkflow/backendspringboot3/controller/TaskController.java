package com.networkflow.backendspringboot3.controller;

import com.networkflow.backendspringboot3.common.R;
import com.networkflow.backendspringboot3.model.request.TaskRequest;
import com.networkflow.backendspringboot3.service.TaskService;
import io.swagger.v3.oas.annotations.Operation;
import io.swagger.v3.oas.annotations.tags.Tag;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.multipart.MultipartFile;

import java.time.LocalDateTime;
import java.time.format.DateTimeFormatter;

@RestController
@RequestMapping("/api/task")
@Tag(name = "任务表接口")
public class TaskController {
    @Autowired
    private TaskService taskService;

    @Operation(summary = "获取所有用户信息")
    @GetMapping("/getAllTask")
    public R getAllTask() {
        return taskService.allTask();
    }

    @Operation(summary = "新建任务")
    @PostMapping("/createTask")
    public R createTask(@RequestParam("taskId") String taskId,
                        @RequestParam("createTime") String createTime,
                        @RequestParam("mode") Integer mode,
                        @RequestParam(name = "model", required = false) String model,
                        @RequestParam("status") Integer status,
                        @RequestParam(name = "pcapFile", required = false) MultipartFile file) {
        TaskRequest taskRequest = new TaskRequest();
        taskRequest.setTaskId(taskId);
        taskRequest.setCreateTime(LocalDateTime.parse(createTime, DateTimeFormatter.ofPattern("yyyy-MM-dd HH:mm:ss")));
        taskRequest.setMode(mode);
        taskRequest.setModel(model);
        taskRequest.setStatus(status);
        return taskService.createTask(taskRequest, file);
    }

    @Operation(summary = "更新任务信息")
    @PostMapping("/updateTask")
    public R updateTask(@RequestParam("taskId") String taskId,
                        @RequestParam("mode") Integer mode,
                        @RequestParam(name = "model", required = false) String model,
                        @RequestParam(name = "pcapFile", required = false) MultipartFile file) {
        TaskRequest taskRequest = new TaskRequest();
        taskRequest.setTaskId(taskId);
        taskRequest.setMode(mode);
        taskRequest.setModel(model);
        return taskService.updateTask(taskRequest, file);
    }


    @Operation(summary = "更新任务状态")
    @PostMapping("/updateTaskStatus")
    public R updateTaskStatus(@RequestParam("taskId") String taskId,
                              @RequestParam("status") Integer status) {
        return taskService.updateTaskStatus(taskId, status);
    }

    @Operation(summary = "删除任务")
    @PostMapping("/deleteTask")
    public R deleteTask(@RequestParam("taskId") String[] taskId) {
        return taskService.deleteTask(taskId);
    }

    @Operation(summary = "开始任务")
    @PostMapping("/startTask")
    public R startTask(@RequestParam("taskId") String[] taskId) {
        return taskService.startTask(taskId);
    }
}
