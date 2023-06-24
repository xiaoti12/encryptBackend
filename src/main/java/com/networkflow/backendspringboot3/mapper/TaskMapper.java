package com.networkflow.backendspringboot3.mapper;

import com.baomidou.mybatisplus.core.mapper.BaseMapper;
import com.networkflow.backendspringboot3.model.domain.Task;
import org.springframework.stereotype.Repository;

@Repository
public interface TaskMapper extends BaseMapper<Task> {

}
