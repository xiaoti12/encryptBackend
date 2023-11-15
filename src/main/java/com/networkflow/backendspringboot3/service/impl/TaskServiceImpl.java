package com.networkflow.backendspringboot3.service.impl;

import cn.hutool.log.Log;
import cn.hutool.log.LogFactory;
import com.baomidou.mybatisplus.core.conditions.query.QueryWrapper;
import com.baomidou.mybatisplus.extension.service.impl.ServiceImpl;
import com.networkflow.backendspringboot3.common.R;
import com.networkflow.backendspringboot3.mapper.TLSFlowMapper;
import com.networkflow.backendspringboot3.mapper.TaskMapper;
import com.networkflow.backendspringboot3.model.domain.TLSFlow;
import com.networkflow.backendspringboot3.model.domain.Task;
import com.networkflow.backendspringboot3.model.request.TaskRequest;
import com.networkflow.backendspringboot3.service.TaskService;
import org.springframework.beans.BeanUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.scheduling.annotation.Async;
import org.springframework.scheduling.annotation.Scheduled;
import org.springframework.stereotype.Component;
import org.springframework.stereotype.Service;
import org.springframework.web.multipart.MultipartFile;

import java.io.BufferedReader;
import java.io.File;
import java.io.IOException;
import java.io.InputStreamReader;
import java.nio.charset.StandardCharsets;
import java.time.LocalDateTime;
import java.time.format.DateTimeFormatter;
import java.util.Arrays;
import java.util.List;
import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;

@Service
public class TaskServiceImpl extends ServiceImpl<TaskMapper, Task> implements TaskService {
    private static final Log log = LogFactory.get();
    private final DetectTask detectTask;
    private final TaskManager taskManager;
    @Autowired
    private TaskMapper taskMapper;
    @Autowired
    private TLSFlowMapper tlsFlowMapper;

    public TaskServiceImpl(DetectTask detectTask, TaskManager taskManager) {
        this.detectTask = detectTask;
        this.taskManager = taskManager;
    }

    @Override
    public R allTask() {
        QueryWrapper<Task> queryWrapper = new QueryWrapper<>();
        queryWrapper.lambda().orderByDesc(Task::getCreateTime);
        return R.success(null, taskMapper.selectList(queryWrapper));
    }

    // 上传文件(返回任务id命名的名字)
    private String uploadFile(MultipartFile uploadFile, String taskId) {
        if (uploadFile == null) {
            return null;
        }
        String fileName = uploadFile.getOriginalFilename();

        // 检查文件存储位置是否存在
        String filePath = System.getProperty("user.dir") + System.getProperty("file.separator") + "corev2"
                + System.getProperty("file.separator") + "upload";
        File file = new File(filePath);
        if (!file.exists()) {
            if (!file.mkdir()) {
                return null;
            }
        }
        // 文件路径
        File dest = new File(filePath + System.getProperty("file.separator") + fileName);
        try {
            uploadFile.transferTo(dest);
            return fileName;
        } catch (IOException e) {
            return null;
        }
    }

    @Override
    public R createTask(TaskRequest createTaskRequest, MultipartFile uploadFile) {
        Task task = new Task();

        BeanUtils.copyProperties(createTaskRequest, task);

        String trueFileName = uploadFile(uploadFile, createTaskRequest.getTaskId());
        if (trueFileName != null) {
            task.setPcapPath(uploadFile.getOriginalFilename());
            // task.setTruePcapPath(trueFileName);
        } else {
            return R.fatal("上传文件失败");
        }

        if (taskMapper.insert(task) > 0) {
            return R.success("添加成功");
        } else {
            return R.error("添加失败");
        }
    }

    @Override
    public R updateTask(TaskRequest createTaskRequest, MultipartFile uploadFile) {
        Task task = new Task();
        BeanUtils.copyProperties(createTaskRequest, task);

        String trueFileName = uploadFile(uploadFile, createTaskRequest.getTaskId());
        if (trueFileName != null) {
            task.setPcapPath(uploadFile.getOriginalFilename());
            // task.setTruePcapPath(trueFileName);
        } else {
            return R.fatal("上传文件失败");
        }

        if (taskMapper.updateById(task) > 0) {
            return R.success("更新成功");
        } else {
            return R.error("更新失败");
        }
    }

