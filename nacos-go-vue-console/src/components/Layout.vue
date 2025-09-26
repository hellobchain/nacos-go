<template>
  <div class="layout">
    <!-- 顶部栏 -->
    <header class="header">
      <div class="left">Nacos-Go 控制台</div>
      <div class="right">
        <button class="logout-btn" @click="logout">退出登录</button>
      </div>
    </header>

    <!-- 导航栏 -->
    <nav class="nav">
      <router-link
        v-for="item in menu"
        :key="item.path"
        :to="item.path"
        class="nav-item"
        active-class="active"
      >
        {{ item.title }}
      </router-link>
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
      { path: '/tenant', title: 'Tenant' },
      { path: '/service', title: '服务' },
      { path: '/config',  title: '配置' },
      { path: '/publish', title: '发布' }
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
  background: #f7f7f7;
}

/* 顶部栏 */
.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0 24px;
  height: 56px;
  background: #fff;
  box-shadow: 0 2px 6px rgba(0, 0, 0, 0.05);
  font-size: 18px;
  font-weight: 500;
  color: #333;
}

.logout-btn {
  padding: 6px 16px;
  border: none;
  border-radius: 4px;
  background: #ff4d4f;
  color: #fff;
  font-size: 14px;
  cursor: pointer;
  transition: opacity 0.2s;
}
.logout-btn:hover {
  opacity: 0.85;
}

/* 导航栏 */
.nav {
  display: flex;
  gap: 4px;
  padding: 8px 24px;
  background: #fff;
  border-bottom: 1px solid #e6e6e6;
}

.nav-item {
  padding: 8px 16px;
  border-radius: 4px;
  font-size: 15px;
  color: #595959;
  text-decoration: none;
  transition: background 0.2s, color 0.2s;
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

/* 内容区 */
.main {
  padding: 24px;
}
</style>