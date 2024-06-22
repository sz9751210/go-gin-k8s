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
      <el-col :span="24">
        <div>
          <el-card
            class="deploy-head-card"
            shadow="never"
            body-style="padding: 10px"
          >
            <el-row>
              <!-- 創建按鈕 -->
              <el-col :span="2">
                <div>
                  <!-- 點擊打開抽屜，填入創建deployment需要的數據 -->
                  <el-button
                    style="border-radius: 2px"
                    icon="Edit"
                    type="primary"
                    @click="createDeploymentDrawer = true"
                    v-loading.fullscreen.lock="fullscreenLoading"
                    >創建</el-button
                  >
                </div>
              </el-col>
              <!-- 搜索框和搜索按鈕 -->
              <el-col :span="6">
                <div>
                  <el-input
                    class="deploy-head-search"
                    clearable
                    placeholder="請輸入"
                    v-model="searchInput"
                  ></el-input>
                  <el-button
                    style="border-radius: 2px"
                    icon="Search"
                    type="primary"
                    plain
                    >搜索</el-button
                  >
                </div>
              </el-col>
            </el-row>
          </el-card>
        </div>
      </el-col>
      <!-- 頭部三 -->
      <el-col :span="24"></el-col>
    </el-row>
    <!-- 抽屜：創建Deployment表單 -->
    <!-- v-model值是bool，用於顯示與隱藏 -->
    <!-- direction 顯示的位置 -->
    <!-- before-close 關閉時觸發，點擊關閉或者點擊空白都會觸發 -->
    <el-drawer
      v-model="createDeploymentDrawer"
      :direction="direction"
      :before-close="handleClose"
    >
      <!-- 插槽，抽屜標題 -->
      <template #title>
        <h4>創建Deployment</h4>
      </template>
      <!-- 插槽，抽屜內容 -->
      <template #default>
        <!-- flex佈局，居中 -->
        <el-row type="flex" justify="center">
          <el-col :span="20">
            <!-- ref綁定控制元件後，js中才能用this.$ref獲取該控制元件 -->
            <!-- rules 定義 form 表單驗證 -->
            <el-form
              ref="createDeployment"
              :rules="createDeploymentRules"
              :model="createDeployment"
              label-width="80px"
            >
              <!-- prop用於rules中的驗證規則key -->
              <el-form-item class="deploy-create-form" label="名稱" prop="name">
                <el-input v-model="createDeployment.name"></el-input>
              </el-form-item>
              <el-form-item
                class="deploy-create-form"
                label="命名空間"
                prop="namespace"
              >
                <el-select
                  v-model="createDeployment.namespace"
                  filterable
                  placeholder="請選擇"
                >
                  <el-option
                    v-for="(namespace, index) in namespaceList"
                    :key="index"
                    :label="namespace.metadata.name"
                    :value="namespace.metadata.name"
                  >
                  </el-option>
                </el-select>
              </el-form-item>
              <!-- 數字輸入框，最小為1，最大為10 -->
              <el-form-item
                class="deploy-create-form"
                label="副本數量"
                prop="replicas"
              >
                <el-input-number
                  v-model="createDeployment.replicas"
                  :min="1"
                  :max="10"
                >
                </el-input-number>
                <!-- 氣泡彈出框，用於提醒上限 -->
                <el-popover
                  placement="top"
                  :width="100"
                  trigger="hover"
                  content="申請副本數上限為10"
                >
                  <template #reference>
                    <el-icon style="width: 2em; font-size: 18px; color: #4795ee"
                      ><WarningFilled
                    /></el-icon>
                  </template>
                </el-popover>
              </el-form-item>
              <el-form-item
                class="deploy-create-form"
                label="Image"
                prop="image"
              >
                <el-input v-model="createDeployment.image"> </el-input>
              </el-form-item>
              <el-form-item
                class="deploy-create-form"
                label="label"
                prop="label_str"
              >
                <el-input
                  v-model="createDeployment.label_str"
                  placeholder="ex. app=k8s-demo,env=prod"
                ></el-input>
              </el-form-item>
              <el-form-item
                class="deploy-create-form"
                label="Resource"
                prop="resource"
              >
                <el-select
                  v-model="createDeployment.resource"
                  plcaeholder="請選擇"
                >
                  <el-option value="0.5/1" label="CPU:0.5 Mem:1G"></el-option>
                  <el-option value="1/2" label="CPU:1 Mem:2G"></el-option>
                  <el-option value="2/4" label="CPU:2 Mem:4G"></el-option>
                  <el-option value="4/8" label="CPU:4 Mem:8G"></el-option>
                </el-select>
              </el-form-item>
              <el-form-item
                class="deploy-create-form"
                label="ContainerPort"
                prop="container_port"
              >
                <el-input
                  v-model="createDeployment.container_port"
                  placeholder="ex. 80"
                ></el-input>
              </el-form-item>
              <el-form-item
                class="deploy-create-form"
                label="HealthCheck"
                prop="health_check"
              >
                <el-switch v-model="createDeployment.health_check" />
              </el-form-item>
              <el-form-item
                class="deploy-create-form"
                label="HealthPath"
                prop="health_path"
              >
                <el-input
                  v-model="createDeployment.health_path"
                  placeholder="ex. /health"
                ></el-input>
              </el-form-item>
            </el-form>
          </el-col>
        </el-row>
      </template>
      <!-- 插槽，footer -->
      <template #footer>
        <!-- 點擊後賦值false，關閉抽屜 -->
        <el-button @click="createDeploymentDrawer = false">取消</el-button>
        <el-button type="primary" @click="submitForm('createDeployment')"
          >立即創建</el-button
        >
      </template>
    </el-drawer>
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
      // 搜尋框內容
      searchInput: "",
      // 創建
      fullscreenLoading: false,
      direction: "rtl",
      createDeploymentDrawer: false,
      createDeployment: {
        name: "",
        namespace: "",
        replicas: 1,
        image: "",
        resource: "",
        health_check: false,
        health_path: "",
        label_str: "",
        label: {},
        container_port: "",
      },
      // 創建請求參數
      createDeploymentData: {
        url: common.k8sDeploymentCreate,
        params: {},
      },
      // 創建請求規則
      createDeploymentRules: {
        name: [{ required: true, message: "請輸入名稱", trigger: "change" }],
        image: [{ required: true, message: "請輸入image", trigger: "change" }],
        namespace: [
          { required: true, message: "請輸入namespace", trigger: "change" },
        ],
        resource: [
          { required: true, message: "請輸入resource", trigger: "change" },
        ],
        label_str: [
          { required: true, message: "請輸入label", trigger: "change" },
        ],
        container_port: [
          {
            required: true,
            message: "請輸入container_port",
            trigger: "change",
          },
        ],
      },
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
    // 處理關閉抽屜
    handleClose(done) {
      this.$confirm("確定關閉嗎?")
        .then(() => {
          done();
        })
        .catch(() => {});
    },
    // 創建 deployment,加 Func 原因是因為 createDeploy 用在屬性了
    createDeployFunc() {
      // 正則匹配，驗證 label 格式
      let reg = new RegExp("(^[A-Za-z]+=[A-Za-z0-9]+).*");
      if (!reg.test(this.createDeployment.label_str)) {
        this.$message.warning({
          message: "label 格式錯誤",
        });
        return;
      }
      // 加載loading動畫
      this.fullscreenLoading = true;
      // 定義 label、cpu、memory變量
      let label = new Map();
      let cpu, memory;
      let a = this.createDeployment.label_str.split(",");
      a.forEach((item) => {
        let b = item.split("=");
        label[b[0]] = b[1];
      });
      // 將deployment規格轉成cpu和memory
      let resourceList = this.createDeployment.resource.split("/");
      cpu = resourceList[0];
      memory = resourceList[1];
      this.createDeploymentData.params = {
        name: this.createDeployment.name,
        // namespace: this.createDeployment.namespace,
        // replicas: this.createDeployment.replicas,
        image: this.createDeployment.image,
        cpu: cpu,
        memory: memory,
        health_check: this.createDeployment.health_check,
        health_path: this.createDeployment.health_path,
        label: label,
        container_port: parseInt(this.createDeployment.container_port),
      };
      httpClient
        .post(this.createDeploymentData.url, this.createDeploymentData.params)
        .then((res) => {
          this.$message.success({
            message: res.msg,
          });
          this.getDeployments();
        })
        .catch((res) => {
          this.$message.error({
            message: res.msg,
          });
        }),
        // 重置表單
        this.resetForm("createDeployment");
      // 關閉加載動畫
      this.fullscreenLoading = false;
      // 關閉抽屜
      this.createDeploymentDrawer = false;
    },
    // 重置表單
    resetForm(formName) {
      this.$refs[formName].resetFields();
    },
    // 提交表單
    submitForm(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          this.createDeployFunc();
          // // 關閉抽屜
          // this.createDeploymentDrawer = false;
          // // 重置輸入框
          // this.createDeployment = {
          //   name: "",
          //   namespace: "",
          //   replicas: 1,
          //   image: "",
          //   resource: "",
          //   health_check: false,
          //   health_path: "",
          //   label_str: "",
          //   label: {},
          //   container_port: "",
          // };
        } else {
          return false;
        }
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

<style scoped>
/* 卡片屬性 */
.deploy-head-card,
.deplou-body-card {
  border-radius: 1px;
  margin-bottom: 5px;
}
/* 搜尋框 */
.deploy-head-search {
  width: 160px;
  margin-right: 10px;
}
</style>
