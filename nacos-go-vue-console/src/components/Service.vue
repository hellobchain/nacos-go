<template>
  <div class="card">
    <div class="box">
    <h2 class="title">服务实例</h2>
    <table>
      <thead><tr><th>服务名</th><th>分组</th><th>IP:端口</th><th>健康</th><th>操作</th></tr></thead>
      <tbody>
        <tr v-for="i in rows" :key="i.ip+i.port">
          <td>{{i.serviceName}}</td><td>{{i.groupName}}</td><td>{{i.ip}}:{{i.port}}</td>
          <td><span :class="i.healthy?'green':'red'">{{i.healthy?'健康':'下线'}}</span></td>
          <td><button @click="der(i.ip,i.port)">下线</button></td>
        </tr>
        <tr v-if="!rows.length"><td colspan="5">暂无实例</td></tr>
      </tbody>
    </table>
  </div>
</div>
</template>
<script>
import { getInstances, deregister } from '@/api/service'
import { Notify } from '@/components/Notify'
export default {
  data() { return { rows: [] } },
  created() { this.load() },
  methods: {
    async load() { this.rows = await getInstances() },
    async der(ip, port) {
      await deregister(ip, port); Notify.success('下线成功'); this.load()
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
</style>
