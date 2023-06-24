package com.networkflow.backendspringboot3.service.impl;

import com.baomidou.mybatisplus.core.conditions.query.QueryWrapper;
import com.baomidou.mybatisplus.extension.service.impl.ServiceImpl;
import com.networkflow.backendspringboot3.common.R;
import com.networkflow.backendspringboot3.mapper.PacketMapper;
import com.networkflow.backendspringboot3.model.domain.Packet;
import com.networkflow.backendspringboot3.service.PacketService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.math.BigInteger;

@Service
public class PacketServiceImpl extends ServiceImpl<PacketMapper, Packet> implements PacketService {
    @Autowired
    private PacketMapper packetMapper;

    @Override
    public R allPacket() {
        QueryWrapper<Packet> queryWrapper = new QueryWrapper<>();
        queryWrapper.lambda().orderByAsc(Packet::getArriveTime);
        return R.success(null, packetMapper.selectList(queryWrapper));
    }

    @Override
    public R getPacketByFlowId(BigInteger flowId) {
        QueryWrapper<Packet> queryWrapper = new QueryWrapper<>();
//        queryWrapper.lambda().eq(Packet::getFlowUEID, flowId).orderByAsc(Packet::getArriveTime);
        BigInteger startFlowId = flowId.subtract(new BigInteger(String.valueOf(1000)));
        BigInteger endFlowId = flowId.add(new BigInteger(String.valueOf(1000)));
        queryWrapper.lambda().between(Packet::getFlowUEID, startFlowId, endFlowId).orderByAsc(Packet::getArriveTime);

        return R.success(null, packetMapper.selectList(queryWrapper));
    }
}
