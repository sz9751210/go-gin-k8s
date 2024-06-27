<template>
    <div>
      <h4>Grafana</h4>
      <el-table :data="tableData" border style="width: 100%">
        <el-table-column label="Name" width="180">
          <template v-slot="scope">
            <a :href="scope.row.link" target="_blank">{{ scope.row.name }}</a>
          </template>
        </el-table-column>
        <el-table-column prop="username" label="UserName" />
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
          },
          {
            link: "http://example.com",
            name: "Jerry",
            username: "No. 123, Elm St, New York",
          },
        ],
      };
    },
  };
  </script>
  