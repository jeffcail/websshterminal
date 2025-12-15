import axios from "./axios";

// 左侧菜单列表
export const menu = () => {
    return axios.getJson("", "/api/menus/list")
}
