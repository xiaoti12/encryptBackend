package com.networkflow.backendspringboot3.config;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.scheduling.annotation.EnableAsync;
import org.springframework.scheduling.concurrent.ThreadPoolTaskExecutor;

import java.util.concurrent.Executor;
import java.util.concurrent.ThreadPoolExecutor;

@Configuration
public class ThreadPoolConfig {
    @Bean(name="checkTaskPool")
    public Executor checkTaskPoolExecutor(){
        ThreadPoolTaskExecutor executor = new ThreadPoolTaskExecutor();
        int processNum = Runtime.getRuntime().availableProcessors(); // 返回可用处理器的Java虚拟机的数量
        int corePoolSize = (int) (processNum / (1 - 0.2));
        int maxPoolSize = (int) (processNum / (1 - 0.5));

        // 核心线程数：线程池创建时候初始化的线程数
        executor.setCorePoolSize(corePoolSize);
        // 最大线程数：线程池最大的线程数，只有在缓冲队列满了之后才会申请超过核心线程数的线程
        executor.setMaxPoolSize(maxPoolSize);
        // 缓冲队列：用来缓冲执行任务的队列
        executor.setQueueCapacity(maxPoolSize * 1000);
        // 允许线程的空闲时间300秒：当超过了核心线程之外的线程在空闲时间到达之后会被销毁
        executor.setKeepAliveSeconds(300);
        // 线程池名的前缀：设置好了之后可以方便我们定位处理任务所在的线程池
        executor.setThreadNamePrefix("checkTaskPoolExecutor-");

        // 缓冲队列满了之后的拒绝策略：由调用线程处理（一般是主线程）
        //AbortPolicy：丢弃任务并抛出 RejectedExecutionException 异常
        //DiscardPolicy：丢弃任务，但是不抛出异常。可能导致无法发现系统的异常状态
        //DiscardOldestPolicy：丢弃队列最前面的任务，然后重新提交被拒绝的任务
        //CallerRunsPolicy：不丢弃任务 由调用线程处理该任务
        executor.setRejectedExecutionHandler(new ThreadPoolExecutor.CallerRunsPolicy());
        // 线程池中corePoolSize线程空闲时间达到keepAliveTime也将关闭
        executor.setAllowCoreThreadTimeOut(true);
        executor.initialize();

        return executor;
    }
}
