<template>
      <el-container class="app-wrapper">
        <!-- 菜单 -->
        <el-aside :width="asideWidth" class="sidebar-container">
            <el-scrollbar style="height:'100%'">
                <el-menu
                    active-text-color="#ffd04b"
                    background-color="#20222A"
                    class="el-menu-vertical-demo"
                    :default-active="activePath"
                    text-color="#fff"
                    unique-opened
                    router
                >
                    <el-menu-item index="board" @click="goUrl('web-ssh-terminal','','board')" style="height: 50px !important;background-color: var(--el-menu-hover-bg-color)  !important;">
                        <span class="backstage-manage">web-ssh-terminal</span>
                    </el-menu-item>
            
                    <el-sub-menu>
                        <template #title>

                            <span class="caidan-auth">节点服务器</span>
                        </template>
                        <el-menu-item>
                    
                            <span class="caidan-auth">节点服务器</span>
                        </el-menu-item>
                    </el-sub-menu>
                </el-menu>
            </el-scrollbar>
        </el-aside>

        <!-- 菜单end -->
        <el-container class="container">
            <!-- 头部 -->
            <el-header  style="font-size: 16px;background-color: #ffffff">
                <div class="toolbar">
                    <div class="toolbar-open-close">
                            <div>
                                <i class="layui-icon"></i>
                            </div>
                    </div>
                    <div class="toolbar-static">

                </div>
                </div>
            </el-header>
            <!-- 头部end -->
            <!-- 面包屑 -->
            <!-- <breadcrumb></breadcrumb> -->
            <!-- 内容 -->
            <el-main style="background-color: #f2f2f2">
                <router-view />
            </el-main>
            <!-- 内容end -->
        </el-container>
    </el-container>
</template>

<script setup>
import { ref,computed,watch } from "vue"                                                                                                                                                                                                                 
import { useRouter } from "vue-router";
import {Location,Document,Menu as IconMenu,Setting,} from '@element-plus/icons-vue'
import variables from "@/assets/styles/variables.scss"
import store from "@/store/index.ts"  
import { ElMessage } from 'element-plus';



const asideWidth = computed(() => {
    return "210"
})

// // 保存信息和url到缓存里面
// const goUrl = (authName, authName2, url) => {
//     store.commit("openNames", authName);
//     store.commit("activeName", authName2);
//     sessionStorage.setItem(`url`,`${url}`)
//     activePath.value = sessionStorage.getItem(`url`)   
// }


</script>

<style lang="scss" scoped>
@import "../../assets/css/index.scss";
</style>
