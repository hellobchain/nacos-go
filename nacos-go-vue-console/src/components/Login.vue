<template>
  <!-- 全屏遮罩，上下左右居中 -->
  <div class="wrapper">
    <div class="box">
      <h2 class="title">登录</h2> 
      <input v-model="user" placeholder="用户名" style="width: 80%; font-size: large; align-self: center;" />
      <input v-model="pwd" type="password" placeholder="密码"  style="width: 80%; font-size: large; align-self: center;" 
      @keyup.enter="handleLogin" />
      <button @click="handleLogin">登录</button>
    </div>
  </div>
</template>

<script>
import { login } from '@/api/auth'
import { Notify } from '@/components/Notify'
export default {
  data() {
    return { user: '', pwd: '' }
  },
  methods: {
    async handleLogin() {
      if (!this.user || !this.pwd) return Notify.error('请输入账号密码')
      try {
        const res = await login({ username: this.user, password: this.pwd })
        this.$store.commit('SET_TOKEN', res.accessToken)
        this.$router.replace('/')
      } catch {
        Notify.error('登录失败，请检查账号密码')
      }
    }
  }
}
</script>

<style scoped>
/* 关键代码：绝对居中 */
.wrapper {
  position: fixed;
  inset: 0;               /* top:0 right:0 bottom:0 left:0 */
  display: flex;
  align-items: center;    /* 垂直居中 */
  justify-content: center;/* 水平居中 */
  background: #f7f7f7;    /* 可视背景，可删 */
}

.box {
  width: 320px;
  padding: 40px 24px;
  background: #fff;
  border-radius: 6px;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.1);
  display: flex;
  flex-direction: column;
  gap: 12px;
}

button {
  align-self: center;     /* 按钮居中 */
  padding: 6px 20px;
  font-size: large;
}

.title {
  text-align: center;   /* 水平居中 */
  font-size: 24px;      /* 字号变大，默认 ~16px */
  margin: 0 0 16px;     /* 去掉默认上下外边距，再留点底距 */
}
</style>
