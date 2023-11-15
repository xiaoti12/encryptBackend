package com.networkflow.backendspringboot3.controller;

import com.networkflow.backendspringboot3.common.R;
import com.networkflow.backendspringboot3.service.AdminService;
import io.swagger.v3.oas.annotations.Operation;
import io.swagger.v3.oas.annotations.tags.Tag;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/api")
@Tag(name = "管理员接口")
public class AdminController {
    @Autowired
    private AdminService adminService;

    // Ant Design Pro
    @Operation(summary = "获取当前的用户")
    @GetMapping("/currentUser")
    public R getCurrentUser() {
        return adminService.getCurrentUser();
    }

    @Operation(summary = "退出登录接口")
    @PostMapping("/login/outLogin")
    public R outLogin() {
        return adminService.outLogin();
    }

    @Operation(summary = "登录接口")
    @PostMapping("/login/account")
    public R login(@RequestParam("username") String username,
            @RequestParam("password") String password) {
        return adminService.login(username, password);
    }
}
