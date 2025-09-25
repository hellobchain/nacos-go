<template>
  <div class="card">
    <h2>发布配置</h2>
    <input v-model="form.dataId" placeholder="DataId"/><br>
    <input v-model="form.group" placeholder="Group"/><br>
    <select v-model="form.tenant">
      <option value="">（空）</option>
      <option v-for="t in tenants" :key="t" :value="t">{{t}}</option>
    </select><br>
    <textarea v-model="form.content" rows="6" placeholder="配置内容"></textarea><br>
    <button @click="submit">发布</button>
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
