<template>
  <div class="pod">
    <el-row>
      <!-- header 1 -->
      <el-col :span="24">
        <div>
          <!-- 包一層卡片 -->
          <el-card
            class="pod-head-card"
            shadow="never"
            :body-style="{ padding: '10px' }"
          >
            <el-row>
              <!-- 命名空間下拉框 -->
              <el-col :span="6">
                <div>
                  <span>命名空間： </span>
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
              <!-- 刷新按鈕 -->
              <el-col :span="2" :offset="16">
                <div>
                  <!-- 每次刷新，都重新調用一次 List 接口，刷新表格數據 -->
                  <el-button style="border-radius: 2px" icon="Refresh" plain
                    >刷新
                  </el-button>
                </div>
              </el-col>
            </el-row>
          </el-card>
        </div>
      </el-col>
      <!-- header 2 -->
      <el-col :span="24"></el-col>
      <!-- body 1 -->
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
            message:
              "name: " +
              res.name +
              " msg: " +
              res.message +
              " code: " +
              res.code,
          });
        });
    },
  },
  watch: {
    namespaceValue: {
      handler() {
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

<style scoped>
.pod-head-card,
.pod-body-card {
  border-radius: 1px;
  margin-bottom: 5px;
}
</style>
