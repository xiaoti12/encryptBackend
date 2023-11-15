package com.networkflow.backendspringboot3.service.impl;

import com.baomidou.mybatisplus.extension.service.impl.ServiceImpl;
import com.networkflow.backendspringboot3.common.R;
import com.networkflow.backendspringboot3.mapper.AdminMapper;
import com.networkflow.backendspringboot3.model.domain.Admin;
import com.networkflow.backendspringboot3.service.AdminService;
import org.springframework.stereotype.Service;
import java.util.HashMap;
import java.util.Map;

@Service
public class AdminServiceImpl extends ServiceImpl<AdminMapper, Admin> implements AdminService {
    private static String access = "";

    @Override
    public R getCurrentUser() {
        if (getAccess().isEmpty()) {
            return R.error("请先登录");
        }
        Map<String, Object> user = new HashMap<>();
        user.put("name", "Admin");
        user.put("avatar", "/buptlogo.png");
        // 返回假数据: name:Admin, "avatar:/buptlogo.png"
        return R.success(null, user);
    }

    @Override
    public R outLogin() {
        access = "";
        Map<String, Object> user = new HashMap<>();
        user.put("name", "");
        user.put("avatar", "");
        return R.success("退出成功", user);
    }

    @Override
    public R login(String username, String password) {
        if (username.equals("admin") && password.equals("admin")) {
            access = "admin";
            return R.success("ok");
        }
        return R.success("error");
    }

    private String getAccess() {
        return access;
    }
}