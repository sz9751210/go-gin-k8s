<template>
  <div class="common-layout">
    <!-- container 佈局 -->
    <el-container style="height: 100vh">
      <!-- 側邊導覽 -->
      <el-aside class="aside" :width="asideWidth"
        >Aside
        <el-affix class="aside-affix">
          <div class="aside-logo">
            <el-image class="logo-image" :src="logo" />
            <span :class="[isCollapse ? 'is-collapse' : '']">
              <span class="logo-name">Kubernetes</span>
            </span>
          </div>
        </el-affix>
        <!-- router 定義vue-router模式，菜單欄的index跟路由規則中的path綁定 -->
        <!-- default-active 默認激活的菜單欄，這裡根據打開的path來找到對應的欄 -->
        <el-menu
          class="aside-menu"
          router
          :default-active="$route.path"
          :collapse="isCollapse"
          background-color="#131b27"
          text-color="#bfcbd9"
          active-text-color="#20a0ff"
        >
          <!-- routers 就是路由規則，router/index.js 中的routes -->
          <div v-for="menu in routers" :key="menu">
            <!-- 第一種情況，children 只有一個子菜單 -->
            <el-menu-item
              class="aside-menu-item"
              v-if="menu.children && menu.children.length == 1"
              :index="menu.path"
            >
              <!-- 處理圖標和菜單欄的名字 -->
              <el-icon>
                <component :is="menu.children[0].icon" />
              </el-icon>
              <template #title>
                {{ menu.children[0].name }}
              </template>
            </el-menu-item>
            <!-- 第二種情況，children 大於一個子菜單  -->
            <el-sub-menu
              class="aside-submenu"
              v-else-if="menu.children && menu.children.length > 1"
              :index="menu.path"
            >
              <!-- 處理父菜單欄 -->
              <template #title>
                <el-icon>
                  <component :is="menu.icon" />
                </el-icon>
                <span :class="[isCollapse ? 'is-collapse' : '']">{{
                  menu.name
                }}</span>
              </template>
              <!-- 處理子菜單 -->
              <el-menu-item
                class="aside-menu-childitem"
                v-for="child in menu.children"
                :key="child"
                :index="child.path"
              >
                <template #title>
                  {{ child.name }}
                </template>
              </el-menu-item>
            </el-sub-menu>
          </div>
        </el-menu>
      </el-aside>
    </el-container>

    <!-- <el-container>
        <el-header>Header</el-header>
        <el-main><router-view></router-view></el-main>
        <el-footer>Footer</el-footer>
    </el-container> -->
  </div>
</template>

<script>
import { useRouter } from "vue-router";

export default {
  data() {
    return {
      logo: require("@/assets/k8s-metrics.png"),
      asideWidth: "220px",
      isCollapse: false,
      routers: [],
    };
  },
  beforeMount() {
    this.routers = useRouter().options.routes;
  },
};
</script>

<style scoped>
/* 側邊欄折疊速度，背景色 */
.aside {
  transition: all 0.5s;
  background-color: #131b27;
}
.aside-logo {
  background-color: #131b27;
  height: 60px;
  color: white;
}
.logo-image {
  width: 40px;
  height: 40px;
  top: 12px;
  padding-left: 12px;
}
.logo-name {
  font-size: 20px;
  font-weight: bold;
  padding: 10px;
}
.is-collapse {
  display: none;
}
/* 菜單欄滾軸不顯示 */
.aside::-webkit-scrollbar {
  display: none;
}
.aside-affix {
  border-bottom-width: 0;
}
/* 右邊邊匡寬度 */
/* .aside-menu{
  border-right-width: 0;
} */

/* 菜单栏的位置以及颜色 */
/* 內邊距左邊的距離 */
/* .aside-menu-item {
  padding-left: 20px !important;
} */
.aside-menu-item.is-active {
  background-color: #1f2a3a;
}
/* 鼠標移過去的背景色變化 */
.aside-menu-item:hover {
  background-color: #142c4e;
}

.aside-submenu{
  padding-left: 0px !important;
}

.aside-menu-childitem {
  padding-left: 40px !important;
}
.aside-menu-childitem.is-active {
  background-color: #1f2a3a;
}
.aside-menu-childitem:hover {
  background-color: #142c4e;
}
</style>
