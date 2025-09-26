<template>
  <div class="card">
    <h2 class="title">配置列表</h2>

    <!-- 空状态 -->
    <div v-if="!rows.length" class="empty">暂无配置</div>

    <!-- 表格 -->
    <table v-else class="config-table">
      <thead>
        <tr>
          <th>DataId</th>
          <th>Group</th>
          <th>内容</th>
          <th>操作</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="c in rows" :key="c.dataId+c.group" class="hover-row">
          <td class="ellipsis">{{ c.dataId }}</td>
          <td class="ellipsis">{{ c.group }}</td>
          <td><pre class="code">{{ c.content }}</pre></td>
          <td class="actions">
            <button class="btn primary" @click="openEdit(c)">编辑</button>
            <button class="btn danger" @click="del(c)">删除</button>
          </td>
        </tr>
      </tbody>
    </table>

    <Edit v-if="showEdit" :item="editItem" @close="showEdit=false" @ok="load"/>
  </div>
</template>

<script>
import { getConfigs, delConfig } from '@/api/config'
import { Notify } from '@/components/Notify'
import Edit from './Edit.vue'
export default {
  components: { Edit },
  data() { return { rows: [], showEdit: false, editItem: null } },
  created() { this.load() },
  methods: {
    async load() { this.rows = await getConfigs() },
    openEdit(c) { this.editItem = c; this.showEdit = true },
    async del(c) {
      if (!confirm(`确认删除 ${c.dataId}?`)) return
      await delConfig(c); Notify.success('删除成功'); this.load()
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
/* 大标题 */
.title {
  text-align: center;
  font-size: 24px;
  margin: 0 0 20px;
}

/* 空状态 */
.empty {
  text-align: center;
  padding: 60px 0;
  color: #999;
  font-size: 15px;
}

/* 整张表 */
.config-table {
  width: 100%;
  border-collapse: collapse;
  font-size: 14px;
  border-radius: 6px;
  overflow: hidden; /* 圆角 */
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.06);
}

/* 表头 */
.config-table th {
  background: #fafafa;
  padding: 10px 12px;
  text-align: left;
  font-weight: 600;
  color: #333;
}

/* 单元格 */
.config-table td {
  padding: 10px 12px;
  border-bottom: 1px solid #f0f0f0;
}

/* 斑马纹 */
.config-table tbody tr:nth-child(even) {
  background: #fcfcfc;
}

/* 悬浮高亮 */
.hover-row:hover {
  background: #f3f9ff !important;
}

/* 内容列：固定宽 + 滚动条 */
.code {
  margin: 0;
  max-height: 60px;
  overflow: auto;
  background: #f6f6f6;
  padding: 4px 6px;
  border-radius: 4px;
  font-size: 13px;
  line-height: 1.4;
}

/* 省略号 */
.ellipsis {
  max-width: 180px;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

/* 按钮组 */
.actions {
  display: flex;
  gap: 8px;
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
.primary {
  background: #1890ff;
  color: #fff;
}
.danger {
  background: #ff4d4f;
  color: #fff;
}
</style>
