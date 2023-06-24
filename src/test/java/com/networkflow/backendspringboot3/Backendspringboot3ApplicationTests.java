package com.networkflow.backendspringboot3;



import com.networkflow.backendspringboot3.model.request.AdminRequest;
import com.networkflow.backendspringboot3.service.AbstractService;
import com.networkflow.backendspringboot3.service.impl.AbstractServiceImpl;
import com.networkflow.backendspringboot3.service.impl.AdminServiceImpl;
import com.networkflow.backendspringboot3.service.impl.TaskServiceImpl;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.scheduling.annotation.Scheduled;

import cn.hutool.log.LogFactory;
import cn.hutool.log.Log;

@SpringBootTest
class Backendspringboot3ApplicationTests {

    @Autowired
    private TaskServiceImpl taskService;
    @Autowired
    private AbstractServiceImpl abstractService;
    private static final Log log = LogFactory.get();
    @Test
    void contextLoads() {
        System.out.println(taskService.allTask());
    }
    @Test
    void test1() {
        System.out.println(abstractService.allAbstract());
    }

}
