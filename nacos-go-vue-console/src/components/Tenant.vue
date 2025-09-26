<template>
  <div class="card">
    <div class="box">
    <h2 class="title">Tenant 管理 <button @click="load" style="float:right">刷新</button></h2>
    <table>
      <thead><tr><th>Tenant</th><th>操作</th></tr></thead>
      <tbody>
        <tr v-for="t in list" :key="t"><td>{{t}}</td><td><button @click="del(t)">删除</button></td></tr>
        <tr v-if="!list.length"><td colspan="2">暂无 Tenant</td></tr>
      </tbody>
    </table>
    <div style="margin-top:8px">
      <input class="input" v-model="newT" placeholder="新 Tenant 名称"/>
      <button @click="add">添加</button>
    </div>
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

<style scoped>
.box {
  width: 600px;
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

.input {
  width: 80%;
  font-size: large;
  align-self: center;
  height: 35px;
}

</style>
