<template>
  <div>
    <!-- 第 1 行：退出按钮 -->
    <div id="logoutBar">
      <button @click="logout">退出登录</button>
    </div>

    <!-- 第 2 行：导航栏 -->
    <div class="nav">
      <router-link to="/tenant">Tenant</router-link>
      <router-link to="/service">服务</router-link>
      <router-link to="/config">配置</router-link>
      <router-link to="/publish">发布</router-link>
    </div>

    <!-- 内容区 -->
    <router-view class="main" />
  </div>
</template>
<script>
import { Notify } from '@/components/Notify'
export default {
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
.logoutBar {
  position: fixed;
  top: 0; left: 0;
  padding: 10px 20px;
  background: #fff;
  border-bottom: 1px solid #e6e6e6;
  z-index: 1;              /* 保持在最上层 */
}

.nav {
  margin-top: 10px;        /* 给固定定位的退出按钮留高 */
  display: flex;
  gap: 60px;
  padding: 5px 10px;
  background: #fff;
  border-bottom: 1px solid #e6e6e6;
  font-style: normal;
  font-size: 18px;
}

.main {
  padding: 20px;
  background: #f7f7f7;
  min-height: 100vh;
}
</style>
