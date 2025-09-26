<template>
  <transition name="fade">
    <div class="mask" @click.self="$emit('close')">
      <div class="dialog">
        <h3 class="title">编辑配置</h3>

        <label class="label">DataId</label>
        <input class="input" :value="item.dataId" disabled />

        <label class="label">Group</label>
        <input class="input" :value="item.group" disabled />

        <label class="label">内容</label>
        <textarea class="textarea" v-model="content" rows="10"></textarea>

        <div class="actions">
          <button class="btn primary" @click="save">保存</button>
          <button class="btn" @click="$emit('close')">取消</button>
        </div>
      </div>
    </div>
  </transition>
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
/* 遮罩 */
.mask {
  position: fixed;
  inset: 0;
  background: rgba(0, 0, 0, 0.45);
  backdrop-filter: blur(4px);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 999;
}

/* 对话框 */
.dialog {
  width: 520px;
  max-width: 90vw;
  background: #fff;
  border-radius: 12px;
  box-shadow: 0 12px 36px rgba(0, 0, 0, 0.18);
  padding: 24px 28px;
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.title {
  margin: 0 0 8px;
  font-size: 20px;
  font-weight: 500;
  text-align: center;
}

.label {
  font-size: 13px;
  color: #595959;
  margin-bottom: -8px;
}

.input,
.textarea {
  width: 100%;
  padding: 8px 12px;
  font-size: 14px;
  border: 1px solid #d9d9d9;
  border-radius: 6px;
  background: #fafafa;
}
.input:disabled {
  color: #8c8c8c;
}
.textarea {
  resize: vertical;
  min-height: 160px;
}

.actions {
  display: flex;
  gap: 12px;
  justify-content: flex-end;
  margin-top: 8px;
}

.btn {
  padding: 6px 16px;
  font-size: 14px;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  transition: opacity 0.2s;
}
.btn:hover {
  opacity: 0.85;
}
.primary {
  background: linear-gradient(90deg, #1890ff 0%, #45b7ff 100%);
  color: #fff;
}

/* 进出动画 */
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.25s;
}
.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>