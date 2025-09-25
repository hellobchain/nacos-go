const style = document.createElement('style')
style.innerHTML = `
#notifyBox{position:fixed;top:12px;left:50%;transform:translateX(-50%);z-index:9999;color:#fff;padding:6px 14px;border-radius:4px;font-size:14px;transition:opacity .3s;opacity:0;pointer-events:none}
.notify-success{background:#52c41a}
.notify-error{background:#ff4d4f}
.notify-info{background:#1890ff}
`
document.head.appendChild(style)
const box = document.createElement('div')
box.id = 'notifyBox'
document.body.appendChild(box)
function show(msg, type = 'info', duration = 2500) {
    box.textContent = msg
    box.className = 'notify-' + type
    box.style.opacity = '1'
    setTimeout(() => (box.style.opacity = '0'), duration)
}
export const Notify = { success: m => show(m, 'success'), error: m => show(m, 'error'), info: m => show(m, 'info') }