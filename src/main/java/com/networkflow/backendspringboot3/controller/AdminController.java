package com.networkflow.backendspringboot3.controller;

import com.networkflow.backendspringboot3.common.R;
import com.networkflow.backendspringboot3.model.request.AdminRequest;
import com.networkflow.backendspringboot3.service.AdminService;
import io.swagger.v3.oas.annotations.Operation;
import io.swagger.v3.oas.annotations.tags.Tag;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.multipart.MultipartFile;

@RestController
@RequestMapping("/api/admin")
@Tag(name = "管理员接口")
public class AdminController {
    @Autowired
    private AdminService adminService;

    @Operation(summary = "获取所有用户信息")
    @GetMapping("/getAllUser")
    public R getAllUser() {
        return adminService.allUser();
    }

    @Operation(summary = "添加用户信息")
    @PostMapping("/addUser")
    public R addUser(@RequestParam("id") Integer id,
                     @RequestParam("name") String name,
                     @RequestParam("password") String password,
                     @RequestParam(name = "file", required = false) MultipartFile file) {
        AdminRequest adminRequest = new AdminRequest();
        adminRequest.setId(id);
        adminRequest.setName(name);
        adminRequest.setPassword(password);
        return adminService.addUser(adminRequest, file);
    }

    @Operation(summary = "更新用户信息")
    @PostMapping("/updateUser")
    public R updateUser(@RequestBody AdminRequest adminRequest) {
        System.out.println(adminRequest);
        return adminService.updateUser(adminRequest);
    }

    @Operation(summary = "删除用户信息")
    @PostMapping("/deleteUser")
    public R deleteUser(@RequestParam("id") Integer userId) {
        return adminService.deleteUser(userId);
    }
}