    @Override
    public R updateTaskStatus(String taskId, Integer status) {
        Task task = new Task();
        task.setTaskId(taskId);
        task.setStatus(status);
        if (taskMapper.updateById(task) > 0) {
            return R.success("更新成功");
        } else {
            return R.error("更新失败");
        }
    }

    @Override
    public R deleteTask(String[] taskIds) {
        for (String taskId : taskIds) {
            // 检查文件存储位置是否存在
            String pcapname = taskMapper.selectById(taskId).getPcapPath();
            String filePath = System.getProperty("user.dir") + System.getProperty("file.separator") + "corev2" +
                    System.getProperty("file.separator") + "upload" + System.getProperty("file.separator") + pcapname;
            File file = new File(filePath);
            if (file.exists()) {
                boolean deleted = file.delete();
                if (deleted) {
                    log.info("文件删除成功");
                } else {
                    log.info("文件删除失败");
                }
            } else {
                log.info("文件不存在");
            }
        }
        if (taskMapper.deleteBatchIds(Arrays.asList(taskIds)) > 0) {
            return R.success("删除成功");
        } else {
            return R.error("删除失败");
        }
    }

    @Override
    public boolean updateTaskByTask(Task task) {
        return taskMapper.updateById(task) > 0;
    }

    @Override
    public R startTask(String[] taskIds) {
        int successCount = 0;
        for (String taskId : taskIds) {
            Task task = new Task();
            task.setTaskId(taskId);
            task.setStatus(1);
            task.setStartTime(null);
            task.setEndTime(null);
            task.setAbnormal(null);
            task.setNormal(null);
            task.setTotal(null);
            // 清除缓存
            deleteCache(taskId);
            if (taskMapper.updateById(task) > 0) {
                successCount++;
            }
        }
        if (successCount == taskIds.length) {
            return R.success("开始成功");
        } else if (successCount > 0 && successCount < taskIds.length) {
            return R.success("部分开始成功");
        } else {
            return R.error("开始失败");
        }
    }

    private void deleteCache(String taskId) {
        tlsFlowMapper.delete(new QueryWrapper<TLSFlow>().lambda().eq(TLSFlow::getTaskID, taskId));
    }

    @Override
    public R exitTask(String[] taskIds) {
        int successCount = 0;
        for (String taskId : taskIds) {
            taskManager.stopTask(taskId);
            Task task = new Task();
            task.setTaskId(taskId);
            task.setStatus(200);
            task.setEndTime(null);
            task.setAbnormal(null);
            task.setNormal(null);
            task.setTotal(null);
            if (taskMapper.updateById(task) > 0) {
                successCount++;
            }
        }
        if (successCount == taskIds.length) {
            return R.success("停止成功");
        } else if (successCount > 0 && successCount < taskIds.length) {
            return R.success("部分停止成功");
        } else {
            return R.error("停止失败");
        }
    }

    @Scheduled(cron = "0/5 * *  * * ? ")
    @Override
    public void checkStatus() {
        // log.info("轮询数据库, 线程名字为 = " + Thread.currentThread().getName());

        QueryWrapper<Task> queryWrapper = new QueryWrapper<>();
        queryWrapper.lambda().eq(Task::getStatus, 1);
        List<Task> list = taskMapper.selectList(queryWrapper);

        for (Task task : list) {
            String currentTime = LocalDateTime.now().format(DateTimeFormatter.ofPattern("yyyy-MM-dd HH:mm:ss"));
            task.setStatus(2);
            task.setStartTime(LocalDateTime.parse(currentTime, DateTimeFormatter.ofPattern("yyyy-MM-dd HH:mm:ss")));
            taskMapper.updateById(task);
            if (taskMapper.updateById(task) > 0) {
                // detectTask.executeGoScript(System.getProperty("user.dir") +
                // System.getProperty("file.separator") + "core" +
                // System.getProperty("file.separator") + "springboot.py", task);
                detectTask.executePythonScript(
                        System.getProperty("user.dir") + System.getProperty("file.separator") + "corev2" +
                                System.getProperty("file.separator") + "main.py",
                        task);

            } else {
                log.info("启动失败");
            }
        }
    }
}

