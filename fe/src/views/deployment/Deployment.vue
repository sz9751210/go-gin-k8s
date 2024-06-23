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
      <!-- 數據表格 -->
      <el-col :span="24">
        <div>
          <!-- 包一層卡片 -->
          <el-card
            class="deploy-body-card"
            shadow="never"
            body-style="padding: 5px"
          >
            <!-- 數據表格 -->
            <!-- v-loading用於加載時的 loading 動畫 -->
            <el-table
              style="width: 100%; font-size: 12px; margin-bottom: 10px"
              :data="deploymentList"
              v-loading="appLoading"
            >
              <!-- 最左側留出 20px的寬度，更加沒關 -->
              <el-table-column width="20"> </el-table-column>
              <!-- deployment name -->
              <el-table-column align="left" label="Deployment Name">
                <template v-slot="scope">
                  <a class="deploy-body-deployname">
                    {{ scope.row.metadata.name }}
                  </a>
                </template>
              </el-table-column>
              <!-- 標籤 -->
              <el-table-column align="center" label="Labels">
                <template v-slot="scope">
                  <!-- for循環，每個label只顯示固定長度，滑鼠懸停後氣泡彈出框顯示完整長度  -->
                  <div
                    v-for="(label, key) in scope.row.metadata.labels"
                    :key="key"
                  >
                    <!-- 氣泡彈出框，用於提醒完整長度 -->
                    <!-- placement 彈出位置 -->
                    <!-- trigger 觸發條件 -->
                    <!-- content 彈出框內容 -->
                    <el-popover
                      placement="right"
                      :width="200"
                      trigger="hover"
                      :content="key + ':' + label"
                    >
                      <!-- <template #reference>{{
                        key + ":" + label
                      }}</template> -->
                      <template #reference>
                        <!-- ellipsis 方法用於剪裁字符串 -->
                        <el-tag style="margin-bottom: 5px" type="warning">{{
                          ellipsis(key + ":" + label)
                        }}</el-tag>
                      </template>
                    </el-popover>
                  </div>
                </template>
              </el-table-column>
              <!-- 容器組 -->
              <el-table-column align="center" label="Containers">
                <!-- 可用數量/總數量，三元運算，若值大於0則顯示，否則顯示0 -->
                <template v-slot="scope">
                  <span>
                    {{
                      scope.row.status.availableReplicas > 0
                        ? scope.row.status.availableReplicas
                        : 0
                    }}/{{
                      scope.row.spec.replicas > 0 ? scope.row.spec.replicas : 0
                    }}
                  </span>
                </template>
              </el-table-column>
              <!-- 創建時間 -->
              <el-table-column
                align="center"
                label="Created Time"
                min-width="100"
              >
                <template v-slot="scope">
                  <el-tag type="info">{{
                    timeTrans(scope.row.metadata.creationTimestamp)
                  }}</el-tag>
                </template>
              </el-table-column>
              <!-- Image -->
              <el-table-column align="center" label="Image" min-width="100">
                <template v-slot="scope">
                  <div
                    v-for="(container, index) in scope.row.spec.template.spec
                      .containers"
                    :key="index"
                  >
                    <el-popover
                      placement="right"
                      :width="200"
                      trigger="hover"
                      :content="container.image"
                    >
                      <template #reference>
                        <el-tag style="margin-bottom: 5px">
                          {{
                            ellipsis(
                              container.image.split(":")[2] == undefined
                                ? container.image
                                : container.image.split(":")[2]
                            )
                          }}
                        </el-tag>
                      </template>
                    </el-popover>
                  </div>
                </template>
              </el-table-column>
              <!-- 操作列，放按鈕 -->
              <el-table-column align="center" label="操作">
                <template v-slot="scope">
                  <el-button
                    size="small"
                    style="border-radius: 5px"
                    icon="Edit"
                    type="primary"
                    plain
                    @click="getDeploymentDetail(scope)"
                    >YAML
                  </el-button>
                  <el-button
                    size="small"
                    style="border-radius: 5px"
                    icon="Plus"
                    type="primary"
                    @click="handleScale(scope)"
                    >擴縮
                  </el-button>
                  <el-button
                    size="small"
                    style="border-radius: 5px"
                    icon="RefreshLeft"
                    type="primary"
                    @click="handleConfirm(scope, '重啟', 'restartDeployment')"
                  >
                    重啟
                  </el-button>
                  <el-button
                    size="small"
                    style="border-radius: 5px"
                    icon="Delete"
                    type="danger"
                    plain
                    @click="handleConfirm(scope, '刪除', 'delDeployment')"
                    >刪除
                  </el-button>
                </template>
              </el-table-column>
            </el-table>
            <!-- 分頁配置 -->
            <!-- background 背景色灰色 -->
            <!-- size-change 單頁大小改變後觸發 -->
            <!-- current-change 當前頁碼改變後觸發 -->
            <!-- current-page 當前頁碼 -->
            <!-- page-size 單頁大小 -->
            <!-- layout 分頁器支持的功能 -->
            <!-- total 數據總條數 -->
            <el-pagination
              class="deploy-body-pagination"
              background
              @size-change="handleSizeChange"
              @current-change="handleCurrentChange"
              :current-page="currentPage"
              :page-sizes="pagesizeList"
              :page-size="pageSize"
              layout="total, sizes, prev, pager, next, jumper"
              :total="deploymentTotal"
            >
            </el-pagination>
          </el-card>
        </div>
      </el-col>
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
    <!-- 展示 yaml 訊息的彈框 -->
    <el-dialog title="YAML INFO" v-model="yamlDialog" width="45%" top="2%">
      <!-- codemirror 編輯器 -->
      <!-- border 帶邊框 -->
      <!-- options 編輯器配置 -->
      <!-- change 編輯器中的內容變化時觸發 -->
      <codemirror
        :value="contentYaml"
        border
        :options="cmOptions"
        height="500"
        style="font-style: 14px"
        @change="onChange"
      ></codemirror>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="yamlDialog = false">Cancel</el-button>
          <el-button type="primary" @click="updateDeployment()"
            >Update</el-button
          >
        </span>
      </template>
    </el-dialog>
    <!-- 調整副本數 -->
    <el-dialog title="調整副本數" v-model="scaleDialog" width="25%">
      <el-input-number
        v-model="scaleNum"
        :min="0"
        :max="30"
        label="描述文字"
      ></el-input-number>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="scaleDialog = false">Cancel</el-button>
          <el-button type="primary" @click="scaleDeployment()"
            >Confirm</el-button
          >
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script>
import common from "../common/Config.js";
import httpClient from "../../utils/request.js";
import json2yaml from "json2yaml";
import yaml2obj from "js-yaml";
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
      // 分頁
      currentPage: 1,
      pagesize: 10,
      pagesizeList: [10, 20, 30],
      // 列表
      appLoading: false,
      deploymentList: [],
      deploymentTotal: 0,
      getDeploymentData: {
        url: common.k8sDeploymentList,
        params: {
          filter_name: "",
          namespace: "",
          page: "",
          limit: "",
        },
      },
      // 編輯器配置
      cmOptions: common.cmOptions,
      contentYaml: "",
      // 詳情
      deploymentDetail: {},
      getDeploymentDetailData: {
        url: common.k8sDeploymentDetail,
        params: {
          deployment_name: "",
          namespace: "",
        },
      },
      // yaml 更新
      yamlDialog: false,
      updateDeploymentData: {
        url: common.k8sDeploymentUpdate,
        params: {
          namespace: "",
          content: "",
        },
      },
      // 擴縮容
      scaleNum: 0,
      scaleDialog: false,
      scaleDeploymentData: {
        url: common.k8sDeploymentScale,
        params: {
          namespace: "",
          deployment_name: "",
          scale_num: "",
        },
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
    // 頁面數量發生變化時觸發
    handleSizeChange(size) {
      this.pagesize = size;
      this.getDeployments();
    },
    // 當前頁碼發生變化時觸發
    handleCurrentChange(currentPage) {
      this.currentPage = currentPage;
      this.getDeployments();
    },
    ellipsis(value) {
      return value.length > 15 ? value.substring(0, 15) + "..." : value;
    },
    // 格林威治時間轉換為本地時間
    timeTrans(timestamp) {
      let date = new Date(new Date(timestamp).getTime() + 8 * 3600 * 1000);
      date = date.toJSON();
      date = date.substring(0, 19).replace("T", " ");
      return date;
    },
    getDeployments() {
      // 表格加載動畫開啟
      this.appLoading = true;
      // getDeploymentsData用於發起deployment列表請求的專用對象，裡面有url和params參數，以下是賦值
      this.getDeploymentData.params.filter_name = this.searchInput;
      this.getDeploymentData.params.namespace = this.namespaceValue;
      this.getDeploymentData.params.page = this.currentPage;
      this.getDeploymentData.params.limit = this.pagesize;
      httpClient
        .get(this.getDeploymentData.url, {
          params: this.getDeploymentData.params,
        })
        .then((res) => {
          this.deploymentList = res.data.items;
          this.deploymentTotal = res.data.total;
          // 表格加載動畫關閉
          // this.appLoading = false;
        })
        .catch((res) => {
          this.$message.error({
            message: res.msg,
          });
        });
      // 表格加載動畫關閉
      this.appLoading = false;
    },
    // json to yaml
    transYaml(content) {
      return json2yaml.stringify(content);
    },
    // yaml to obj
    transObj(content) {
      return yaml2obj.load(content);
    },
    // 編輯器內容變化時觸發的方式
    // codemirror 的 value 是綁定 contentYaml
    onChange(val) {
      this.contentYaml = val;
    },
    // 獲取 deployment 詳情，e參數表示傳入的scope插槽，row是該行的數據
    getDeploymentDetail(e) {
      this.getDeploymentDetailData.params.deployment_name = e.row.metadata.name;
      this.getDeploymentDetailData.params.namespace = this.namespaceValue;
      httpClient
        .get(this.getDeploymentDetailData.url, {
          params: this.getDeploymentDetailData.params,
        })
        .then((res) => {
          // 響應成功，獲取詳情
          this.deploymentDetail = res.data;
          // console.log(this.deploymentDetail);
          // 將對象轉成 yaml 格式
          this.contentYaml = this.transYaml(this.deploymentDetail);
          console.log(this.contentYaml);
          // 打開彈出框
          // this.deploymentDetailDrawer = true;
          this.yamlDialog = true;
        })
        .catch((res) => {
          this.$message.error({
            message: res.msg,
          });
        });
    },
    // 擴縮容的中間方法
    handleScale(e) {
      // 打開彈出框
      this.scaleDialog = true;
      this.deploymentDetail = e.row;
      // 傳入當前 deployment 的副本數
      this.scaleNum = e.row.spec.replicas;
    },
    // scale deployment
    scaleDeployment() {
      this.scaleDeploymentData.params.deployment_name =
        this.deploymentDetail.metadata.name;
      this.scaleDeploymentData.params.namespace = this.namespaceValue;
      this.scaleDeploymentData.params.scale_num = this.scaleNum;
      httpClient
        .put(this.scaleDeploymentData.url, this.scaleDeploymentData.params)
        .then((res) => {
          this.$message.success({
            message: res.msg,
          });
          // 重新獲取 deployment 列表
          this.getDeployments();
        })
        .catch((res) => {
          this.$message.error({
            message: res.msg,
          });
        });
      this.scaleDialog = false;
    },
  },
  watch: {
    namespaceValue: {
      handler() {
        // 將 namespace 的值放入本地，用於 path 切換時依舊能獲取到
        localStorage.setItem("namespace", this.namespaceValue);
        this.currentPage = 1;
        this.getDeployments();
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
    this.getDeployments();
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
/* 數據表格 deployment name 顏色 */
.deploy-body-deployname {
  color: #4795ee;
}
/* deployment name hover */
.deploy-body-deployname:hover {
  color: rgb(84, 138, 238);
  cursor: pointer;
  font-weight: bold;
}
</style>
