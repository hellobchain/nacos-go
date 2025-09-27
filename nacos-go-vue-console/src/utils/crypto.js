import CryptoJS from 'crypto-js'

const aseKey = 'kGx1ae1Sc0qCo88F' // 16位
// aes ecb 加密 补位
export function encryptByAesEcb(data) {
    const keyHex = CryptoJS.enc.Utf8.parse(aseKey)
    const encrypted = CryptoJS.AES.encrypt(data, keyHex, {
        mode: CryptoJS.mode.ECB,
        padding: CryptoJS.pad.ZeroPadding
    })
    return encrypted.ciphertext.toString()
}
