package com.networkflow.backendspringboot3.service.impl;

import cn.hutool.log.Log;
import cn.hutool.log.LogFactory;
import com.baomidou.mybatisplus.core.conditions.query.QueryWrapper;
import com.baomidou.mybatisplus.extension.service.impl.ServiceImpl;
import com.networkflow.backendspringboot3.common.R;
import com.networkflow.backendspringboot3.mapper.AdminMapper;
import com.networkflow.backendspringboot3.model.domain.Admin;
import com.networkflow.backendspringboot3.model.request.AdminRequest;
import com.networkflow.backendspringboot3.service.AdminService;
import jakarta.servlet.http.HttpSession;
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
import java.text.SimpleDateFormat;
import java.util.Date;
import java.util.List;
import java.util.concurrent.Executor;


@Service
public class AdminServiceImpl extends ServiceImpl<AdminMapper, Admin> implements AdminService {
    private static final Log log = LogFactory.get();
    private final ScriptTask scriptTask;
    @Autowired
    private AdminMapper adminMapper;

    public AdminServiceImpl(ScriptTask scriptTask) {
        this.scriptTask = scriptTask;
    }

    @Override
    public R verifyPasswd(AdminRequest adminRequest, HttpSession session) {
        QueryWrapper<Admin> queryWrapper = new QueryWrapper<>();
        queryWrapper.eq("name", adminRequest.getName());
        queryWrapper.eq("password", adminRequest.getPassword());
        if (adminMapper.selectCount(queryWrapper) > 0) {
            session.setAttribute("admin", adminRequest.getName());
            return R.success("登录成功");
        } else {
            return R.error("登录失败");
        }
    }

    @Override
    public R allUser() {
        QueryWrapper<Admin> queryWrapper = new QueryWrapper<>();
        queryWrapper.lambda().orderByAsc(Admin::getId);
        return R.success(null, adminMapper.selectList(queryWrapper));
    }

    @Override
    public R addUser(AdminRequest addUserRequest, MultipartFile upload_file) {
        Admin admin = new Admin();
        BeanUtils.copyProperties(addUserRequest, admin);
        if (upload_file != null) {
            String fileName = upload_file.getOriginalFilename();
            String filePath = System.getProperty("user.dir") + System.getProperty("file.separator") + "pcap";
            File file = new File(filePath);
            if (!file.exists()) {
                if (!file.mkdir()) {
                    return R.fatal("创建文件失败");
                }
            }
            File dest = new File(filePath + System.getProperty("file.separator") + fileName);
            String storeUrlPath = "/pcap/" + fileName;
            try {
                upload_file.transferTo(dest);
            } catch (IOException e) {
                return R.fatal("上传失败" + e.getMessage());
            }
            admin.setFilepath(storeUrlPath);
        } else
            admin.setFilepath(null);
        if (adminMapper.insert(admin) > 0) {
            return R.success("添加成功");
        } else {
            return R.error("添加失败");
        }
    }

    @Override
    public R updateUser(AdminRequest updateUserRequest) {
        Admin admin = new Admin();
        BeanUtils.copyProperties(updateUserRequest, admin);
        if (adminMapper.updateById(admin) > 0) {
            return R.success("更新成功");
        } else {
            return R.error("更新失败");
        }
    }

    @Override
    public R deleteUser(Integer userId) {
        if (adminMapper.deleteById(userId) > 0) {
            return R.success("删除成功");
        } else {
            return R.error("删除失败");
        }
    }

    // @Scheduled(cron="0/1 * *  * * ? ")
    @Override
    public void checkUser(){
        log.info("执行Java的线程名字为 = " + Thread.currentThread().getName());

        QueryWrapper<Admin> queryWrapper = new QueryWrapper<>();
        queryWrapper.lambda().eq(Admin::getId, 555);
        List<Admin> list = adminMapper.selectList(queryWrapper);
        for (Admin admin: list) {
            Admin newAdmin = new Admin();
            newAdmin.setId(admin.getId());
            newAdmin.setName("1233");
            newAdmin.setFilepath(admin.getFilepath());
            SimpleDateFormat sdf = new SimpleDateFormat();// 格式化时间
            sdf.applyPattern("yyyy-MM-dd HH:mm:ss");// a为am/pm的标记

            if (adminMapper.updateById(newAdmin) > 0) {
                log.info("Java中, 开始时间：" + sdf.format(new Date())); // 输出已经格式化的现在时间（24小时制）
                scriptTask.executePythonScript(System.getProperty("user.dir") + System.getProperty("file.separator") + "core" +
                        System.getProperty("file.separator") + "springboot.py");
                log.info("Java中, 结束时间：" + sdf.format(new Date())); // 输出已经格式化的现在时间（24小时制）
            } else {
                log.info("更新失败");
            }
        }
    }
}

@Component
class ScriptTask {
    private static final Log log = LogFactory.get();
    @Async("checkTaskPool")
    public void executePythonScript(String scriptPath) {
        log.info("执行Python的线程名字为 = " + Thread.currentThread().getName());
        try {
            String condaEnv = "base";
            ProcessBuilder processBuilder = new ProcessBuilder("D:\\ProgramData\\anaconda3\\Scripts\\conda.exe", "run", "-n", condaEnv, "python", scriptPath);
            Process process = processBuilder.start();

            // 处理脚本的输出
            BufferedReader reader = new BufferedReader(new InputStreamReader(process.getInputStream()));
            String line;
            while ((line = reader.readLine()) != null) {
                log.info(line);
            }
            int exitCode = process.waitFor();
            log.info("Python脚本执行完毕，退出码：" + exitCode);
        } catch (IOException | InterruptedException e) {
            e.printStackTrace();
        }
    }
}