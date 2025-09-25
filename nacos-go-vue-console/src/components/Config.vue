<template>
  <div class="card">
    <h2>配置列表</h2>
    <table>
      <thead><tr><th>DataId</th><th>Group</th><th>内容</th><th>操作</th></tr></thead>
      <tbody>
        <tr v-for="c in rows" :key="c.dataId+c.group">
          <td>{{c.dataId}}</td><td>{{c.group}}</td><td><pre>{{c.content}}</pre></td>
          <td>
            <button @click="openEdit(c)">编辑</button>
            <button @click="del(c)">删除</button>
          </td>
        </tr>
        <tr v-if="!rows.length"><td colspan="4">暂无配置</td></tr>
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