<template>
  <div class="wrapper">
    <div class="glass">
      <div class="header">
        <h2 class="title">用户管理</h2>
        <button class="btn primary" @click="openEdit()">新增用户</button>
      </div>

      <!-- 列表 -->
      <table class="list">
        <thead>
          <tr>
            <th>用户名</th>
            <th>角色</th>
            <th>操作</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="u in users" :key="u.username" class="hover-row">
            <td>{{ u.username }}</td>
            <td>{{ u.role || '-' }}</td>
            <td class="actions">
              <button class="btn text" v-if="u.role !== 'admin'" @click="openEdit(u)">修改</button>
              <button class="btn text danger" v-if="u.role !== 'admin'" @click="del(u.username)">删除</button>
            </td>
          </tr>
          <tr v-if="!users.length">
            <td colspan="3" class="empty">暂无用户</td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- 新增/编辑弹框 -->
    <transition name="fade">
      <div v-if="showDialog" class="mask" @click.self="closeDialog">
        <div class="dialog">
          <h3 class="dialogTitle">{{ isAdd ? '新增用户' : '编辑用户' }}</h3>

          <label class="label">用户名</label>
          <input
            v-model="form.username"
            class="input"
            :disabled="!isAdd"
            placeholder="请输入用户名"
          />

          <label class="label">密码</label>
          <input
            v-model="form.password"
            type="password"
            class="input"
            placeholder="6-20 位字符（留空则不改）"
          />

          <label class="label">角色</label>
          <select v-model="form.role" class="input">
            <option value="user">普通用户</option>
            <option value="admin">管理员</option>
          </select>

          <div class="dialogActions">
            <button class="btn" @click="closeDialog">取消</button>
            <button class="btn primary" @click="submit">保存</button>
          </div>
        </div>
      </div>
    </transition>
  </div>
</template>

<script>
import { Notify } from '@/components/Notify'
import { addUser,updateUser,deleteUser,getUserList } from '@/api/auth'
export default {
  name: 'User',
  data() {
    return {
      users: [],
      showDialog: false,
      isAdd: true,
      form: { username: '', password: '', role: 'user' }
    }
  },
  created() {
    this.load()
  },
  methods: {
    async load() {
      this.users = await getUserList()
    },
    openEdit(user = null) {
      this.isAdd = !user
      this.form = user
        ? { username: user.username, password: '', role: user.role || 'user' }
        : { username: '', password: '', role: 'user' }
      this.showDialog = true
    },
    closeDialog() {
      this.showDialog = false
    },
    async submit() {
      const { username, password, role } = this.form
      if (!username) return Notify.warning('请输入用户名')
      if (this.isAdd && !password) return Notify.warning('请输入密码')
      if (password && password.length < 4) return Notify.warning('密码至少 4 位')

      try {
        if (this.isAdd) {
          await addUser({ username, password, role })
          Notify.success('新增成功')
        } else {
          await updateUser({ username, password, role })
          Notify.success('修改成功')
        }
        this.closeDialog()
        this.load()
      } catch (e) {
        Notify.error(e.message || '操作失败')
      }
    },
    async del(username) {
      if (!confirm(`确认删除用户 ${username} ？`)) return
      await deleteUser(username)
      Notify.success('删除成功')
      this.load()
    }
  }
}
</script>

<style scoped>
.wrapper {
  min-height: 100vh;
  display: flex;
  align-items: flex-start;
  justify-content: center;
  background: #f7f7f7;
  padding: 40px 20px;
}

.glass {
  width: 800px;
  max-width: 90vw;
  background: #fff;
  border-radius: 12px;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.08);
  padding: 28px 32px;
}

.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
}

.title {
  margin: 0;
  font-size: 22px;
  font-weight: 500;
  color: #262626;
}

.list {
  width: 100%;
  border-collapse: collapse;
  font-size: 14px;
}

.list th {
  background: #fafafa;
  padding: 10px 12px;
  text-align: left;
  font-weight: 600;
  color: #333;
}

.list td {
  padding: 10px 12px;
  border-bottom: 1px solid #f0f0f0;
}

.hover-row:hover {
  background: #f3f9ff !important;
}

.actions {
  text-align: right;
}

.empty {
  text-align: center;
  padding: 40px 0;
  color: #999;
}

.btn {
  padding: 4px 12px;
  font-size: 13px;
  border: none;
  border-radius: 4px;
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
.text {
  background: transparent;
  color: #1890ff;
}
.text.danger {
  color: #ff4d4f;
}

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
  width: 420px;
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

.label {
  font-size: 13px;
  color: #595959;
  margin-bottom: -8px;
}

.input {
  width: 100%;
  padding: 8px 12px;
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

.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.25s;
}
.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>