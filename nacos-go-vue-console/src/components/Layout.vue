<template>
  <div class="layout">
    <!-- 顶部栏 -->
    <header class="header">
      <div class="left">Nacos-Go 控制台</div>
    </header>

    <!-- 导航栏 + 底部菜单 -->
    <nav class="nav">
      <div class="top">
        <router-link
          v-for="item in menu"
          :key="item.path"
          :to="item.path"
          class="nav-item"
          active-class="active"
        >
          {{ item.title }}
        </router-link>
      </div>

      <!-- 左侧底部退出 -->
      <div class="bottom">
        <a class="logout-item" @click="logout">退出登录</a>
      </div>
    </nav>

    <!-- 内容区 -->
    <main class="main">
      <router-view />
    </main>
  </div>
</template>

<script>
import { Notify } from '@/components/Notify'
export default {
  data: () => ({
    menu: [
      { path: '/tenant', title: '空间管理' },
      { path: '/service', title: '服务管理' },
      { path: '/config',  title: '配置管理' },
      { path: '/publish', title: '发布管理' },
      { path: '/user', title: '用户管理'},
      { path: '/profile', title: '个人中心'}
    ]
  }),
  methods: {
    logout() {
      this.$store.commit('SET_TOKEN', '')
      this.$router.replace('/login')
      Notify.success('已退出')
    }
  }
}
</script>

<style scoped>
.layout {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  background: #f7f7f7;
}

.header {
  display: flex;
  align-items: center;
  padding: 0 24px;
  height: 56px;
  background: #fff;
  box-shadow: 0 2px 6px rgba(0,0,0,.05);
  font-size: 18px;
  font-weight: 500;
  color: #333;
}

.nav {
  width: 200px;                     /* 左侧菜单宽度 */
  background: #fff;
  border-right: 1px solid #e6e6e6;
  display: flex;
  flex-direction: column;
  justify-content: space-between;   /* 顶部导航 + 底部退出 */
  padding: 12px 0;
  position: fixed;
  left: 0;
  top: 56px;
  bottom: 0;
}

.top {
  display: flex;
  flex-direction: column;
  gap: 4px;
  padding: 0 12px;
}

.nav-item {
  padding: 10px 14px;
  border-radius: 4px;
  font-size: 15px;
  color: #595959;
  text-decoration: none;
  transition: background .2s, color .2s;
}
.nav-item:hover {
  background: #f3f9ff;
  color: #1890ff;
}
.nav-item.active {
  background: #e6f7ff;
  color: #1890ff;
  font-weight: 500;
}

.bottom {
  padding: 0 12px;
}
.logout-item {
  display: block;
  padding: 10px 14px;
  border-radius: 4px;
  font-size: 15px;
  color: #595959;
  cursor: pointer;
  transition: background .2s;
}
.logout-item:hover {
  background: #ffece8;
  color: #ff4d4f;
}

.main {
  margin-left: 200px;   /* 对应 nav 宽度 */
  padding: 24px;
}
</style>