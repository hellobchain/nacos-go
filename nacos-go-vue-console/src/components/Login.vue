<template>
  <div class="wrapper">
    <div class="glass">
      <h2 class="title">登录</h2>

      <input v-model="user" class="input" placeholder="用户名" />
      <input v-model="pwd" class="input" type="password" placeholder="密码" @keyup.enter="handleLogin" />

      <button class="btn" @click="handleLogin">登录</button>
    </div>
  </div>
</template>

<script>
import { login,getUserInfo } from '@/api/auth'

import { Notify } from '@/components/Notify'
import { encryptByAesEcb } from '@/utils/crypto'
export default {
  data() {
    return { user: '', pwd: '' }
  },
  methods: {
    async handleLogin() {
      if (!this.user || !this.pwd) return Notify.error('请输入账号密码')
      try {
        const res = await login({ username: this.user, password: encryptByAesEcb(this.pwd) })
        Notify.success('登录成功')
        this.$store.commit('SET_TOKEN', res.accessToken)
        // 获取用户信息
        const userInfo = await getUserInfo()
        localStorage.setItem('userName', userInfo.userName)
        localStorage.setItem('role', userInfo.role)
        this.$router.replace('/')

      } catch {
        Notify.error('登录失败，请检查账号密码')
      }
    }
  }
}
</script>

<style scoped>
/* 全屏居中 */
.wrapper {
  height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg, #83a4d4 0%, #b6fbff 100%);
}

/* 玻璃卡片 */
.glass {
  width: 360px;
  padding: 40px 32px;
  background: rgba(255, 255, 255, 0.75);
  border-radius: 16px;
  backdrop-filter: blur(10px);
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.15);
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 20px;
}

/* 标题 */
.title {
  margin: 0 0 8px;
  font-size: 26px;
  font-weight: 500;
  color: #262626;
}

/* 输入框 */
.input {
  width: 100%;
  padding: 12px 16px;
  font-size: 16px;
  border: 1px solid #d9d9d9;
  border-radius: 8px;
  transition: border-color 0.3s, box-shadow 0.3s;
}
/* 输入框 - 聚焦 */
.input:focus {
  border-color: #1890ff;
  box-shadow: 0 0 0 3px rgba(24, 144, 255, 0.15);
  outline: none;
}
/* 按钮 */
.btn {
  width: 100%;
  padding: 12px;
  font-size: 16px;
  font-weight: 500;
  color: #fff;
  background: linear-gradient(90deg, #1890ff 0%, #45b7ff 100%);
  border: none;
  border-radius: 8px;
  cursor: pointer;
  transition: transform 0.2s, box-shadow 0.2s;
}
/* 按钮 - 悬浮和点击 */
.btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 16px rgba(24, 144, 255, 0.35);
}
/* 按钮 - 点击 */
.btn:active {
  transform: translateY(0);
}
</style>

