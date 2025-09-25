<template>
  <div class="mask" @click.self="$emit('close')">
    <div class="dialog">
      <h3>编辑配置</h3>
      <input :value="item.dataId" disabled><br>
      <input :value="item.group" disabled><br>
      <textarea v-model="content" rows="10"></textarea><br>
      <button @click="save">保存</button>
      <button @click="$emit('close')">取消</button>
    </div>
  </div>
</template>
<script>
import { publish } from '@/api/config'
import { Notify } from '@/components/Notify'
export default {
  props: { item: Object },
  data() { return { content: this.item.content } },
  methods: {
    async save() {
      if (!this.content) return Notify.error('内容不能为空')
      await publish({ ...this.item, content: this.content })
      Notify.success('保存成功')
      this.$emit('ok'); this.$emit('close')
    }
  }
}
</script>
<style scoped>
.mask{position:fixed;inset:0;background:rgba(0,0,0,.4);z-index:9999}
.dialog{position:absolute;top:50%;left:50%;transform:translate(-50%,-50%);background:#fff;padding:20px;border-radius:4px;width:420px}
input,textarea{width:100%;margin-bottom:8px}
</style>