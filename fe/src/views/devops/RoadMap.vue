<template>
  <el-timeline style="max-width: 600px">
    <el-timeline-item center placement="top">
      <div>
        <h4>Learn a Programming Language</h4>
      </div>
      <el-row>
        <el-col :span="5">
          <el-card style="margin-right: 10px">
            <h4>Python</h4>
          </el-card>
        </el-col>
        <el-col :span="5">
          <el-card style="margin-right: 10px">
            <h4>Go</h4>
          </el-card>
        </el-col>
      </el-row>
    </el-timeline-item>
    <el-timeline-item center placement="top">
      <div>
        <h4>Operating Systems</h4>
      </div>
      <el-row>
        <el-col :span="5">
          <el-card style="margin-right: 10px">
            <h4>Linux</h4>
          </el-card>
        </el-col>
        <el-col :span="5">
          <el-card style="margin-right: 10px">
            <h4>Unix</h4>
          </el-card>
        </el-col>
      </el-row>
    </el-timeline-item>
  </el-timeline>
</template>
