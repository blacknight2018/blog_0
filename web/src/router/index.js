import Vue from 'vue'
import Router from 'vue-router'
import Home from '@/components/Home'
import Admin from '@/components/Admin'
import Archive from '@/components/Archive'
import Article from '@/components/Article'
import ArticleDetail from '@/components/ArticleDetail'

Vue.use(Router)

export default new Router({
    mode: 'history',
    routes: [{
        path: '/Home',
        name: 'Home',
        component: Home
    }, {
        path: '/index',
        redirect: '/'
    }, {
        path: '/admin',
        name: 'Admin',
        component: Admin
    }, {
        path: '/archive',
        name: 'archive',
        component: Archive
    }, {
        path: '/article',
        name: 'article',
        component: Article
    }, {
        path: '/articledetail',
        name: 'articledetail',
        component: ArticleDetail
    }]
})