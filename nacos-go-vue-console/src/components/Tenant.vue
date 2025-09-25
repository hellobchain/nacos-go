<template>
  <div class="card">
    <h2>Tenant 管理 <button @click="load" style="float:right">刷新</button></h2>
    <table>
      <thead><tr><th>Tenant</th><th>操作</th></tr></thead>
      <tbody>
        <tr v-for="t in list" :key="t"><td>{{t}}</td><td><button @click="del(t)">删除</button></td></tr>
        <tr v-if="!list.length"><td colspan="2">暂无 Tenant</td></tr>
      </tbody>
    </table>
    <div style="margin-top:8px">
      <input v-model="newT" placeholder="新 Tenant 名称"/>
      <button @click="add">添加</button>
    </div>
  </div>
</template>
<script>
import { getTenants, addTenant, delTenant } from '@/api/tenant'
import { Notify } from '@/components/Notify'
export default {
  data() {
    return { list: [], newT: '' }
  },
  created() { this.load() },
  methods: {
    async load() { this.list = await getTenants() },
    async add() {
      if (!this.newT) return Notify.error('名称不能为空')
      await addTenant(this.newT); this.newT = ''; Notify.success('添加成功'); this.load()
    },
    async del(t) {
      if (!confirm(`确认删除 Tenant：${t} ？`)) return
      await delTenant(t); Notify.success('删除成功'); this.load()
    }
  }
}
</script>