import { createStore } from 'vuex'
import { menu } from '@/request/api'
import createPersistedstate from 'vuex-persistedstate'

export default createStore({
  state: {
        /** 左侧菜单列表 */
        menuListAll: [],
        /** 左边导航栏: 当前展开的一级菜单名 */
        openNames: '',
        /** 左边菜单栏: 当前选中的菜单名 */
        activeName: '',
  },
  getters: {
  },
  mutations: {
    activeName(state: any, name) {
      state.activeName = name;
    },
    openNames(state: any, name) {
      state.openNames = name
    },
    // 菜单列表
    async menuList(state) {
      try {
        let res = await menu();
        if (res) {
          state.menuListAll = res.data
          // console.log(state.menuListAll);
                 
          // state.menuListAll = state.menuListAll
        }
      } catch (err) {

      }
    },
  },
  actions: {
    /** 请求左侧菜单 */
    loadmenuList(obj) {
      window.localStorage.getItem('token') && obj.commit('menuList')
    }
  },
  modules: {
  },
  plugins: [
    createPersistedstate({
      storage: window.sessionStorage,
      key: "store",
      reducer(state) {
        return {
          menuListAll: state.menuListAll,
          openNames: state.openNames,
          activeName: state.activeName
        }
      }
    })
  ]
})
