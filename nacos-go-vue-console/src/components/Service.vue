<template>
  <div class="card">
    <h2 class="title">服务实例</h2>

    <!-- 空状态 -->
    <div v-if="!rows.length" class="empty">暂无实例</div>

    <!-- 表格 -->
    <table v-else class="list">
      <thead>
        <tr>
          <th>服务名</th>
          <th>分组</th>
          <th>IP:端口</th>
          <th>健康</th>
          <th>操作</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="i in rows" :key="i.ip+i.port" class="hover-row">
          <td class="name">{{ i.serviceName }}</td>
          <td class="ellipsis">{{ i.groupName }}</td>
          <td>{{ i.ip }}:{{ i.port }}</td>
          <td>
            <span :class="['tag', i.healthy ? 'success' : 'danger']">
              {{ i.healthy ? '健康' : '下线' }}
            </span>
          </td>
          <td class="actions">
            <button class="btn danger" @click="der(i.ip,i.port)">下线</button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
import { getInstances, deregister } from '@/api/service'
import { Notify } from '@/components/Notify'
export default {
  data: () => ({ rows: [] }),
  created() { this.load() },
  methods: {
    async load() { this.rows = await getInstances() },
    async der(ip, port) {
      await deregister(ip, port)
      Notify.success('下线成功')
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

.title {
  margin: 0 0 16px;
  font-size: 20px;
  font-weight: 500;
  text-align: center;
}

.empty {
  text-align: center;
  padding: 60px 0;
  color: #999;
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

.ellipsis {
  max-width: 160px;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

/* 状态标签 */
.tag {
  display: inline-block;
  padding: 4px 10px;
  border-radius: 4px;
  font-size: 12px;
  font-weight: 500;
}
.success {
  background: #f6ffed;
  color: #52c41a;
  border: 1px solid #b7eb8f;
}
.danger {
  background: #fff1f0;
  color: #ff4d4f;
  border: 1px solid #ffccc7;
}

.actions {
  text-align: right;
}

.btn {
  padding: 4px 12px;
  border: none;
  border-radius: 4px;
  font-size: 13px;
  cursor: pointer;
  transition: opacity 0.2s;
}
.btn:hover {
  opacity: 0.85;
}
.danger {
  background: #ff4d4f;
  color: #fff;
}
</style>