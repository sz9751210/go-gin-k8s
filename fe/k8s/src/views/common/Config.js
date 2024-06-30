export default {
  k8sNamespaceList: "http://localhost:9090/api/k8s/namespaces",
  k8sDeploymentCreate: "http://localhost:9090/api/k8s/deployment/create",
  k8sDeploymentList: "http://localhost:9090/api/k8s/deployments",
  k8sDeploymentDetail: "http://localhost:9090/api/k8s/deployment/detail",
  k8sDeploymentUpdate: "http://localhost:9090/api/k8s/deployment/update",
  k8sDeploymentScale: "http://localhost:9090/api/k8s/deployment/scale",
  k8sDeploymentRestart: "http://localhost:9090/api/k8s/deployment/restart",
  k8sDeploymentDelete: "http://localhost:9090/api/k8s/deployment/delete",
  // 編輯器設定
  cmOptions: {
    // 語言及語法模式
    mode: "text/yaml",
    // 主題
    theme: "idea",
    // 顯示行數
    lineNumbers: true,
    // 智能縮進
    smartIndent: true,
    // 智能縮進長度
    smartIndent: 4,
    // 顯示選中行的樣式
    styleActiveLine: true,
    // 當光標位於匹配的方括號旁邊時，都會高亮顯示
    matchBrackets: true,
    readOnly: false,
    // 自動換行
    lineWrapping: true,
  },
};
