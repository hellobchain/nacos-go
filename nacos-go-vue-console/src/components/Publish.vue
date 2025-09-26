<template>
  <div class="card">
    <div class="box"> 
      <h2 class='title'>发布配置</h2>
      <input class="input" v-model="form.dataId" placeholder="DataId"/><br>
      <input class="input" v-model="form.group" placeholder="Group"/><br>
      <select class="input" v-model="form.tenant">
        <option value="">（空）</option>
        <option v-for="t in tenants" :key="t" :value="t">{{t}}</option>
      </select><br>
      <textarea class="textarea" v-model="form.content" rows="6" placeholder="配置内容"></textarea><br>
      <button @click="submit">发布</button>
    </div>
  </div>
</template>
<script>
import { publish } from '@/api/config'
import { getTenants } from '@/api/tenant'
import { Notify } from '@/components/Notify'
export default {
  data() {
    return { tenants: [], form: { dataId: '', group: '', tenant: '', content: '' } }
  },
  created() { getTenants().then(r => (this.tenants = r)) },
  methods: {
    async submit() {
      const f = this.form
      if (!f.dataId || !f.group || !f.content) return Notify.error('请填写完整')
      await publish(f)
      Notify.success('发布成功')
      Object.keys(f).forEach(k => (f[k] = ''))
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

.textarea {
  width: 80%;
  font-size: large;
  align-self: center;
}
</style>
