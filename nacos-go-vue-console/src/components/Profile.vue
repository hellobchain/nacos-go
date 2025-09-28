<template>
  <div class="wrapper">
    <div class="glass">
      <h2 class="title">个人中心</h2>

      <!-- 用户信息 -->
      <div class="info">
        <div class="row">
          <span class="label">用户名称：</span>
          <span class="val">{{ userInfo.username }}</span>
        </div>
        <div class="row">
          <span class="label">角色：</span>
          <span class="val">{{ userInfo.role || '普通用户' }}</span>
        </div>
      </div>

      <!-- 功能按钮 -->
      <div class="actions">
        <button class="btn primary" @click="openPwdDialog">修改密码</button>
      </div>
    </div>

    <!-- 修改密码弹框 -->
    <transition name="fade">
      <div v-if="showDialog" class="mask" @click.self="closeDialog">
        <div class="dialog">
          <h3 class="dialogTitle">修改密码</h3>

          <label class="label">原密码</label>
          <input
            v-model="pwdForm.oldPwd"
            type="password"
            class="input"
            placeholder="请输入原密码"
          />

          <label class="label">新密码</label>
          <input
            v-model="pwdForm.newPwd"
            type="password"
            class="input"
            placeholder="4-20 位字符"
          />

          <label class="label">确认密码</label>
          <input
            v-model="pwdForm.confirm"
            type="password"
            class="input"
            placeholder="再次输入新密码"
            @keyup.enter="submitPwd"
          />

          <div class="dialogActions">
            <button class="btn" @click="closeDialog">取消</button>
            <button class="btn primary" @click="submitPwd">保存</button>
          </div>
        </div>
      </div>
    </transition>
  </div>
</template>

<script>
import { Notify } from '@/components/Notify'
import { changePassword } from '@/api/auth'
import { encryptByAesEcb } from '@/utils/crypto'

export default {
  name: 'Profile',
  data() {
    return {
      userInfo: {
        username: localStorage.getItem('username') || 'nacos',
        role: localStorage.getItem('role') || '管理员'
      },
      showDialog: false,
      pwdForm: {
        oldPwd: '',
        newPwd: '',
        confirm: ''
      }
    }
  },
  created() {
    // 演示：从 localStorage 拿用户名，真实项目请改 vuex
    this.userInfo.username = localStorage.getItem('username') || 'nacos'
  },
  methods: {
    openPwdDialog() {
      this.pwdForm = { oldPwd: '', newPwd: '', confirm: '' }
      this.showDialog = true
    },
    closeDialog() {
      this.showDialog = false
    },
    async submitPwd() {
      const { oldPwd, newPwd, confirm } = this.pwdForm
      if (!oldPwd || !newPwd) return Notify.warning('请填写完整')
      if (newPwd.length < 4) return Notify.warning('新密码至少 4 位')
      if (newPwd !== confirm) return Notify.warning('两次密码不一致')
      try {
        await changePassword({
          username: this.userInfo.username,
          oldPassword: encryptByAesEcb(oldPwd),
          newPassword: encryptByAesEcb(newPwd)
        })
        Notify.success('密码已修改，请重新登录')
        this.closeDialog()
        // 修改密码后常见做法：直接退出
        this.$store.commit('SET_TOKEN', '')
        this.$router.replace('/login')
      } catch (e) {
        Notify.error(e.message || '修改失败')
      }
    }
  }
}
</script>

<style scoped>
.wrapper {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg, #83a4d4 0%, #b6fbff 100%);
  padding: 40px 20px;
}

.glass {
  width: 480px;
  max-width: 90vw;
  background: rgba(255, 255, 255, 0.75);
  border-radius: 16px;
  backdrop-filter: blur(10px);
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.15);
  padding: 32px;
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.title {
  margin: 0 0 8px;
  font-size: 24px;
  font-weight: 500;
  text-align: center;
  color: #262626;
}

.info {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.row {
  display: flex;
  align-items: center;
}

.label {
  width: 80px;
  font-size: 15px;
  color: #595959;
}

.val {
  font-size: 15px;
  color: #262626;
  font-weight: 500;
}

.actions {
  display: flex;
  justify-content: center;
  margin-top: 8px;
}

.btn {
  padding: 8px 20px;
  font-size: 15px;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  transition: opacity 0.2s;
}
.btn:hover {
  opacity: 0.85;
}
.primary {
  background: linear-gradient(90deg, #1890ff 0%, #45b7ff 100%);
  color: #fff;
}

/* 弹窗遮罩 */
.mask {
  position: fixed;
  inset: 0;
  background: rgba(0, 0, 0, 0.45);
  backdrop-filter: blur(4px);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 999;
}

.dialog {
  width: 440px;
  max-width: 90vw;
  background: #fff;
  border-radius: 12px;
  box-shadow: 0 12px 36px rgba(0, 0, 0, 0.18);
  padding: 24px 28px;
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.dialogTitle {
  margin: 0 0 8px;
  font-size: 18px;
  font-weight: 500;
  text-align: center;
}

.input {
  width: 94%;
  padding: 10px 14px;
  font-size: 14px;
  border: 1px solid #d9d9d9;
  border-radius: 6px;
  transition: border-color 0.3s, box-shadow 0.3s;
}
.input:focus {
  border-color: #1890ff;
  box-shadow: 0 0 0 3px rgba(24, 144, 255, 0.15);
  outline: none;
}

.dialogActions {
  display: flex;
  gap: 12px;
  justify-content: flex-end;
  margin-top: 8px;
}

/* 动画 */
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.25s;
}
.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>