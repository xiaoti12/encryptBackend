package com.networkflow.backendspringboot3.service.impl;

import cn.hutool.log.Log;
import cn.hutool.log.LogFactory;
import com.baomidou.mybatisplus.core.conditions.query.QueryWrapper;
import com.baomidou.mybatisplus.extension.service.impl.ServiceImpl;
import com.networkflow.backendspringboot3.common.R;
import com.networkflow.backendspringboot3.mapper.PacketMapper;
import com.networkflow.backendspringboot3.mapper.TaskMapper;
import com.networkflow.backendspringboot3.mapper.TimeFlowMapper;
import com.networkflow.backendspringboot3.mapper.UEFlowMapper;
import com.networkflow.backendspringboot3.mapper.TLSFlowMapper;
import com.networkflow.backendspringboot3.model.domain.Task;
import com.networkflow.backendspringboot3.model.domain.TimeFlow;
import com.networkflow.backendspringboot3.model.domain.UEFlow;
import com.networkflow.backendspringboot3.model.domain.TLSFlow;
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
import java.time.LocalDateTime;
import java.time.format.DateTimeFormatter;
import java.util.Arrays;
import java.util.List;
import java.util.concurrent.Executor;
import java.util.concurrent.ThreadPoolExecutor;

@Service
public class TaskServiceImpl extends ServiceImpl<TaskMapper, Task> implements TaskService {
    private static final Log log = LogFactory.get();
    private final DetectTask detectTask;
    @Autowired
    private TaskMapper taskMapper;
    @Autowired
    private TimeFlowMapper timeFlowMapper;
    @Autowired
    private UEFlowMapper ueFlowMapper;
    @Autowired
    private TLSFlowMapper tlsFlowMapper;
    @Autowired
    private PacketMapper packetMapper;

    public TaskServiceImpl(DetectTask detectTask) {
        this.detectTask = detectTask;
    }

    @Override
    public R allTask() {
        QueryWrapper<Task> queryWrapper = new QueryWrapper<>();
        queryWrapper.lambda().orderByAsc(Task::getCreateTime);
        return R.success(null, taskMapper.selectList(queryWrapper));
    }

    @Override
    public R createTask(TaskRequest createTaskRequest, MultipartFile uploadFile) {
        Task task = new Task();

        BeanUtils.copyProperties(createTaskRequest, task);
        if (uploadFile != null) {
            String fileName = uploadFile.getOriginalFilename();
            String filePath = System.getProperty("user.dir") + System.getProperty("file.separator") + "core_python" + System.getProperty("file.separator") + "upload";
            File file = new File(filePath);
            if (!file.exists()) {
                if (!file.mkdir()) {
                    return R.fatal("创建文件失败");
                }
            }
            File dest = new File(filePath + System.getProperty("file.separator") + fileName);
            String storeUrlPath = fileName;
            try {
                uploadFile.transferTo(dest);
            } catch (IOException e) {
                return R.fatal("上传失败" + e.getMessage());
            }
            task.setPcapPath(storeUrlPath);
        } else
            task.setPcapPath(null);
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
        if (uploadFile != null) {
            String fileName = uploadFile.getOriginalFilename();
            String filePath = System.getProperty("user.dir") + System.getProperty("file.separator") + "core_python" + System.getProperty("file.separator") + "upload";
            File file = new File(filePath);
            if (!file.exists()) {
                if (!file.mkdir()) {
                    return R.fatal("创建文件失败");
                }
            }
            File dest = new File(filePath + System.getProperty("file.separator") + fileName);
            String storeUrlPath = fileName;
            try {
                uploadFile.transferTo(dest);
            } catch (IOException e) {
                return R.fatal("上传失败" + e.getMessage());
            }
            task.setPcapPath(storeUrlPath);
        } else
            task.setPcapPath(null);
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
            // 清除缓存
            // timeFlowMapper.delete(new QueryWrapper<TimeFlow>().lambda().eq(TimeFlow::getTaskID, taskId));
            // ueFlowMapper.delete(new QueryWrapper<UEFlow>().lambda().eq(UEFlow::getTaskID, taskId));
            tlsFlowMapper.delete(new QueryWrapper<TLSFlow>().lambda().eq(TLSFlow::getTaskID, taskId));

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

    @Scheduled(cron = "0/5 * *  * * ? ")
    @Override
    public void checkStatus() {
        log.info("执行Java的线程名字为 = " + Thread.currentThread().getName());

        QueryWrapper<Task> queryWrapper = new QueryWrapper<>();
        queryWrapper.lambda().eq(Task::getStatus, 1);
        List<Task> list = taskMapper.selectList(queryWrapper);

        log.info("list = " + list);
        for (Task task : list) {
            String currentTime = LocalDateTime.now().format(DateTimeFormatter.ofPattern("yyyy-MM-dd HH:mm:ss"));
            task.setStatus(2);
            task.setStartTime(LocalDateTime.parse(currentTime, DateTimeFormatter.ofPattern("yyyy-MM-dd HH:mm:ss")));
            taskMapper.updateById(task);
            if (taskMapper.updateById(task) > 0) {
                // detectTask.executeGoScript(System.getProperty("user.dir") + System.getProperty("file.separator") + "core" +
                //         System.getProperty("file.separator") + "springboot.py", task);
                detectTask.executePythonScript(System.getProperty("user.dir") + System.getProperty("file.separator") + "core_python" +
                        System.getProperty("file.separator") + "main.py", task);
            } else {
                log.info("启动成功");
            }
        }
    }
}

@Component
class DetectTask {
    private static final Log log = LogFactory.get();
    @Autowired
    private TaskMapper taskMapper;
    @Autowired
    private Executor checkTaskPool;

