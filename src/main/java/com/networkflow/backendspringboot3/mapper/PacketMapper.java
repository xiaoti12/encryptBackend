package com.networkflow.backendspringboot3.mapper;

import com.baomidou.mybatisplus.core.mapper.BaseMapper;
import com.networkflow.backendspringboot3.model.domain.Packet;
import org.springframework.stereotype.Repository;

@Repository
public interface PacketMapper extends BaseMapper<Packet> {

}
