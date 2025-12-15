<template>  
    <div class="term1">
        <div ref="terminalBox" style="height: 60vh;"></div>
    </div>
</template>

<script setup>
import 'xterm/css/xterm.css'
import { Terminal } from 'xterm'
import { FitAddon } from 'xterm-addon-fit'
import { ref, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router';

let terminalBox = ref(null)
let term
let socket
const router = useRouter();
const route = useRoute();

const ipaddress  = ref(route.query.ipaddress);
const username = ref(route.query.username);
const password = ref(route.query.password);
const port = ref(route.query.port);


onMounted(() => {
    //创建一个客户端
    term = new Terminal({
        rendererType: 'canvas', //使用这个能解决vim不显示或乱码
        cursorBlink: true,
        cursorStyle: "bar",
    })
    // term.write
    // 将客户端挂载到dom上
    const fitAddon = new FitAddon()
    term.loadAddon(fitAddon)
    term.open(terminalBox.value)
    fitAddon.fit()

    // 创建socket连接
    term.write('正在连接...\r\n');
    socket = new WebSocket("ws://" + location.hostname +  ":5555/ssh")

    socket.binaryType = "arraybuffer";

    // 打开socket监听事件的方法
    socket.onopen = function () {
        term.write('连接成功...\r\n');
        fitAddon.fit()
        term.onData(function (data) {
            // socket.send(JSON.stringify({ type: "stdin", data: data }))
            // console.log(data)
            socket.send(data)
            // console.log(data)
        });
        // ElMessage.success("会话成功连接！")
        var jsonStr = `{"username":"${username.value}", "ipaddress":"${ipaddress.value}", "port":${port.value}, "password":"${password.value}"}`
        var datMsg = window.btoa(jsonStr)
        // socket.send(JSON.stringify({ ip: ip.value, name: name.value, password: password.value }))
        socket.send(datMsg)
    }
    socket.onclose = function () {
        term.writeln('连接关闭');
    }
    socket.onerror = function (err) {
        // console.log(err)
        term.writeln('读取数据异常：', err);
    }
    // 接收数据
    socket.onmessage = function (recv) {
        try {
            term.write(recv.data)
        } catch (e) {
            console.log('unsupport data', recv.data)
        }
    }

    window.addEventListener("resize", () => {
        fitAddon.fit()
    }, false)

})


</script>

<style lang="scss" scoped>
.upload {
    min-height: 100px;
}
.term1 {
    margin-left: 60px;
}
.go_out {
    margin-left: -89%;
    margin-top: 20px;
    margin-bottom: 20px;
}
</style>
