package com.networkflow.backendspringboot3.service;

import com.baomidou.mybatisplus.extension.service.IService;
import com.networkflow.backendspringboot3.common.R;
import com.networkflow.backendspringboot3.model.domain.Admin;
import com.networkflow.backendspringboot3.model.request.AdminRequest;
import jakarta.servlet.http.HttpSession;
import org.springframework.web.multipart.MultipartFile;

public interface AdminService extends IService<Admin> {

    R getCurrentUser();

    R outLogin();

    R login(String username, String password);
}