@Component
class DetectTask {
    private static final Log log = LogFactory.get();
    private final TaskManager taskManager;
    @Autowired
    private TaskMapper taskMapper;

    DetectTask(TaskManager taskManager) {
        this.taskManager = taskManager;
    }

    @Async("checkTaskPool")
    public void executePythonScript(String scriptPath, Task currentTask) {
        try {
            String line;
            BufferedReader reader;
            String pyEnv = "/home/fsc/anaconda3/envs/tig/bin/python";
            // Go解析脚本
            log.info("任务: " + currentTask.getTaskId() + " 执行检测, 线程名字为 = " + Thread.currentThread().getName());
            ProcessBuilder processBuilder;
            if (currentTask.getMode() == 0) {
                processBuilder = new ProcessBuilder(pyEnv, scriptPath,
                        "--taskid", currentTask.getTaskId(),
                        "--model", currentTask.getModel(),
                        "--mode", "pcap",
                        "--file_path", currentTask.getPcapPath());
            } else {
                processBuilder = new ProcessBuilder(pyEnv, scriptPath,
                        "--taskid", currentTask.getTaskId(),
                        "--model", currentTask.getModel(),
                        "--mode", "live",
                        "--interface", currentTask.getNetcard());
            }
            processBuilder.redirectErrorStream(true); // 合并标准输出和标准错误流
            Process process = processBuilder.start();
            taskManager.addTaskProcess(currentTask.getTaskId(), process);

            // log.info("任务: " + currentTask.getTaskId() + " 检测脚本运行的PID为:" + process.pid());
            reader = new BufferedReader(new InputStreamReader(process.getInputStream(), StandardCharsets.UTF_8));
            while ((line = reader.readLine()) != null) {
                log.info("任务: " + currentTask.getTaskId() + " " + line);
            }
            int exitCode = process.waitFor();
            reader.close();
            taskManager.stopTask(currentTask.getTaskId());

            Task task = new Task();
            task.setTaskId(currentTask.getTaskId());
            if (exitCode == 0)
                task.setStatus(5);
            else
                task.setStatus(100);
            String currentTime = LocalDateTime.now().format(DateTimeFormatter.ofPattern("yyyy-MM-dd HH:mm:ss"));
            task.setEndTime(LocalDateTime.parse(currentTime, DateTimeFormatter.ofPattern("yyyy-MM-dd HH:mm:ss")));
            taskMapper.updateById(task);

            log.info("任务: " + currentTask.getTaskId() + " 检测成功, 已停止, 退出码为: " + exitCode);
        } catch (IOException | InterruptedException e) {
            e.printStackTrace();
            Task task = new Task();
            task.setTaskId(currentTask.getTaskId());
            task.setStatus(100);
            String currentTime = LocalDateTime.now().format(DateTimeFormatter.ofPattern("yyyy-MM-dd HH:mm:ss"));
            task.setEndTime(LocalDateTime.parse(currentTime, DateTimeFormatter.ofPattern("yyyy-MM-dd HH:mm:ss")));
            taskMapper.updateById(task);
            taskManager.stopTask(currentTask.getTaskId());
            log.info("任务: " + currentTask.getTaskId() + " 检测失败, 已停止");
        }
    }
}

@Component
class TaskManager {
    private static final Log log = LogFactory.get();
    private final Map<String, Process> taskProcesses = new ConcurrentHashMap<>();

    public void addTaskProcess(String taskId, Process process) {
        taskProcesses.put(taskId, process);
        log.info("添加任务: " + taskId + " Map中任务数: " + taskProcesses.size());
    }

    public void stopTask(String taskId) {
        Process process = taskProcesses.get(taskId);
        if (process != null) {
            process.destroy();
            taskProcesses.remove(taskId);
        }
        log.info("删除任务: " + taskId + " Map中任务数: " + taskProcesses.size());
    }
}