package com.networkflow.backendspringboot3.mapper;

import com.baomidou.mybatisplus.core.mapper.BaseMapper;
import com.networkflow.backendspringboot3.model.domain.TLSFlow; // to change import 
import org.springframework.stereotype.Repository;

@Repository
public interface TLSFlowMapper extends BaseMapper<TLSFlow> {
}
