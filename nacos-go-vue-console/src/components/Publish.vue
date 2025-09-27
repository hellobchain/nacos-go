<template>
  <div class="wrapper">
    <div class="glass">
      <h2 class="title">发布配置</h2>

      <input class="input" v-model="form.dataId" placeholder="DataId" />
      <input class="input" v-model="form.group" placeholder="Group" />

      <select class="input" v-model="form.tenant">
        <!-- <option value="">（空）</option> -->
        <option v-for="t in tenants" :key="t" :value="t">{{ t }}</option>
      </select>

      <textarea class="textarea" v-model="form.content" rows="6" placeholder="配置内容" />

      <button class="btn" @click="submit">发布</button>
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
/* 全屏渐变背景 */
.wrapper {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg, #83a4d4 0%, #b6fbff 100%);
  padding: 40px 20px;
}

/* 玻璃卡片 */
.glass {
  width: 480px;
  max-width: 90vw;
  padding: 32px;
  background: rgba(255, 255, 255, 0.75);
  border-radius: 16px;
  backdrop-filter: blur(10px);
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.15);
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.title {
  margin: 0 0 8px;
  font-size: 24px;
  font-weight: 500;
  text-align: center;
  color: #262626;
}

.input,
.textarea {
  width: 100%;
  padding: 10px 14px;
  font-size: 15px;
  border: 1px solid #d9d9d9;
  border-radius: 8px;
  transition: border-color 0.3s, box-shadow 0.3s;
  background: #fff;
}

.input:focus,
.textarea:focus {
  border-color: #1890ff;
  box-shadow: 0 0 0 3px rgba(24, 144, 255, 0.15);
  outline: none;
}

.textarea {
  resize: vertical;
  min-height: 120px;
}

.btn {
  margin-top: 8px;
  width: 100%;
  padding: 12px;
  font-size: 16px;
  font-weight: 500;
  color: #fff;
  background: linear-gradient(90deg, #1890ff 0%, #45b7ff 100%);
  border: none;
  border-radius: 8px;
  cursor: pointer;
  transition: transform 0.2s, box-shadow 0.2s;
}

.btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 16px rgba(24, 144, 255, 0.35);
}

.btn:active {
  transform: translateY(0);
}
</style>

