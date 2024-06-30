<template>
  <div>
    <h4>Grafana</h4>
    <el-table :data="tableData" border style="width: 100%">
      <el-table-column label="Website" width="180">
        <template v-slot="scope">
          <a :href="scope.row.link" target="_blank" class="custom-link">{{ scope.row.name }}</a>
        </template>
      </el-table-column>
      <el-table-column width="250" prop="username" label="UserName" />
      <el-table-column label="PassWord">
        <template v-slot="scope">
          <div style="display: flex; align-items: center">
            <span>{{
              scope.row.showPassword ? scope.row.password : "******"
            }}</span>
            <el-icon
              @click="togglePassword(scope.row)"
              style="margin-left: 10px; cursor: pointer"
            >
              <component :is="scope.row.showPassword ? 'hide' : 'view'" />
            </el-icon>
            <el-button
              @click="copyPassword(scope.row.password)"
              style="
                margin-left: 10px;
                border: none;
                background-color: transparent;
              "
              icon="copy-document"
            >
            </el-button>
          </div>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
export default {
  data() {
    return {
      tableData: [
        {
          link: "http://example.com",
          name: "Tom",
          username: "No. 189, Grove St, Los Angeles",
          password: "aaaaaaaaaaa",
          showPassword: false,
        },
        {
          link: "http://example.com",
          name: "Jerry",
          username: "No. 123, Elm St, New York",
          password: "123456",
          showPassword: false,
        },
      ],
    };
  },
  methods: {
    togglePassword(row) {
      row.showPassword = !row.showPassword;
    },
    async copyPassword(password) {
      try {
        await navigator.clipboard.writeText(password);
        this.$message({
          message: "Password copied to clipboard",
          type: "success",
        });
      } catch (err) {
        this.$message({
          message: "Failed to copy password",
          type: "error",
        });
      }
    },
  },
};
</script>
<style scoped>
.custom-link {
  color: #409EFF; /* Element Plus primary color */
  text-decoration: none;
  font-weight: bold;
  transition: color 0.3s ease;
}

.custom-link:hover {
  color: #66b1ff;
  text-decoration: underline;
}
</style>