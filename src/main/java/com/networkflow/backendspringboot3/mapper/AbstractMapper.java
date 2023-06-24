package com.networkflow.backendspringboot3.mapper;

import com.baomidou.mybatisplus.core.mapper.BaseMapper;
import com.networkflow.backendspringboot3.model.domain.Abstract;
import org.springframework.stereotype.Repository;

@Repository
public interface AbstractMapper extends BaseMapper<Abstract> {
}
