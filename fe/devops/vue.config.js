const { defineConfig } = require("@vue/cli-service");
module.exports = defineConfig({
  devServer: {
    host: "0.0.0.0",
    port: 7070,
    open: true, // 啟動後自動打開頁面
  },
  transpileDependencies: true, // 忽略警告
  lintOnSave: false, // 禁用eslint, 關閉語法檢查
});