    @Async("checkTaskPool")
    public void executePythonScript(String scriptPath, Task currentTask) {
        log.info("执行Python的线程名字为 = " + Thread.currentThread().getName());
        try {
            String condaEnv = "tig";
            ProcessBuilder processBuilder = new ProcessBuilder("/home/fsc/anaconda3/envs/tig/bin/python", scriptPath, "--file_path", currentTask.getPcapPath(), "--taskid", currentTask.getTaskId(),"--model",currentTask.getModel());
            // processBuilder.directory(new File("/home/fsc/liujy/platform_display/backend/core_python"));
            Process process = processBuilder.start();

            // 处理脚本的输出
            BufferedReader reader = new BufferedReader(new InputStreamReader(process.getInputStream()));
            String line;
            while ((line = reader.readLine()) != null) {
                log.info(line);
            }
            int exitCode = process.waitFor();
            log.info("Python脚本执行完毕，退出码：" + exitCode);

            String currentTime = LocalDateTime.now().format(DateTimeFormatter.ofPattern("yyyy-MM-dd HH:mm:ss"));
            Task task = new Task();
            task.setTaskId(currentTask.getTaskId());
            if (exitCode == 0)
                task.setStatus(5);
            else
                task.setStatus(100);
            task.setEndTime(LocalDateTime.parse(currentTime, DateTimeFormatter.ofPattern("yyyy-MM-dd HH:mm:ss")));

            if (taskMapper.updateById(task) > 0) {
                if (exitCode == 0)
                    log.info("检测完成");
                else{
                    log.info("检测失败");
                }
            } else {
                log.info("检测失败");
            }
        } catch (IOException | InterruptedException e) {
            e.printStackTrace();
        } finally {
            if (checkTaskPool instanceof ThreadPoolExecutor) {
                ((ThreadPoolExecutor) checkTaskPool).remove(Thread.currentThread());
            }
        }
    }
    @Async("checkTaskPool")
    public void executeGoScript(String scriptPath, Task currentTask) {
        log.info("执行Go的线程名字为 = " + Thread.currentThread().getName());
        try {
            ProcessBuilder processBuilder = new ProcessBuilder("C:\\Users\\HorizonHe\\sdk\\go1.20.4\\bin\\go.exe", "run", "main.go", "--pcap_path", "..\\upload\\"+currentTask.getPcapPath(), "--taskid", currentTask.getTaskId());
            processBuilder.directory(new File("E:\\Code\\web\\backendspringboot3\\core_go\\sctp_flowmap"));
            Process process = processBuilder.start();

            BufferedReader reader = new BufferedReader(new InputStreamReader(process.getInputStream()));
            String line;
            while ((line = reader.readLine()) != null) {
                log.info(line);
            }

            int exitCode = process.waitFor();
            log.info("Go脚本执行完毕，退出码：" + exitCode);

            Task task = new Task();
            task.setTaskId(currentTask.getTaskId());
            if (exitCode == 0)
                task.setStatus(3);
            else {
                task.setStatus(100);
            }

            if (taskMapper.updateById(task) > 0) {
                if (exitCode == 0)
                    log.info("解析完成");
                else{
                    log.info("解析失败");
                    return;
                }
            } else {
                log.info("解析失败");
                return;
            }

            executePythonScript(System.getProperty("user.dir") + System.getProperty("file.separator") + "core_go\\python" +
                        System.getProperty("file.separator") + "springboot.py", currentTask);

        } catch (IOException | InterruptedException e) {
            e.printStackTrace();
        }
    }
}