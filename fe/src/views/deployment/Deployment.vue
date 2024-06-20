<template>
  <div class="deploy">
    <el-row>
      <!-- 頭部一 -->
      <el-col :span="24">
        <div>
          <!-- 包一層卡片 -->
          <el-card class="deploy-head-card" shadow="never">
            <el-row>
              <!-- 命名空間下拉框 -->
              <el-col :span="6">
                <div>
                  <span>命名空間：</span>
                  <el-select
                    v-model="namespaceValue"
                    filterable
                    placeholder="請選擇"
                  >
                    <el-option
                      v-for="(item, index) in namespaceList"
                      :key="index"
                      :label="item.metadata.name"
                      :value="item.metadata.name"
                    >
                    </el-option>
                  </el-select>
                </div>
              </el-col>
              <el-col :span="2" :offset="16">
                <div>
                  <!-- 每次刷新，都重新調一次list接口，刷新表格中的數據 -->
                  <el-button style="border-radius: 2px" icon="Refresh" plain
                    >刷新</el-button
                  >
                </div>
              </el-col>
            </el-row>
          </el-card>
        </div>
      </el-col>
      <!-- 頭部二 -->
      <el-col :span="24"></el-col>
      <!-- 頭部三 -->
      <el-col :span="24"></el-col>
    </el-row>
  </div>
</template>

<script>
import common from "../common/Config.js";
import httpClient from "../../utils/request.js";
export default {
  data() {
    return {
      namespaceValue: "default",
      namespaceList: [],
      namespaceListUrl: common.k8sNamespaceList,
    };
  },
  methods: {
    getNamespaces() {
      httpClient
        .get(this.namespaceListUrl)
        .then((res) => {
          this.namespaceList = res.data.items;
        })
        .catch((res) => {
          this.$message.error({
            message: res.msg,
          });
        });
    },
  },
  watch: {
    namespaceValue: {
      handler() {
        // 將 namespace 的值放入本地，用於 path 切換時依舊能獲取到
        localStorage.setItem("namespace", this.namespaceValue);
      },
    },
  },
  beforeMount() {
    if (
      localStorage.getItem("namespace") !== undefined &&
      localStorage.getItem("namespace") !== null
    ) {
      this.namespaceValue = localStorage.getItem("namespace");
    }
    this.getNamespaces();
  },
};
</script>

<style>
.deploy-head-card {
  padding: "10px";
}
</style>
