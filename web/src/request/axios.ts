import axios from "axios";
import router from "@/router";
import { ElMessage } from "element-plus";

const service = axios.create({
    baseURL: process.env.VUE_APP_BASE_API,
    timeout: 5000
})

service.interceptors.request.use(
    config => {
        return config
    },
    error => {
        return Promise.reject(new Error(error))
    }
)

service.interceptors.response.use(
    response => {
        debugger
        return response
    }, error => {
        debugger
        return Promise.reject(new Error(error.response.data))
    }
)

interface BZResponse<U> {
    success: boolean;
    obj: U;
    errMsg?: string;
    code?: number;
}

function checkCode(res: any) {
    if (res.status === -404) {
        ElMessage.error("请求未找到")
    }
    if (res.data && !res.data.success) {
        ElMessage.error("未请求到数据")
    }
    return res
}

function checkStatus(response: any) {
    // 未登陆 || Token错误 || 账号禁用
    if ( response && (/^(200|304|400)$/.test(response.status)) ) {
        if (response.data && (response.data.code === 21004 || 
            response.data.code === 21005 || 
            response.data.code === 21006 || 
            response.data.code === 21008)) {
                router.push('/login');
        } else if (response.data.code !== 2000) { // 错误提示
            ElMessage.error(response.data.msg);
        } else if (response.data.success) {
            return response.data;
        } else if (response.data.msg) { // 提示
            ElMessage.error(response.data.msg);
            return response.data
        } else {
            ElMessage.error('请求出错')
            return false
        }
    } else {
        return false
    }
}

export default {
    service,
    getJson<U>(data: any, url: string) {
        let token: any = "Bearer "+window.localStorage.getItem("token")
        return axios({
            method: "GET",
            url: url,
            data: data,
            headers: {
                "Authorization": token,
                "key": "2A3F43B2E46DDAFC6E12C2B386704EF4",
                "secret": "7FE6461015477B0F24ADD487CBC4A398",
                "sign": "dbbcbb3266ce08f5b8549bcf43bda750",
                "Content-type": "application/json;charset=UTF-8"
            }
        }).then((res: any) => {
            return checkStatus(res);
        }).catch((err) => {
            return checkCode(err);
        });
    },
    postJson<U>(data: any, url: string) {
        let token: any = "Bearer "+window.localStorage.getItem("token")
        return axios({
            method: "POST",
            url: url,
            data: data,
            headers: {
                "Authorization": token,
                "key": "2A3F43B2E46DDAFC6E12C2B386704EF4",
                "secret": "7FE6461015477B0F24ADD487CBC4A398",
                "sign": "dbbcbb3266ce08f5b8549bcf43bda750",
                "Content-type": "application/json;charset=UTF-8"
            }
        }).then((res: any) => {
            return checkStatus(res);
        }).catch((err) => {
            return checkCode(err);
        });
    },

    postJson2<U>(data: any, url: string) {
        let token: any = "Bearer "+window.localStorage.getItem("token")
        return axios({
            method: "POST",
            url: url,
            data: data,
            headers: {
                "Authorization": token,
                "key": "2A3F43B2E46DDAFC6E12C2B386704EF4",
                "secret": "7FE6461015477B0F24ADD487CBC4A398",
                "sign": "dbbcbb3266ce08f5b8549bcf43bda750",
                'Content-Type': 'multipart/form-data'
            }
        }).then((res: any) => {
            return checkStatus(res);
        }).catch((err) => {
            return checkCode(err);
        });
    },
}
