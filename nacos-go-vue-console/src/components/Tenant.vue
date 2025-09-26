<template>
  <div class="card">
    <div class="header">
      <h2 class="title">空间管理</h2>
      <button class="refresh" @click="load">刷新</button>
    </div>

    <!-- 列表 -->
    <table class="list">
      <thead>
        <tr><th>Tenant</th><th>操作</th></tr>
      </thead>
      <tbody>
        <tr v-for="t in list" :key="t" class="hover-row">
          <td class="name">{{ t }}</td>
          <td class="actions">
            <button class="btn danger" @click="del(t)">删除</button>
          </td>
        </tr>
        <tr v-if="!list.length">
          <td colspan="2" class="empty">暂无空间</td>
        </tr>
      </tbody>
    </table>

    <!-- 新增 -->
    <div class="add-line">
      <input v-model="newT" class="input" placeholder="新空间名称" @keyup.enter="add" />
      <button class="btn primary" @click="add">添加</button>
    </div>
  </div>
</template>

<script>
/* 原有逻辑保持不变 */
import { getTenants, addTenant, delTenant } from '@/api/tenant'
import { Notify } from '@/components/Notify'
export default {
  data: () => ({ list: [], newT: '' }),
  created() { this.load() },
  methods: {
    async load() { this.list = await getTenants() },
    async add() {
      if (!this.newT) return Notify.error('名称不能为空')
      await addTenant(this.newT)
      this.newT = ''
      Notify.success('添加成功')
      this.load()
    },
    async del(t) {
      if (!confirm(`确认删除空间${t} ？`)) return
      await delTenant(t)
      Notify.success('删除成功')
      this.load()
    }
  }
}
</script>

<style scoped>
.card {
  background: #fff;
  border-radius: 6px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.06);
  padding: 20px 24px;
}

.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 16px;
}

.title {
  margin: 0;
  font-size: 24px;
  font-weight: 500;
}

.refresh {
  padding: 4px 12px;
  font-size: 13px;
  border: none;
  border-radius: 4px;
  background: #e6f7ff;
  color: #1890ff;
  cursor: pointer;
  transition: opacity 0.2s;
}
.refresh:hover {
  opacity: 0.85;
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

/* 斑马纹 */
.list tbody tr:nth-child(even) {
  background: #fcfcfc;
}

/* 悬浮高亮 */
.hover-row:hover {
  background: #f3f9ff !important;
}

.name {
  font-weight: 500;
  color: #262626;
}

.actions {
  text-align: right;
}

.empty {
  text-align: center;
  padding: 40px 0;
  color: #999;
}

.add-line {
  margin-top: 16px;
  display: flex;
  gap: 8px;
}

.input {
  flex: 1;
  padding: 6px 10px;
  font-size: 14px;
  border: 1px solid #d9d9d9;
  border-radius: 4px;
}

.btn {
  padding: 6px 16px;
  border: none;
  border-radius: 4px;
  font-size: 14px;
  cursor: pointer;
  transition: opacity 0.2s;
}
.btn:hover {
  opacity: 0.85;
}
.primary {
  background: #1890ff;
  color: #fff;
}
.danger {
  background: #ff4d4f;
  color: #fff;
}
</style>